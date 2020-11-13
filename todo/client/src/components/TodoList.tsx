import { IconButton, List, ListItem, ListItemSecondaryAction, ListItemText } from '@material-ui/core';
import { Delete } from '@material-ui/icons';
import React from 'react';

const todos = [
  "TODO 1",
  "TODO 2"
]

const TodoList = () => {
  return (
    <div>
      <List dense={true}>
        {todos.map(todo => (
          <ListItem key={todo}>
            <ListItemText
              primary={todo}
            />
            <ListItemSecondaryAction>
              <IconButton>
                <Delete />
              </IconButton>
            </ListItemSecondaryAction>
          </ListItem>
        ))}
      </List>
    </div>
  );
};

export default TodoList;