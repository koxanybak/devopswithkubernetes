import { IconButton, List, ListItem, ListItemSecondaryAction, ListItemText } from '@material-ui/core';
import { Delete, Done } from '@material-ui/icons';
import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../reducers';
import { initializeTodos, Todo as TodoType, updateTodo } from '../reducers/todoReducer';

type TodoProps = {
  todo: TodoType;
}

const Todo: React.FC<TodoProps> = ({ todo }) => {
  const dispatch = useDispatch()
  const toggleDone = () => {
    dispatch(updateTodo({ ...todo, done: !todo.done }))
  }

  return (
    <ListItem>
      <ListItemText
        primary={todo.name}
      />
      <ListItemText
        primary={todo.done ? "Done" : "Not done"}
      />
      <ListItemSecondaryAction>
        <IconButton onClick={toggleDone}>
          <Done />
        </IconButton>
        <IconButton>
          <Delete />
        </IconButton>
      </ListItemSecondaryAction>
    </ListItem>
  );
};

const TodoList = () => {
  const todos = useSelector((state: RootState) => state.todos.todos)
  const dispatch = useDispatch()
  useEffect(() => {
    dispatch(initializeTodos())
  }, [dispatch])

  return (
    <div>
      <List dense={true}>
        {todos && typeof todos !== "string" ? todos.map(todo => (
          <Todo key={todo.id} todo={todo} />
        )) : null}
      </List>
    </div>
  );
};

export default TodoList;