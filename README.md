# Todo List

A simple API driven todo list.

Frontend: React

Backend: Golang

Database: NONE(persisted in memory on the backend)


## GET /todos
Retrieve all tasks.

Returns 200 or 500.

## GET /todos/uuid
Retrieve the task with uuid.

Returns 200 or 500 or 404.

## POST /todos
Add task to the pending list. Provide task in the request body. Returns the uuid for the newly added task.

Returns 200.

## PATCH /todos/uuid
Toggles the status of the task with uuid.

Returns 200 or 404.

## DELETE /todos
Clears the entire task list.

Returns 200.

## DELETE /todos/uuid
Removes the task with uuid from the task list.

Returns 200 or 404.
