package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ericprd/personal-web-api/controllers"
	"github.com/ericprd/personal-web-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
)

var (
	server 							*gin.Engine
	projectService 			services.ProjectService
	ProjectController 	controllers.ProjectController
	ctx 								context.Context
	projectCollection 	*mongo.Collection
	mongoClient 				*mongo.Client
	err 								error
	_ = godotenv.Load()
)

func init() {
	ctx = context.TODO()
	DB_URI := os.Getenv("DB_URL")
	mongoConn := options.Client().ApplyURI(DB_URI)
	mongoClient, err = mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB Connected")
	projectCollection = mongoClient.Database("personal-web").Collection("projects")
	projectService = services.NewProjectService(projectCollection, ctx)
	ProjectController = controllers.NewProject(projectService)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)
	basePath := server.Group("/v1")
	ProjectController.CreateProjectRoutes(basePath)

	server.Run()
}
