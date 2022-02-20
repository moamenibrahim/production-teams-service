package router

import (
	_ "github.com/lib/pq"
	"production_service/utils"
)

var db = utils.SetupDB()
