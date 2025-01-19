package controller

import (
	"OSS-Matching-ServerSide/internal/controller/dto"
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

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

func (c *RequiredSkillController) Create(ctx echo.Context) error {
	req := new(dto.CreateRequiredSkillRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	skill, err := c.requiredSkillService.Create(req.JobPostingID, req.Name, req.Level)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, skill)
}
