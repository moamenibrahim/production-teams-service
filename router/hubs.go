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

func GetHubs(w http.ResponseWriter, r *http.Request) {
	var response models.HubResponse
	fmt.Println("Getting hubs...")

	rows, err := db.Query("SELECT * FROM hubs")
	if err != nil {
		utils.CheckErr(w, r, err)
		return
	}

	var hubs []models.Hub
	for rows.Next() {
		var hubID string
		var hubName string
		var geoLocation string

		err = rows.Scan(&hubID, &hubName, &geoLocation)
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

		hubs = append(hubs, models.Hub{ID: hubID, Name: hubName, GeoLocation: geoLocation})
	}

	response = models.HubResponse{Type: "success", Data: hubs}
	json.NewEncoder(w).Encode(response)
}

func CreateHub(w http.ResponseWriter, r *http.Request) {
	hubID := r.FormValue("hubid")
	hubName := r.FormValue("hubname")
	geolocation := r.FormValue("geolocation")

	var response = models.HubResponse{}
	if hubID == "" || hubName == "" {
		response = models.HubResponse{Type: "error", Message: "You are missing hubID or hubName parameter."}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println("Inserting new hub with ID: " + hubID + " and name: " + hubName + " and geoLocation: " + geolocation)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO hubs(id, name, geo_location) VALUES($1, $2, $3) returning id;", hubID, hubName, geolocation).Scan(&lastInsertID)
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

		data := []models.Hub{}
		data = append(data, models.Hub{ID: hubID, Name: hubName, GeoLocation: geolocation})
		response = models.HubResponse{Type: "success", Data: data, Message: "The hub has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteHub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hubID := params["hubid"]
	var response = models.HubResponse{}
	if hubID == "" {
		response = models.HubResponse{Type: "error", Message: "You are missing hubID parameter."}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println("Deleting hub from DB")

		result, err := db.Exec("DELETE FROM hubs where id=$1", hubID)
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

		response = models.HubResponse{Type: "success", Message: "The hub has been deleted successfully!"}
	}
	json.NewEncoder(w).Encode(response)
}

func GetHub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hubID := params["hubid"]
	var response = models.HubResponse{}
	if hubID == "" {
		response = models.HubResponse{Type: "error", Message: "You are missing hubID parameter."}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println("Getting hub from DB")

		var ID string
		var hubName string
		var geoLocation string
		err := db.QueryRow("SELECT id, name, geo_location FROM hubs WHERE name=$1", hubName).Scan(&ID, &hubName, &geoLocation)
		if err != nil {
			utils.CheckErr(w, r, err)
			return
		}

		hubs := []models.Hub{}
		hubs = append(hubs, models.Hub{ID: ID, Name: hubName, GeoLocation: geoLocation})
		response = models.HubResponse{Type: "success", Data: hubs, Message: "The hub has been deleted successfully!"}
	}
	json.NewEncoder(w).Encode(response)
}
