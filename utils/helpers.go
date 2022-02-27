package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lithammer/shortuuid"

	"github.com/Shyp/go-dberror"
	pq "github.com/lib/pq"
)

func CheckErr(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		if pgerr, ok := err.(*pq.Error); ok {
			dberr := dberror.GetError(pgerr)
			switch e := dberr.(type) {
			case *dberror.Error:
				// TODO: Correctly map postgres errors to HTTP response errors
				if strings.Contains(e.Error(), "does not exist") {
					http.Error(w, e.Error(), http.StatusNotFound)
				} else if strings.Contains(e.Error(), "already exist") {
					http.Error(w, e.Error(), http.StatusConflict)
				}
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(err)
	}
}

func GenShortUUID() string {
	return shortuuid.New()
}
