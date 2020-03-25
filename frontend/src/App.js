import React from "react";
import Todo from "./component/Todo";
import TodoForm from "./component/TodoForm";
import axios from "axios"
import "./App.css";

function App() {
  const [todos, setTodos] = React.useState([]);
  const url = "http://localhost:9999/todos/";

  React.useEffect(() => {
    axios.get(url).then(function (response) {
      setTodos(response.data);
      console.log(response.data);
    }).catch(function (error) {
      console.log(error);
    });
  }, []);

  const addTodo = todo => {
    axios.post(url, todo)
      .then(function (response) {
        console.log(response.data);
        var task = {
          id: response.data,
          task: todo,
          complete: false
        }
        const newTodos = [...todos, task];
        setTodos(newTodos);
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  const completeTodo = index => {
    axios.patch(url + index)
      .then(() => {
        const newTodos = [...todos];
        newTodos.forEach((todo) => {
          if (todo.id === index) {
            todo.complete = !todo.complete;
          }
        });
        setTodos(newTodos);
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  const removeTodo = index => {
    axios.delete(url + index)
      .then(() => {
        const newTodos = [...todos].filter((todo) => todo.id !== index);
        setTodos(newTodos);
      }).catch((error) => {
        console.log(error);
      });
  };

  return (
    <div className="app">
      <div className="todo-list">
        {todos.map((todo) => (
          <Todo
            key={todo.id}
            index={todo.id}
            todo={todo}
            completeTodo={completeTodo}
            removeTodo={removeTodo}
          />
        ))}
        <TodoForm addTodo={addTodo} />
      </div>
    </div>
  );
}

export default App;