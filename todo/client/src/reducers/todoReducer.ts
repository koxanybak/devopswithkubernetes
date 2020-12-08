import { Dispatch } from "redux"
import { addTodo, getTodos, putTodo } from "../services/todos"

export type Todo = {
  id: number;
  name: string;
  done: boolean;
}

export type State = {
  todos?: Todo[] | null;
}

// const CREATE_TODO = "CREATE_TODO"
// const INITIALIZE = "INITIALIZE"
type CreateTodoAction = {
  type: "CREATE_TODO";
  payload: {
    todo: Todo;
  }
}
type UpdateTodoAction = {
  type: "UPDATE_TODO";
  payload: {
    todo: Todo;
  }
}
type InitializeAction = {
  type: "INITIALIZE";
  payload: {
    todos: Todo[];
  }
}

export type TodoAction = CreateTodoAction | InitializeAction | UpdateTodoAction

export const todoReducer = (state: State = { todos: null }, action: TodoAction): State => {
  switch (action.type) {
    case "INITIALIZE":
      return { ...state, todos: action.payload.todos }
    case "CREATE_TODO":
      const todos = state.todos ? state.todos : []
      return { ...state, todos: todos.concat(action.payload.todo) }
    case "UPDATE_TODO":
      return {
        ...state,
        todos: state.todos ? state.todos.map(t => (
          t.id === action.payload.todo.id ? action.payload.todo : t
        )) : null,
      }
    default:
      return state
  }
}

export const createTodo = (name: string) => {
  return async (dispatch: Dispatch<TodoAction>) => {
    try {
      const newTodo = await addTodo({ name })
      dispatch({
        type: "CREATE_TODO",
        payload: {
          todo: newTodo,
        },
      })
    } catch (e) {
      console.error(e)
    }
  }
}

export const updateTodo = (todo: Todo) => {
  return async (dispatch: Dispatch<TodoAction>) => {
    try {
      const newTodo = await putTodo(todo)
      dispatch({
        type: "UPDATE_TODO",
        payload: {
          todo: newTodo,
        },
      })
    } catch (e) {
      console.error(e)
    }
  }
}

export const initializeTodos = () => {
  return async (dispatch: Dispatch<TodoAction>) => {
    try {
      const todos = await getTodos()
      dispatch({
        type: "INITIALIZE",
        payload: {
          todos
        }
      })
    } catch (e) {
      console.error(e)
    }
  }
}