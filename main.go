package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rusdiy/amezink_task/config/database"
	"github.com/rusdiy/amezink_task/controller"
	"github.com/rusdiy/amezink_task/repository"
	"github.com/rusdiy/amezink_task/service"
)

func main() {
	// Initializing database connection
	var dbase database.Database
	err := database.MySQLConn(&dbase)
	if err != nil {
		fmt.Println(err)
	}

	// Initializing The Repository
	recordRepository := repository.RecordRepository{
		DB: dbase,
	}

	// Initializing the Service
	recordService := service.RecordService{
		RecordRepo: recordRepository,
	}

	// Initializing the Controller
	recordController := controller.RecordController{
		RecordService: recordService,
	}

	// Initializing the Mux Router
	router := mux.NewRouter()

	// Routing the Endpoint to the Controller
	router.HandleFunc("/marks", recordController.GetMarks).Methods("GET")

	// Console Log and Start Listening
	fmt.Println("Listening to http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
