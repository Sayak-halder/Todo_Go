package main

import (
	"fmt"
	"log"
	"os"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct{
	ID 			primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Completed 	bool 				`json:"completed"`
	Body    	string 				`json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Server Up!!")
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept", 
	}))

	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error in loading .env file!")
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client,err := mongo.Connect(context.Background(),clientOptions)

	if err != nil{
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(),nil)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGODB Atlas!")

	collection = client.Database("golang_db").Collection("todos")

	PORT := os.Getenv("PORT")

	app.Get("/api/todos",getTodo)
	app.Post("/api/todos",createTodo)
	app.Patch("/api/todos/:id",updateTodo)
	app.Delete("/api/todos/:id",deleteTodo)

	log.Fatal(app.Listen("0.0.0.0:"+PORT))
}

//Get the todo collection
func getTodo(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(),bson.M{})

	if err!=nil{
		return c.Status(500).SendString("Error in fetching todos")
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()){
		var todo Todo
		if err := cursor.Decode(&todo); err!=nil{
			return c.Status(500).SendString("Error in fetching todos");
		}
		todos=append(todos,todo)
	}

	if len(todos)==0{
		return c.Status(200).SendString("No todos found")
	}

	return c.JSON(todos)
}

// Created a Todo
func createTodo(c *fiber.Ctx)error{
	todo := new(Todo)

	if err := c.BodyParser(todo);err!=nil{
		return c.Status(500).SendString("Error in parsing body")
	}

	if todo.Body == ""{
		return c.Status(400).JSON(fiber.Map{"error":"Body cannot be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(),todo)
	if err!=nil{
		return c.Status(500).SendString("Error in inserting todo")
	}

	todo.ID=insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(200).JSON(fiber.Map{"success":"Task created successfully!"})
}

// Upadetd a Todo
func updateTodo(c *fiber.Ctx) error{
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":"Invalid id!"})
	}

	filter := bson.M{"_id":objectID}
	update := bson.M{"$set":bson.M{"completed":true}}

	_,err = collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		return c.Status(500).JSON(fiber.Map{"error":"Error in updating todo!"})
	}

	return c.Status(200).JSON(fiber.Map{"success":true})
}

// Delete a todo
func deleteTodo(c *fiber.Ctx)error{
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":"Invalid id!"})
	}

	filter := bson.M{"_id":objectID}
	_,err = collection.DeleteOne(context.Background(),filter)

	if err!=nil{
		return c.Status(500).JSON(fiber.Map{"error":"Error in deleting todo!"})
	}

	return c.Status(200).JSON(fiber.Map{"success":"Task deleted successfully!"})
}