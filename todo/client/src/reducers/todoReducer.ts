import { Dispatch } from "redux"
import { addTodo, getTodos } from "../services/todos"

export type Todo = {
  name: string;
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
type InitializeAction = {
  type: "INITIALIZE";
  payload: {
    todos: Todo[];
  }
}

export type TodoAction = CreateTodoAction | InitializeAction

export const todoReducer = (state: State = { todos: null }, action: TodoAction): State => {
  switch (action.type) {
    case "INITIALIZE":
      return { ...state, todos: action.payload.todos }
    case "CREATE_TODO":
      const todos = state.todos ? state.todos : []
      return { ...state, todos: todos.concat(action.payload.todo) }
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