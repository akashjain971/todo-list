import React from 'react'

function Todo(props) {

    const { todo, index, completeTodo, removeTodo } = props;

    return (
        <div className="todo">
            <p style={{ textDecoration: todo.isCompleted ? "line-through" : "" }}>{todo.text}</p>
            <div>
                <button onClick={() => completeTodo(index)}>Complete</button>
                <button onClick={() => removeTodo(index)}>X</button>
            </div>
        </div>
    );
}

export default Todo;