import React, { useState, useEffect } from "react";
import { Todo } from "../../types/todo";
import { todoService } from "../../services/todoService";
import TodoItem from "../TodoItem/TodoItem";
import TodoForm from "../TodoForm/TodoForm";
import "./TodoList.css";

const TodoList: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Carregar tarefas ao iniciar o componente
  useEffect(() => {
    fetchTodos();
  }, []);

  // Buscar todas as tarefas
  const fetchTodos = async () => {
    setLoading(true);
    try {
      const data = await todoService.getTodos();
      // Garante que data é sempre um array, mesmo se a API retornar null
      setTodos(Array.isArray(data) ? data : []);
      setError(null);
    } catch (err) {
      setError("Erro ao carregar tarefas. Tente novamente mais tarde.");
      console.error(err);
      // Garante que todos é sempre um array, mesmo em caso de erro
      setTodos([]);
    } finally {
      setLoading(false);
    }
  };

  // Adicionar nova tarefa
  const handleAddTodo = async (title: string) => {
    try {
      const newTodo = await todoService.addTodo(title);
      if (newTodo) {
        setTodos([...todos, newTodo]);
      } else {
        // Se a API retornar null, mostra um erro mas mantém a lista atual
        setError("Erro ao adicionar tarefa. Tente novamente.");
        setTimeout(() => setError(null), 3000); // Remove a mensagem após 3 segundos
      }
    } catch (err) {
      setError("Erro ao adicionar tarefa. Tente novamente.");
      console.error(err);
      setTimeout(() => setError(null), 3000);
    }
  };

  // Atualizar tarefa existente
  const handleUpdateTodo = async (updatedTodo: Todo) => {
    try {
      const result = await todoService.updateTodo(updatedTodo);
      if (result) {
        setTodos(
          todos.map((todo) => (todo.id === updatedTodo.id ? updatedTodo : todo))
        );
      } else {
        // Se a API retornar null, mantém o estado atual e mostra um erro
        setError("Erro ao atualizar tarefa. Tente novamente.");
        setTimeout(() => setError(null), 3000); // Remove a mensagem após 3 segundos
      }
    } catch (err) {
      setError("Erro ao atualizar tarefa. Tente novamente.");
      console.error(err);
      setTimeout(() => setError(null), 3000);
    }
  };

  // Excluir tarefa
  const handleDeleteTodo = async (id: string) => {
    try {
      const success = await todoService.deleteTodo(id);
      if (success) {
        setTodos(todos.filter((todo) => todo.id !== id));
      } else {
        // Se a API retornar false, mantém o estado atual e mostra um erro
        setError("Erro ao excluir tarefa. Tente novamente.");
        setTimeout(() => setError(null), 3000); // Remove a mensagem após 3 segundos
      }
    } catch (err) {
      setError("Erro ao excluir tarefa. Tente novamente.");
      console.error(err);
      setTimeout(() => setError(null), 3000);
    }
  };

  return (
    <div className="todo-list-container">
      <div className="todo-list-header">
        <h1>Lista de Tarefas</h1>
      </div>

      <TodoForm onAddTodo={handleAddTodo} />

      <div className="todo-list-content">
        {loading ? (
          <p>Carregando tarefas...</p>
        ) : error ? (
          <p>{error}</p>
        ) : todos.length > 0 ? (
          todos.map((todo) => (
            <TodoItem
              key={todo.id}
              todo={todo}
              onUpdate={handleUpdateTodo}
              onDelete={handleDeleteTodo}
            />
          ))
        ) : (
          <p className="empty-message">
            Nenhuma tarefa encontrada. Adicione uma nova tarefa!
          </p>
        )}
      </div>
    </div>
  );
};

export default TodoList;
