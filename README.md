# to-go

Backend for a small ToDo App, using the [Gin Framework](https://github.com/gin-gonic/gin).

# Run
To build the Docker Image run `./build.sh`.  
After that you can bring the app up using `docker-compose up`

## API

The API consists of the following endpoints:

- GET `/tasks`  
Returns all tasks

- POST `/tasks`  
Creates a new Task

- GET `/tasks/:id`  
Returns a specific task by Id

- PUT `/tasks/:id`  
Updates a specific task by Id

- DELETE `/tasks/:id`  
Deletes a specific task by ID

## Tasks

A task is a JSON Object like this:

```json
{
	"id": "id",
	"description": "Task Description",
	"date": "0001-01-01T00:00:00Z",
	"progress": 100,
	"finished": false
}
```
