# Project Task Management APP
A RESTful API Project for project task management

## Installation & Run

Before running API server, you should set the database config on [config.go](go-rest-todo/config/config.go)
```go
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "guest",
			Password: "Guest0000!",
			Name:     "todoapp",
			Charset:  "utf8",
		},
	}
}
```

```bash
# Build and Run
cd go-todo-rest-api-example
go build
./go-todo-rest-api-example

# API Endpoint : http://127.0.0.1:3000
```

## Structure
```
├── app
│   ├── app.go
│   ├── handler          // API handlers
│   │   ├── common.go    // Common response functions
│   │   ├── projects.go  // APIs for Project model
│   │   └── tasks.go     // APIs for Task model
│   └── model
│       └── model.go     // Models for application
├── config
│   └── config.go        // Configuration
└── main.go
```

## API

#### /projects
* `GET` : Get all projects
* `POST` : Create a new project
```
#Sample Request Body
{
  "Title": "ProjectTest001",
  "Archived": true,
  "Tasks": [
    {
      "Title": "Testtask001",
      "Priority": "01",
      "Done": true,
      "ProjectID": 1
    }
  ]
}
```

#### /projects/:title
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project

#### /projects/:title/archive
* `PUT` : Archive a project
* `DELETE` : Restore a project 

#### /projects/:title/tasks
* `GET` : Get all tasks of a project
* `POST` : Create a new task in a project

#### /projects/:title/tasks/:id
* `GET` : Get a task of a project
* `PUT` : Update a task of a project
* `DELETE` : Delete a task of a project

#### /projects/:title/tasks/:id/complete
* `PUT` : Complete a task of a project
* `DELETE` : Undo a task of a project

