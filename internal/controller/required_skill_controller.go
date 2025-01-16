package controller

import (
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RequiredSkillController struct {
	requiredSkillService service.RequiredSkillService
}

func NewRequiredSkillController(rss service.RequiredSkillService) *RequiredSkillController {
	return &RequiredSkillController{
		requiredSkillService: rss,
	}
}

type CreateRequiredSkillRequest struct {
	JobPostingID uuid.UUID `json:"job_posting_id" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Level        string    `json:"level"`
}

func (c *RequiredSkillController) Create(ctx echo.Context) error {
	req := new(CreateRequiredSkillRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	skill, err := c.requiredSkillService.Create(req.JobPostingID, req.Name, req.Level)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, skill)
}
