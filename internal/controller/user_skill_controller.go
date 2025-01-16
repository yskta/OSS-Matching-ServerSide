package controller

import (
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

	"github.com/google/uuid"
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

type CreateUserSkillRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Name   string    `json:"name" validate:"required"`
	Level  string    `json:"level"`
}

func (c *UserSkillController) Create(ctx echo.Context) error {
	req := new(CreateUserSkillRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	skill, err := c.userSkillService.Create(req.UserID, req.Name, req.Level)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, skill)
}
