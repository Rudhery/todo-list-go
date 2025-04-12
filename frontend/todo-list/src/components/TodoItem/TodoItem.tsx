import React, { useState } from "react";
import { Todo } from "../../types/todo";
import "./TodoItem.css";

interface TodoItemProps {
  todo: Todo;
  onUpdate: (todo: Todo) => void;
  onDelete: (id: string) => void;
}

const TodoItem: React.FC<TodoItemProps> = ({ todo, onUpdate, onDelete }) => {
  const [isEditing, setIsEditing] = useState(false);
  const [editValue, setEditValue] = useState(todo.title);

  const handleToggleDone = () => {
    onUpdate({ ...todo, done: !todo.done });
  };

  const handleEdit = () => {
    setIsEditing(true);
  };

  const handleSave = () => {
    if (editValue.trim()) {
      onUpdate({ ...todo, title: editValue });
      setIsEditing(false);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      handleSave();
    } else if (e.key === "Escape") {
      setIsEditing(false);
      setEditValue(todo.title);
    }
  };

  return (
    <div className={`todo-item-container ${todo.done ? "done" : ""}`}>
      <input
        className="todo-checkbox"
        type="checkbox"
        checked={todo.done}
        onChange={handleToggleDone}
      />

      {isEditing ? (
        <input
          className="edit-input"
          type="text"
          value={editValue}
          onChange={(e) => setEditValue(e.target.value)}
          onBlur={handleSave}
          onKeyDown={handleKeyDown}
          autoFocus
        />
      ) : (
        <span
          className={`todo-title ${todo.done ? "done" : ""}`}
          onClick={handleEdit}
        >
          {todo.title}
        </span>
      )}

      <div className="todo-actions">
        <button onClick={handleEdit} disabled={todo.done} title="Editar tarefa">
          Editar
        </button>
        <button onClick={() => onDelete(todo.id)} title="Excluir tarefa">
          Excluir
        </button>
      </div>
    </div>
  );
};

export default TodoItem;
