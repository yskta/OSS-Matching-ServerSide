package controller

import (
	"OSS-Matching-ServerSide/internal/controller/dto"
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

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

func (c *ProjectContributorController) Create(ctx echo.Context) error {
	req := new(dto.CreateProjectContributorRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	contributor, err := c.projectContributorService.Create(req.ProjectID, req.UserID, req.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, contributor)
}
