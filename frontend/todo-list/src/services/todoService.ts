import { Todo } from "../types/todo";

const API_URL = "http://192.168.0.181:8080/api";

export const todoService = {
  // Buscar todas as tarefas
  async getTodos(): Promise<Todo[]> {
    try {
      const response = await fetch(`${API_URL}/todos`, {
        credentials: "include",
      });
      if (!response.ok) {
        throw new Error("Erro ao buscar tarefas");
      }
      const data = await response.json();
      // Garante que sempre retornamos um array, mesmo se a API retornar null
      return Array.isArray(data) ? data : [];
    } catch (error) {
      console.error("Erro ao buscar tarefas:", error);
      return [];
    }
  },

  // Adicionar uma nova tarefa
  async addTodo(title: string): Promise<Todo | null> {
    try {
      const response = await fetch(`${API_URL}/todos`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title,
          done: false,
          created_at: new Date().toISOString(),
        }),
      });

      if (!response.ok) {
        throw new Error("Erro ao adicionar tarefa");
      }

      const data = await response.json();
      // Verifica se a resposta é válida
      if (!data || typeof data !== "object" || !data.id) {
        console.error("Resposta inválida ao adicionar tarefa:", data);
        return null;
      }

      return data;
    } catch (error) {
      console.error("Erro ao adicionar tarefa:", error);
      return null;
    }
  },

  // Atualizar uma tarefa existente
  async updateTodo(todo: Todo): Promise<Todo | null> {
    try {
      const response = await fetch(`${API_URL}/todos/${todo.id}`, {
        method: "PUT",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(todo),
      });

      if (!response.ok) {
        throw new Error("Erro ao atualizar tarefa");
      }

      const data = await response.json();
      // Verifica se a resposta é válida
      if (!data || typeof data !== "object" || !data.id) {
        console.error("Resposta inválida ao atualizar tarefa:", data);
        return null;
      }

      return data;
    } catch (error) {
      console.error("Erro ao atualizar tarefa:", error);
      return null;
    }
  },

  // Excluir uma tarefa
  async deleteTodo(id: string): Promise<boolean> {
    try {
      const response = await fetch(`${API_URL}/todos/${id}`, {
        method: "DELETE",
        credentials: "include",
      });

      return response.ok;
    } catch (error) {
      console.error("Erro ao excluir tarefa:", error);
      return false;
    }
  },
};
