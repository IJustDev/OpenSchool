# Development

## Project structure
```
..
|-- CHANGELOG.md
|-- Dockerfile
|-- README.md
|-- api
|   |-- admin.go
|   |-- auth.go
|   |-- auth_test.go
|   |-- docs
|   |-- go.mod
|   |-- go.sum
|   |-- main.go
|   |-- models
|   |-- routes.go
|   |-- server.go
|   |-- token.go
|   `-- user_repository.go
|-- package-lock.json
|-- package.json
|-- public
|   `-- index.html
|-- resources
|   `-- logo.svg
|-- specification.md
|-- src
|   |-- App.js
|   |-- components
|   `-- index.js
`-- webpack.config.js

7 directories, 21 files
```

- The `api/` folder contains all the code for the backend that needs to be run in order to run the frontend.
- The `src/` folder consists of the frontend (react) code. Before you run it, you will need to install npm dependencies.
- The `public/` folder is used for the index.html and may be used for images in the future.
- The `resources/` folder is currently only meant for the repository resources.

## Setup
Set up a development environment by cloning the repository.

```sh
$ git clone git@github.com:IJustDev/OpenSchool.git
# or
$ git clone https://github.com/IJustDev/OpenSchool
```

### Run Backend
```sh
# Navigate to the api/ folder
$ cd api/
# And run it
$ go run .
```

### Run frontend
You will need to configure the backend with an SQL database.

```sh
# Install necessary dependencies
$ npm i
# Run webpack dev server
$ npm run dev
```
