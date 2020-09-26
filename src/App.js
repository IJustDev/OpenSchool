import React, { Component } from "react";
import { hot } from "react-hot-loader";

import LoginForm from "./components/LoginForm";
import RegisterForm from "./components/RegisterForm";
import Navigation from "./components/Navigation";

import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

class App extends Component {
  render() {
    return (
      <div className="App">
        <Navigation>
        </Navigation>
        <Container>
          <Row>
              <LoginForm />
          </Row>
        </Container>
      </div>
    );
  }
}

export default hot(module)(App);
