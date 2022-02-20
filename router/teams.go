package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"production_service/models"
	"production_service/utils"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting teams...")
	rows, err := db.Query("SELECT * FROM teams")
	utils.CheckErr(err)

	var teams []models.Team
	for rows.Next() {
		var id string
		var teamName string
		var teamType string

		err = rows.Scan(&id, &teamName, &teamType)
		utils.CheckErr(err)
		teams = append(teams, models.Team{ID: id, Name: teamName, TeamType: teamType})
	}

	var response = models.TeamResponse{Type: "success", Data: teams}
	json.NewEncoder(w).Encode(response)
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.FormValue("teamid")
	teamName := r.FormValue("teamname")
	geolocation := r.FormValue("geolocation")

	var response = models.TeamResponse{}
	if teamID == "" || teamName == "" {
		response = models.TeamResponse{Type: "error", Message: "You are missing teamID or teamName parameter."}
	} else {
		fmt.Println("Inserting team into DB")
		fmt.Println("Inserting new team with ID: " + teamID + " and name: " + teamName)
		var lastInsertID int
		err := db.QueryRow("INSERT INTO teams(teamID, teamName, geolocation) VALUES($1, $2, $3) returning id;", teamID, teamName, geolocation).Scan(&lastInsertID)
		utils.CheckErr(err)
		response = models.TeamResponse{Type: "success", Message: "The team has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teamID := params["teamid"]
	var response = models.TeamResponse{}
	if teamID == "" {
		response = models.TeamResponse{Type: "error", Message: "You are missing teamID parameter."}
	} else {
		fmt.Println("Deleting team from DB")
		_, err := db.Exec("DELETE FROM teams where teamID = $1", teamID)
		utils.CheckErr(err)
		response = models.TeamResponse{Type: "success", Message: "The team has been deleted successfully!"}
	}
	json.NewEncoder(w).Encode(response)
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teamID := params["teamid"]
	var response = models.TeamResponse{}
	if teamID == "" {
		response = models.TeamResponse{Type: "error", Message: "You are missing teamID parameter."}
	} else {
		fmt.Println("Deleting team from DB")
		_, err := db.Exec("SELECT $1 FROM teams", teamID)
		utils.CheckErr(err)
		response = models.TeamResponse{Type: "success", Message: "The team has been deleted successfully!"}
	}
	json.NewEncoder(w).Encode(response)
}
