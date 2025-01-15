package controller

import (
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProjectContributorController struct {
	projectContributorService service.ProjectContributorService
}

func NewProjectContributorController(pcs service.ProjectContributorService) *ProjectContributorController {
	return &ProjectContributorController{
		projectContributorService: pcs,
	}
}

type CreateProjectContributorRequest struct {
	ProjectID uuid.UUID `json:"project_id" validate:"required"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	Role      string    `json:"role" validate:"required"`
}

func (c *ProjectContributorController) Create(ctx echo.Context) error {
	req := new(CreateProjectContributorRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	contributor, err := c.projectContributorService.Create(req.ProjectID, req.UserID, req.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	
	return ctx.JSON(http.StatusCreated, contributor)
}
