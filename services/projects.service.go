package services

import (
	"personal-web-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectService interface {
	CreateProject(*models.Project) error
	GetProject(*primitive.ObjectID) (*models.Project, error)
	GetAllProjects() ([]*models.Project, error)
	UpdateProject(*models.Project) error
	DeleteProject(*primitive.ObjectID) error
}
