package services

import (
	"context"
	"errors"
	"personal-web-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectImpl struct {
	projectCollection *mongo.Collection
	ctx context.Context
}

func NewProjectService(projectCollection *mongo.Collection, ctx context.Context) ProjectService {
	return &ProjectImpl {
		projectCollection: projectCollection,
		ctx: ctx,
	}
}

func (p *ProjectImpl) CreateProject(project *models.Project) error {
	_, err := p.projectCollection.InsertOne(p.ctx, project)
	return err
}

func (p *ProjectImpl) GetProject(id *primitive.ObjectID) (*models.ResponseProject, error) {

	return nil, nil
}

func (p *ProjectImpl) GetAllProjects() ([]*models.ResponseProject, error) {
	var projects []*models.ResponseProject

	cursor, err := p.projectCollection.Find(p.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(p.ctx) {
		var project models.ResponseProject
		err := cursor.Decode(&project)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(p.ctx)
	if len(projects) == 0 {
		return nil, errors.New("No data found")
	}
	return projects, nil
}

func (p *ProjectImpl) UpdateProject(project *models.Project) error {
	return nil
}

func (p *ProjectImpl) DeleteProject(id *primitive.ObjectID) error {
	return nil
}
