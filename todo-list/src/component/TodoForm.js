import React from 'react'

function TodoForm(props) {
    const [value, setValue] = React.useState("");

    const handleSubmit = e => {
        e.preventDefault();
        if (!value) return;
        props.addTodo(value);
        setValue("");
    };

    const inputStyle = {
        width: "75%",
        boxSizing: "border-box"
    }

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                maxLength="25"
                className="input"
                value={value}
                style={inputStyle}
                onChange={event => setValue(event.target.value)}
            />
            <button type="submit">Add Todo</button>
        </form >
    );
}

export default TodoForm;