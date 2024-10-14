package service

import (
	"anekazoo-api/internal/domain"
	"anekazoo-api/internal/repository"
)

type AnimalService struct {
	repo *repository.AnimalRepository
}

func NewAnimalService(repo *repository.AnimalRepository) *AnimalService {
	return &AnimalService{repo: repo}
}

func (s *AnimalService) GetAllAnimals() ([]domain.Animal, error) {
	return s.repo.GetAllAnimals()
}

func (s *AnimalService) GetAnimalByID(id int) (*domain.Animal, error) {
	return s.repo.GetAnimalByID(id)
}

func (s *AnimalService) CreateAnimal(animal domain.Animal) error {
	return s.repo.CreateAnimal(animal)
}

func (s *AnimalService) UpdateAnimal(animal domain.Animal) error {
	return s.repo.UpdateAnimal(animal)
}

func (s *AnimalService) DeleteAnimal(id int) error {
	return s.repo.DeleteAnimal(id)
}
