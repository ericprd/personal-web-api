package controllers

import (
	"net/http"

	"github.com/ericprd/personal-web-api/helpers"
	"github.com/ericprd/personal-web-api/models"
	"github.com/ericprd/personal-web-api/services"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	ProjectService services.ProjectService
}

func NewProject(projectService services.ProjectService) ProjectController {
	return ProjectController{
		ProjectService: projectService,
	}
}

func (p *ProjectController) CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		helpers.ResponseHandler(c, err.Error(), http.StatusBadRequest)
		return
	}

	newProject := models.Project {
		Title: project.Title,
		Description: project.Description,
		ImageUrl: project.ImageUrl,
		Tags: project.Tags,
		Stacks: project.Stacks,
	}
	if err := p.ProjectService.CreateProject(&newProject); err != nil {
		helpers.ResponseHandler(c, err.Error(), http.StatusBadGateway)
		return
	}

	helpers.ResponseHandler(c, newProject, http.StatusCreated)

}

func (p *ProjectController) GetProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "message": "Hello" })
}

func (p *ProjectController) GetAllProjects(c *gin.Context) {
	projects, err := p.ProjectService.GetAllProjects()
	if err != nil {
		helpers.ResponseHandler(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseHandler(c, projects, http.StatusOK)

}

func (p *ProjectController) UpdateProject(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{ "message": "Hello" })
}

func (p *ProjectController) DeleteProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "message": "Hello" })
}

func (p *ProjectController) CreateProjectRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/projects")
	route.GET("/", p.GetAllProjects)
	route.POST("/", p.CreateProject)
	route.GET("/:title", p.GetProject)
	route.PUT("/:title", p.UpdateProject)
	route.DELETE("/:title", p.DeleteProject)
}
