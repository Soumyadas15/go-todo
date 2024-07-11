package main

import (
    "log"
    "net/http"
    "backend/router"
	"backend/db"
)

func main() {
    db.InitCluster();
    defer db.CloseCluster()

    
    r := router.Router()
    log.Fatal(http.ListenAndServe(":8080", r))
}