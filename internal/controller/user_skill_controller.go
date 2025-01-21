package controller

import (
	"OSS-Matching-ServerSide/internal/controller/dto"
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserSkillController struct {
	userSkillService service.UserSkillService
}

func NewUserSkillController(uss service.UserSkillService) *UserSkillController {
	return &UserSkillController{
		userSkillService: uss,
	}
}

// @Summary Create a new user skill
// @Description Create a new user skill with user ID, name, and level
// @Tags user_skills
// @Accept json
// @Produce json
// @Param request body dto.CreateUserSkillRequest true "User skill creation request"
// @Success 201 {object} dto.UserSkillResponse
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 500 {object} echo.HTTPError "Server error"
// @Router /user_skills [post]
func (c *UserSkillController) Create(ctx echo.Context) error {
	req := new(dto.CreateUserSkillRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	skill, err := c.userSkillService.Create(req.UserID, req.Name, req.Level)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &dto.UserSkillResponse{
		ID:        skill.ID.String(),
		UserID:    skill.UserID.String(),
		Name:      skill.Name,
		Level:     skill.Level.String,
		CreatedAt: skill.CreatedAt.Time,
		UpdatedAt: skill.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusCreated, response)
}
