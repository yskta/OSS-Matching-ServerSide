// UserController handles HTTP requests for user operations
package controller

import (
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

// リクエスト構造体
type CreateUserRequest struct {
	GithubID string `json:"github_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// @Summary Create a new user
// @Description Create a new user with the given information
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "User creation request"
// @Success 201 {object} model.User
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /users [post]
func (c *UserController) Create(ctx echo.Context) error {
	req := new(CreateUserRequest)
	// リクエストボディのJSONをGo構造体にマッピング
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := c.userService.Create(req.GithubID, req.Name, req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, user)
}

// @Summary Get a user by ID
// @Description Get a user's information by their UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID" format(uuid)
// @Success 200 {object} model.User
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /users/{id} [get]
func (c *UserController) Get(ctx echo.Context) error {
	// 文字列のIDをUUIDに変換
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id format")
	}

	user, err := c.userService.Get(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

// @Summary Update a user
// @Description Update a user's information by their UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID" format(uuid)
// @Param request body UpdateUserRequest true "User update request"
// @Success 200 {object} model.User
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /users/{id} [put]
func (c *UserController) Update(ctx echo.Context) error {
	// 文字列のIDをUUIDに変換
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id format")
	}

	req := new(UpdateUserRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := c.userService.Update(id, req.Name, req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

// @Summary Delete a user
// @Description Delete a user by their UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID" format(uuid)
// @Success 204 "No Content"
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /users/{id} [delete]
func (c *UserController) Delete(ctx echo.Context) error {
	// 文字列のIDをUUIDに変換
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id format")
	}

	if err := c.userService.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
