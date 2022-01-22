# Tototo Do

Tototo Do is a simple web based todo list application.

![Tototo Do](tototodo.gif)

By Tototo Do:

- You can add your todo item
- You can see the todo items that you added before

Table of Contents
========

- [Tototo Do](#tototo-do)
- [Table of Contents](#table-of-contents)
  - [Technologies](#technologies)
  - [Folder Structure](#folder-structure)
  - [Installation](#installation)
    - [Building up](#building-up)
    - [Starting](#starting)
  - [Deployment](#deployment)
  - [Running tests](#running-tests)
    - [Backend tests](#backend-tests)
    - [Frontend tests](#frontend-tests)
  - [Etymology](#etymology)

## Technologies

This app uses Vue on the frontend side, and Golang with Gorilla Mux router on the backend RESTful API. It uses MySQL as database. 

## Folder Structure

```shell
tototodo/
├── .circleci
│   └── config.yml
├── backend
│   ├── _config
│   │   └── config.go
│   ├── features
│   │   └── ...
│   ├── src
│   │   ├── database.go
│   │   └── handlers.go
│   ├── main_test.go
│   ├── main.go
│   └── ...
├── docker
│   ├── backend.dockerfile
│   └── frontend.dockerfile
├── frontend
│   ├── src
│   │   ├── components
│   │   │    └── ...
│   │   ├── App.vue
│   │   └── main.js
│   ├── public
│   │   ├── index.html
│   │   └── ...
│   ├── tests
│   │   └── ...
│   └── ...
├── docker-compose.yml
└── ...
```

## Installation

Tototo Do is a dockerized application. You can use docker-compose to build up all project. 

It uses a MySQL database that located in a container.

### Building up

```bash
docker-compose build
```

### Starting

```bash
docker-compose up
```

REST API port will be ``8000``, UI serving port will be ``8001`` and database port will be ``33066``.

## Deployment

Tototo Do uses Circle CI to CI/CD process. 

## Running tests

### Backend tests
You should use ``godog`` to run tests

### Frontend tests
You should use ``npm run test:e2e`` to run tests

## Etymology
"Tototo Do" name is derrived from the opening sequence of the [Beethoven's 5th Symphony](https://www.youtube.com/watch?v=_4IRMYuE1hI).