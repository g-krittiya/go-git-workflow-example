package services

import "main.go/models"

type Student interface {
	GetStudents() ([]models.Student, error)
}
