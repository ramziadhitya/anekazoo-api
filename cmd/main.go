package main

import (
	"anekazoo-api/internal/handler"
	"anekazoo-api/internal/repository"
	"anekazoo-api/internal/service"
	"anekazoo-api/pkg/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Setup database
	db, err := db.ConnectMySQL()
	if err != nil {
		log.Fatal("Failed to connect to MySQL: ", err)
	}
	defer db.Close()

	animalRepo := repository.NewAnimalRepository(db)
	animalService := service.NewAnimalService(animalRepo)

	animalHandler := handler.NewAnimalHandler(animalService)

	router := mux.NewRouter()
	router.HandleFunc("/v1/animals", animalHandler.GetAllAnimals).Methods("GET")
	router.HandleFunc("/v1/animals/{id:[0-9]+}", animalHandler.GetAnimalByID).Methods("GET")
	router.HandleFunc("/v1/animals", animalHandler.CreateAnimal).Methods("POST")
	router.HandleFunc("/v1/animals", animalHandler.UpdateAnimal).Methods("PUT")
	router.HandleFunc("/v1/animals/{id:[0-9]+}", animalHandler.DeleteAnimal).Methods("DELETE")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
