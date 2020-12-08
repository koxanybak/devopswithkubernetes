import axios from "axios";
import { Todo } from "../reducers/todoReducer";

export const addTodo = (todo: Pick<Todo, "name">) => {
  return axios.post("/api/todos/", todo).then(res => res.data)
}

export const putTodo = (todo: Todo) => {
  return axios.put("/api/todos/", todo).then(res => res.data)
}

export const getTodos = () => {
  return axios.get("/api/todos/").then(res => res.data)
}