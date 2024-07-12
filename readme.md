
# Go Todo Application

  

This a todo app made using Golang. Follow the instructions below to clone the repository and run the application.

  

## Prerequisites

  

Ensure you have the following installed on your system:

- [Go](https://golang.org/dl/) (version 1.20 or higher)
- [Git](https://git-scm.com/)
- [Node.js](https://nodejs.org/en) (version 18.17.0 or higher)

  

## Getting Started

  

### Clone the Repository

  

First, clone the repository using Git:

 
```sh
git  clone  https://github.com/Soumyadas15/go-todo
```

Move to project directory
```sh
cd go-todo
```

### Running the backend

Move to 'backend' directory
```sh
cd backend
```

And run to install the go dependencies
```sh
go mod tidy
```

After that, paste this command to build the app and create executable
```sh
go build -tags netgo -ldflags '-s -w' -o app
```
After that, paste this command as it is
```sh
./app --databaseURI="cleo.centralindia.cloudapp.azure.com:9042"
```

Wait for a while to see these logs in the console
```sh
2024/07/12 12:39:46 Db connected
2024/07/12 12:39:46 Starting server on port 8080
```

After running, you visit http://localhost:8080/swagger/index.html#/todo to see a detailed API documentation

### Running the frontend locally

After running the backend, navigate to the project directory using
```sh
cd ..
```
Then navigate to the frontend directory
```sh
cd frontend
```

In the root directory, create a file called .env and paste these and save
```sh
BACKEND_URL="http://localhost:8080"
NEXT_PUBLIC_BACKEND_URL="http://localhost:8080"
NEXTAPP_URL="http://localhost:8080"
```
After that, run the command to install dependencies
```sh
npm install
```
And finally run the app using
```sh
npm run dev
```

Visit http://localhost:3000 on your browser to see it running.

## Backend structure

### Schema diagram

![image](./db%20schema.png)

### APIs:

To get extensive details of the apis, visit http://localhost:8080/swagger/index.html#/todo after running the backend server

### Authentication

When a user logs in via the `/api/auth/login` endpoint, a JWT token is returned. This token is then stored in cookies on the frontend. For every subsequent request, the JWT token is retrieved from the cookies and sent along with the request.

I used bcrypt to hash passwords before storing them in the database.

### Todo table

To achieve database level pagination, I had to make certain changes in the database schema. I used `user_id` as partition key and `id` as clustering key. 

### Additional

I hosted the database container on an Azure virtual machine to avoid the hassle of creating containers before running the app.