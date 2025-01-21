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

// @Summary Create a new required skill
// @Description Create a new required skill with job posting ID, name, and level
// @Tags required_skills
// @Accept json
// @Produce json
// @Param request body dto.CreateRequiredSkillRequest true "Required skill creation request"
// @Success 201 {object} dto.RequiredSkillResponse
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 500 {object} echo.HTTPError "Server error"
// @Router /required_skills [post]
func (c *RequiredSkillController) Create(ctx echo.Context) error {
	req := new(dto.CreateRequiredSkillRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	skill, err := c.requiredSkillService.Create(req.JobPostingID, req.Name, req.Level)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &dto.RequiredSkillResponse{
		ID:           skill.ID.String(),
		JobPostingID: skill.JobPostingID.String(),
		Name:         skill.Name,
		Level:        skill.Level.String,
		CreatedAt:    skill.CreatedAt.Time,
		UpdatedAt:    skill.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusCreated, response)
}
