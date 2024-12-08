# Golang Chat

Final Project Mata Kuliah Pemrograman Berbasis Kerangka Kerja

## Architecture

1. Domain Driven Design

   ![software architecture](./assets/architecture.png)

2. Hub

   ![chat-hub](./assets/join_room.jpg)
   ![hub-arch](./assets/hub_architecture.jpg)

## How to use

Requirements: Go, Node, PostgreSQL

1. Clone the repo

   ```bash
   git clone https://github.com/firaniaputri23/GolangFInalProject.git
   cd GolangFInalProject
   ```

2. Set up environment variables and run the server

   Copy .env.example to .env

   ```bash
   cd server
   cp .env.example .env
   ```

   ...and then edit `.env`

   Run the server

   ```bash
   go run cmd/main.go
   ```

3. Run the client with another terminal

   ```bash
   cd client
   npm run dev
   ```

4. Access `http://localhost:3000` with your browser
