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

// @Summary Create a new project contributor
// @Description Create a new project contributor with user ID and role
// @Tags project_contributors
// @Accept json
// @Produce json
// @Param request body dto.CreateProjectContributorRequest true "Project contributor creation request"
// @Success 201 {object} dto.ProjectContributorResponse
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 500 {object} echo.HTTPError "Server error"
// @Router /project_contributors [post]
func (c *ProjectContributorController) Create(ctx echo.Context) error {
	req := new(dto.CreateProjectContributorRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	contributor, err := c.projectContributorService.Create(req.ProjectID, req.UserID, req.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// DBモデルからDTOのレスポンス型に変換
	response := &dto.ProjectContributorResponse{
		ProjectID:           contributor.ProjectID.String(),
		UserID:              contributor.UserID.String(),
		Role:                contributor.Role,
		CanManageJobPosting: contributor.CanManageJobPosting.Bool,
		CreatedAt:           contributor.CreatedAt.Time,
		UpdatedAt:           contributor.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusCreated, response)
}
