package repository

import (
	"anekazoo-api/internal/domain"
	"database/sql"
	"errors"
)

type AnimalRepository struct {
	db *sql.DB
}

func NewAnimalRepository(db *sql.DB) *AnimalRepository {
	return &AnimalRepository{db: db}
}

func (r *AnimalRepository) GetAllAnimals() ([]domain.Animal, error) {
	rows, err := r.db.Query("SELECT id, name, class, legs FROM animals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []domain.Animal
	for rows.Next() {
		var animal domain.Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs); err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}

func (r *AnimalRepository) GetAnimalByID(id int) (*domain.Animal, error) {
	var animal domain.Animal
	err := r.db.QueryRow("SELECT id, name, class, legs FROM animals WHERE id = ?", id).
		Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("animal not found")
		}
		return nil, err
	}
	return &animal, nil
}

func (r *AnimalRepository) CreateAnimal(animal domain.Animal) error {
	_, err := r.db.Exec("INSERT INTO animals (name, class, legs) VALUES (?, ?, ?)", animal.Name, animal.Class, animal.Legs)
	return err
}

func (r *AnimalRepository) UpdateAnimal(animal domain.Animal) error {
	result, err := r.db.Exec("UPDATE animals SET name = ?, class = ?, legs = ? WHERE id = ?", animal.Name, animal.Class, animal.Legs, animal.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("animal not found")
	}

	return nil
}

func (r *AnimalRepository) DeleteAnimal(id int) error {
	result, err := r.db.Exec("DELETE FROM animals WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("animal not found")
	}

	return nil
}
