package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"

	"production_service/migrate"
	"production_service/router"
)

func SetupHandlers() {
	// Init the mux router and endpoints
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", router.GetHome).Methods("GET")

	muxRouter.HandleFunc("/users", router.GetUsers).Methods("GET")
	muxRouter.HandleFunc("/users", router.CreateUser).Methods("POST")
	muxRouter.HandleFunc("/users/{username}", router.GetUser).Methods("GET")
	muxRouter.HandleFunc("/users/{userid}", router.DeleteUser).Methods("DELETE")

	muxRouter.HandleFunc("/hubs", router.GetHubs).Methods("GET")
	muxRouter.HandleFunc("/hubs", router.CreateHub).Methods("POST")
	muxRouter.HandleFunc("/hubs/{hubname}", router.GetHub).Methods("GET")
	muxRouter.HandleFunc("/hubs/{hubid}", router.DeleteHub).Methods("DELETE")

	muxRouter.HandleFunc("/teams", router.GetTeams).Methods("GET")
	muxRouter.HandleFunc("/teams", router.CreateTeam).Methods("POST")
	muxRouter.HandleFunc("/teams/{teamname}", router.GetTeam).Methods("GET")
	muxRouter.HandleFunc("/teams/{teamid}", router.DeleteTeam).Methods("DELETE")

	// Serve the app
	fmt.Println("Server at 8080")
	handler := cors.Default().Handler(muxRouter)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func main() {
	// Migrate DB with data
	migrate.MigrateDB()

	// Setup handlers and endpoints
	SetupHandlers()
}
