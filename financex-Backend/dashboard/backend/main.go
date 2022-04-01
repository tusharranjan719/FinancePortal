package main

import (
	"flag"
	"fmt"
	"github.com/Baumanar/bill-split/backend/data"
	_ "github.com/lib/pq"
)

func main() {
	var flagvar bool
	flag.BoolVar(&flagvar, "demo", false, "run the backend with demo data")
	flag.Parse()
	var app App
	app.Initialize()
	if flagvar {
		// Reset Database
		fmt.Println("Running in demo mode")
		data.SetupDB()
		// Populate the databse
		PopulateDB()
	}

	app.SetRoutes()
	app.Run()

}
