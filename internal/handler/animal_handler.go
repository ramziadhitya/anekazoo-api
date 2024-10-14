package handler

import (
	"anekazoo-api/internal/domain"
	"anekazoo-api/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AnimalHandler struct {
	service *service.AnimalService
}

func NewAnimalHandler(service *service.AnimalService) *AnimalHandler {
	return &AnimalHandler{service: service}
}

func (h *AnimalHandler) GetAllAnimals(w http.ResponseWriter, r *http.Request) {
	animals, err := h.service.GetAllAnimals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(animals) == 0 {
		http.Error(w, "no animals found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(animals)
}

func (h *AnimalHandler) GetAnimalByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid animal ID", http.StatusBadRequest)
		return
	}

	animal, err := h.service.GetAnimalByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(animal)
}

func (h *AnimalHandler) CreateAnimal(w http.ResponseWriter, r *http.Request) {
	var animal domain.Animal
	if err := json.NewDecoder(r.Body).Decode(&animal); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateAnimal(animal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(animal)
}

func (h *AnimalHandler) UpdateAnimal(w http.ResponseWriter, r *http.Request) {
	var animal domain.Animal
	if err := json.NewDecoder(r.Body).Decode(&animal); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateAnimal(animal); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(animal)
}

func (h *AnimalHandler) DeleteAnimal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid animal ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteAnimal(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
