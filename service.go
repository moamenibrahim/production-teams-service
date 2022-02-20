package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"encoding/json"

	"production_service/migrate"
	"production_service/router"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am here! :)")
	json.NewEncoder(w).Encode("I am Here! :)")
}

func main() {
	// Migrate DB with data
	migrate.MigrateDB()

	// Init the mux router and endpoints
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", GetHome).Methods("GET")

	muxRouter.HandleFunc("/teams", router.GetTeams).Methods("GET")
	muxRouter.HandleFunc("/teams", router.CreateTeam).Methods("POST")
	muxRouter.HandleFunc("/teams/{teamid}", router.GetTeam).Methods("GET")
	muxRouter.HandleFunc("/teams/{teamid}", router.DeleteTeam).Methods("DELETE")

	// Serve the app
	fmt.Println("Server at 8080")
	handler := cors.Default().Handler(muxRouter)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
