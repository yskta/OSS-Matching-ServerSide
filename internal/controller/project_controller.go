package controller

import (
	"OSS-Matching-ServerSide/internal/controller/dto"
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProjectController struct {
	projectService service.ProjectService
}

func NewProjectController(ps service.ProjectService) *ProjectController {
	return &ProjectController{
		projectService: ps,
	}
}

func (c *ProjectController) Create(ctx echo.Context) error {
	req := new(dto.CreateProjectRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	project, err := c.projectService.Create(req.GithubRepoID, req.Name, req.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// DBモデルからDTOのレスポンス型に変換
	response := &dto.ProjectResponse{
		ID:           project.ID.String(),
		GithubRepoID: project.GithubRepoID,
		Name:         project.Name,
		Description:  project.Description.String,
		IsActive:     project.IsActive.Bool,
		CreatedAt:    project.CreatedAt.Time,
		UpdatedAt:    project.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (c *ProjectController) Get(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id format")
	}

	project, err := c.projectService.Get(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// DBモデルからDTOのレスポンス型に変換
	response := &dto.ProjectResponse{
		ID:           project.ID.String(),
		GithubRepoID: project.GithubRepoID,
		Name:         project.Name,
		Description:  project.Description.String,
		IsActive:     project.IsActive.Bool,
		CreatedAt:    project.CreatedAt.Time,
		UpdatedAt:    project.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *ProjectController) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id format")
	}

	req := new(dto.UpdateProjectRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := c.projectService.Update(id, req.Name, req.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// DBモデルからDTOのレスポンス型に変換
	response := &dto.ProjectResponse{
		ID:           project.ID.String(),
		GithubRepoID: project.GithubRepoID,
		Name:         project.Name,
		Description:  project.Description.String,
		IsActive:     project.IsActive.Bool,
		CreatedAt:    project.CreatedAt.Time,
		UpdatedAt:    project.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *ProjectController) Delete(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id format")
	}

	if err := c.projectService.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
