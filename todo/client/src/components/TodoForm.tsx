import React from 'react';
import { Formik } from "formik";
import { Button, FormControl, TextField } from '@material-ui/core';
import * as yup from "yup"
import { useDispatch } from 'react-redux';
import { createTodo } from '../reducers/todoReducer';

type FormValues = {
  name: string;
}

const initialValues: FormValues = {
  name: ""
}

const validationSchema = yup.object().shape({
  name: yup.string().required("Required TODO")
})

const TodoForm = () => {
  const dispatch = useDispatch()
  
  return (
    <div>
      <Formik
        initialValues={initialValues}
        validationSchema={validationSchema}
        onSubmit={(values, actions) => {
          console.log(values)
          dispatch(createTodo(values.name))
          actions.setSubmitting(false)
        }}
      >
        {({
          values,
          isSubmitting,
          handleSubmit,
          handleChange
        }) => (
          <form onSubmit={handleSubmit}>
            <FormControl>
              <TextField
                label="TODO name"
                name="name"
                value={values.name}
                onChange={handleChange}
                variant="outlined"
              />
            </FormControl>
            <Button
              type="submit"
              disabled={isSubmitting}
              variant="contained"
              color="primary"
            >
              Create TODO
            </Button>
          </form>
        )}
      </Formik>
    </div>
  );
};

export default TodoForm;