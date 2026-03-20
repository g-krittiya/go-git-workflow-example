package repositories

import "main.go/models"

type StudentRepository interface {
	GetAll() ([]models.Student, error)
}
