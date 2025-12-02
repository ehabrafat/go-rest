-- name: ListTodos :many
SELECT 
* 
FROM 
    todos;

-- name: GetTodoById :one
SELECT * FROM todos where id = $1;
