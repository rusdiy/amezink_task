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
	dbase, err := database.MySQLConn()
	if err != nil {
		fmt.Println(err)
	}

	recordRepository := repository.RecordRepository{
		DB: *dbase,
	}

	recordService := service.RecordService{
		RecordRepo: recordRepository,
	}

	recordController := controller.RecordController{
		RecordService: recordService,
	}


	router := mux.NewRouter()
	router.HandleFunc("/marks", recordController.GetMarks).Methods("GET")
	
	fmt.Println("Listening to http://localhost:40004")
	log.Fatal(http.ListenAndServe("localhost:40004", router))
}