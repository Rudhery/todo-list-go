import React, { useState } from "react";
import "./TodoForm.css";

interface TodoFormProps {
  onAddTodo: (title: string) => void;
}

const TodoForm: React.FC<TodoFormProps> = ({ onAddTodo }) => {
  const [title, setTitle] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (title.trim()) {
      onAddTodo(title.trim());
      setTitle("");
    }
  };

  return (
    <form className="todo-form-container" onSubmit={handleSubmit}>
      <input
        className="todo-input"
        type="text"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Adicionar nova tarefa..."
        autoFocus
      />
      <button className="todo-button" type="submit" disabled={!title.trim()}>
        Adicionar
      </button>
    </form>
  );
};

export default TodoForm;
