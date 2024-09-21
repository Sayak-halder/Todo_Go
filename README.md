# Simple Todo App

A minimal Todo app built with a Go backend and a React frontend using Chakra UI and `react-query` for state management.

## Features

- Add, edit, and delete tasks
- Responsive design with light/dark mode support
- Real-time task updates using `react-query`
- MongoDB for persistent storage

## Project Structure

```shell
├── client          # Frontend (React)
│   ├── src
│   │   ├── assets  # Images & icons
│   │   ├── chakra  # Theme configuration
│   │   └── components
│   │       ├── Navbar.tsx    # Top navigation
│   │       ├── TodoForm.tsx  # Form for adding todos
│   │       ├── TodoList.tsx  # List displaying todos
│   │       └── TodoItem.tsx  # Individual task item
├── server          # Backend (Go)
├── .env            # Environment variables
└── README.md       # Project documentation

```


## Installation

### Backend (Go)

1. Clone the repo:  
   ```shell
   git clone https://github.com/Sayak-halder/Todo_Go.git
   ```
2. Navigate to `server/`:  
   `cd server`
3. Set up `.env` with MongoDB credentials.
4. Install dependencies:  
   ```shell
   go mod tidy
   ```
5. Start server:  
   ```shell
   air
   ```

### Frontend (React)

1. Navigate to `client/`:  
   `cd client`
2. Install dependencies:  
   ```shell
   npm install
   ```
3. Run development server:  
   ```shell
   npm run dev
   ```

## Frontend Overview

- **Navbar**: Includes a dark mode toggle using Chakra UI.
- **TodoForm**: Form for adding new tasks, integrated with `react-query` to handle POST requests.
- **TodoList**: Displays tasks from the backend, showing loading states and an empty state when no tasks are available.
- **TodoItem**: Represents each task with options to mark complete or delete, with real-time updates.

## Backend API Endpoints

| Method   | Endpoint          | Description               |
| -------- | ----------------- | ------------------------- |
| `GET`    | `/api/todos`       | Fetch all todos           |
| `POST`   | `/api/todos`       | Create a new todo         |
| `PATCH`  | `/api/todos/:id`   | Update a todo's status    |
| `DELETE` | `/api/todos/:id`   | Delete a todo             |
