package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"database/sql"
)

var db *sql.DB

func main() {

	router := NewRouter()
	time.Sleep(time.Second * 10)
	fmt.Println("starting to connect to DB...")
	db = ConnectToPG("nest")
	fmt.Println("intermediate check")
	SetupTable(db, "listings")
	ReadSeed(db)
	fmt.Println("connected")
	log.Fatal(http.ListenAndServe(":8080", router))
}
