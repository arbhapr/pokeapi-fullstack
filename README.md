# My Pokemon Project

This repository contains a full-stack web application that allows users to catch and manage Pokémon.
The project is structured with two submodules: a Go backend and a ReactJS + Vite frontend.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Development](#development)
- [Docker](#docker)
- [License](#license)


## Prerequisites

Before setting up the project, ensure you have the following installed:

- [Git](https://git-scm.com/)
- [Go](https://golang.org/doc/install) (for backend)
- [Node.js](https://nodejs.org/) and [Yarn](https://yarnpkg.com/) (for frontend)
- [Docker](https://www.docker.com/get-started) (for containerization)

## Setup

1. **Clone the repository with submodules**:

   ```bash
   git clone --recurse-submodules <repository-url>
   cd my-pokemon
   ```

2. **Initialize and update submodules (if not done automatically)**:
    ```bash
    git submodule init
    git submodule update --recursive
    ```

3. **Backend Setup**:

   Navigate to the backend folder:

   ```bash
   cd backend
   ```

   Install Go dependencies:

    ```bash
    cp .env.example .env
    go mod tidy
    ```

4. **Frontend Setup**:
   
   Navigate to the frontend folder:

    ```bash
    cd ../frontend
    ```

   Install Node.js dependencies:

    ```bash
    cp .env.example .env
    yarn install
    ```

## Development
### Running Backend Locally
1. Navigate to the backend folder:

   ```bash
   cd ../backend
   ```

2. Start the Go server:

   ```bash
   go run main.go
   ```

The backend will run on **http://localhost:3000**.

### Running Frontend Locally
1. Navigate to the frontend folder:

   ```bash
   cd ../frontend
   ```

2. Start the React development server:

   ```bash
   yarn dev
   ```

The frontend will run on **http://localhost:5173** by default.

## Docker
### Building and Running with Docker Compose
To run both the backend and frontend using Docker:

1. Ensure Docker is running on your machine.

2. Build and start the containers:

   ```bash
   cp .env.example .env
   docker-compose up --build
   ```

   * The backend will be available at **http://localhost:3000**.
   * The frontend will be available at **http://localhost:80**.

### Stopping the Containers
To stop the running containers:

   ```bash
   docker-compose down
   ```

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

This project created by **[Arbha Pradana](https://linkedin.com/in/arbhapr)**, following the challenge of **PINTRO** Golang Technical Test.
