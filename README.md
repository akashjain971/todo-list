# todo-list

A simple API driven task list.

The tasks are stored in memory and does not persist in any database.

Maintains tasks as complete or pending.

## GET /todos
Retrieve all items of the list as "complete" or "pending".

## GET /todos/\<task\>
Retrieve the status of the \<task\>

## POST /todos/\<task\>
Add task to the pending list

## PATCH /todos/\<task\>
Flips the status of the \<task\>

## DELETE /todos
Clears the task list

## DELETE /todos/\<task\>
Removes the \<task\> from the task list
