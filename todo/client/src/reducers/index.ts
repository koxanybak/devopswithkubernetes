import { applyMiddleware, combineReducers, createStore } from "redux";
import thunk from "redux-thunk";
import { todoReducer } from "./todoReducer";

const rootReducer = combineReducers({
  todos: todoReducer,
})

export const store = createStore(
  rootReducer,
  undefined,
  applyMiddleware(thunk)
)

export type RootState = ReturnType<typeof rootReducer>