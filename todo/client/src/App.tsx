import { Grid } from '@material-ui/core';
import React from 'react';
import TodoForm from './components/TodoForm';
import TodoList from './components/TodoList';

const size = 6

const App = () => {
  return (
    <div>
      <Grid
        container
        spacing={2}
        direction="column"
      >
        <Grid item md={size}>
          <img src="/api/image" alt="daily" />
        </Grid>
        <Grid item md={size}>
          <TodoForm />
        </Grid>
        <Grid item md={4}>
          <TodoList />
        </Grid>
      </Grid>
    </div>
  );
};

export default App;