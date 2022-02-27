package router

import (
	"fmt"
	"net/http"
	"encoding/json"
	
	_ "github.com/lib/pq"
	"production_service/utils"
)

var db = utils.SetupDB()

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am here! :)")
	json.NewEncoder(w).Encode("I am Here! :)")
}
