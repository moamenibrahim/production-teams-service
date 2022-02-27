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
	var response models.TeamResponse
	fmt.Println("Getting teams...")

	rows, err := db.Query("SELECT * FROM teams")
	if err != nil {
		utils.CheckErr(w, r, err)
		return
	}

	var teams []models.Team
	for rows.Next() {
		var teamID string
		var teamName string
		var teamType string

		err = rows.Scan(&teamID, &teamName, &teamType)
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

		teams = append(teams, models.Team{ID: teamID, Name: teamName, TeamType: teamType})
	}

	response = models.TeamResponse{Type: "success", Data: teams}
	json.NewEncoder(w).Encode(response)
}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.FormValue("teamid")
	teamName := r.FormValue("teamname")
	teamType := r.FormValue("teamtype")

	var response = models.TeamResponse{}
	if teamID == "" || teamName == "" {
		response = models.TeamResponse{Type: "error", Message: "You are missing teamID or teamName parameter."}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println("Inserting new team with ID: " + teamID + " and name: " + teamName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO teams(id, name, team_type) VALUES($1, $2, $3) returning id;", teamID, teamName, teamType).Scan(&lastInsertID)
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

		data := []models.Team{}
		data = append(data, models.Team{ID: teamID, Name: teamName, TeamType: teamType})
		response = models.TeamResponse{Type: "success", Data: data, Message: "The team has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teamID := params["teamid"]
	var response = models.TeamResponse{}
	if teamID == "" {
		response = models.TeamResponse{Type: "error", Message: "You are missing teamID parameter."}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println("Deleting team from DB")

		result, err := db.Exec("DELETE FROM teams where teamID = $1", teamID)
		rows, _ := result.RowsAffected()
		if rows == 0 {
			http.Error(w, http.StatusText(http.StatusNotFound),
				http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

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
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println("Getting team from DB")

		var ID string
		var teamName string
		var teamType string
		err := db.QueryRow("SELECT id, name, team_type FROM teams WHERE name=$1", teamID).Scan(&ID, &teamName, &teamType)
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

		teams := []models.Team{}
		teams = append(teams, models.Team{ID: ID, Name: teamName, TeamType: teamType})
		response = models.TeamResponse{Type: "success", Data: teams}
	}
	json.NewEncoder(w).Encode(response)
}
