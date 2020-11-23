import { IconButton, List, ListItem, ListItemSecondaryAction, ListItemText } from '@material-ui/core';
import { Delete } from '@material-ui/icons';
import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../reducers';
import { initializeTodos } from '../reducers/todoReducer';

const TodoList = () => {
  const todos = useSelector((state: RootState) => state.todos.todos)
  const dispatch = useDispatch()
  useEffect(() => {
    dispatch(initializeTodos())
  }, [])

  return (
    <div>
      <List dense={true}>
        {todos ? todos.map(todo => (
          <ListItem key={todo.name}>
            <ListItemText
              primary={todo.name}
            />
            <ListItemSecondaryAction>
              <IconButton>
                <Delete />
              </IconButton>
            </ListItemSecondaryAction>
          </ListItem>
        )) : null}
      </List>
    </div>
  );
};

export default TodoList;