package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"
)

func main() {
	// Initialize database connection
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	// Auto migrate the schema
	if err := dbConn.GetDB().AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("could not migrate database: %s", err)
	}

	// Initialize repositories, services and handlers
	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	// Initialize websocket hub
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	// Setup and start router
	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
