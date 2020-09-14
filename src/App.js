import React, { Component } from "react";
import { hot } from "react-hot-loader";

import LoginForm from "./components/LoginForm";
import RegisterForm from "./components/RegisterForm";

import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

class App extends Component {
  render() {
    return (
      <div className="App">
        <Container>
          <Row>
            <Col>
              <LoginForm />
            </Col>
            <Col>
              <RegisterForm />
            </Col>
          </Row>
        </Container>
      </div>
    );
  }
}

export default hot(module)(App);
