import React, { Component } from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";

class RegisterForm extends Component {
  render() {
    return (
      <div className="LoginForm">
        <h1>Register</h1>
        <Form>
          <Form.Group controlId="test">
            <Form.Label>Username</Form.Label>
            <Form.Control type="email" placeholder="Username" />
          </Form.Group>
          <Form.Group>
            <Form.Label>Password</Form.Label>
            <Form.Control type="password" placeholder="Password" />
          </Form.Group>
          <Button variant="primary" type="submit">
            Log in
          </Button>
        </Form>
      </div>
    );
  }
}

export default RegisterForm;
