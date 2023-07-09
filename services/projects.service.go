package services

import (
	"personal-web-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectService interface {
	CreateProject(*models.Project) error
	GetProject(*primitive.ObjectID) (*models.ResponseProject, error)
	GetAllProjects() ([]*models.ResponseProject, error)
	UpdateProject(*models.Project) error
	DeleteProject(*primitive.ObjectID) error
}
