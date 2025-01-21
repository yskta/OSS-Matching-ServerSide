package controller

import (
	"OSS-Matching-ServerSide/internal/controller/dto"
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

// @Summary Create a new user
// @Description Create a new user with GitHub account
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "User creation request"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 500 {object} echo.HTTPError "Server error"
// @Router /users [post]
func (c *UserController) Create(ctx echo.Context) error {
	req := new(dto.CreateUserRequest)
	// リクエストボディのJSONをGo構造体にマッピング
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := c.userService.Create(req.GithubID, req.Name, req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// DBモデルからDTOのレスポンス型に変換
	response := &dto.UserResponse{
		ID:        user.ID.String(),
		GithubID:  user.GithubID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusCreated, response)
}

// @Summary Get a user by ID
// @Description Get a user's detailed information by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID" format(uuid)
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} echo.HTTPError "Invalid ID format"
// @Failure 404 {object} echo.HTTPError "User not found"
// @Failure 500 {object} echo.HTTPError "Server error"
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

	// DBモデルからDTOのレスポンス型に変換
	response := &dto.UserResponse{
		ID:        user.ID.String(),
		GithubID:  user.GithubID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusOK, response)
}

// @Summary Update user
// @Description Update user information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID" format(uuid)
// @Param request body dto.UpdateUserRequest true "User update request"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} echo.HTTPError "Invalid request format or ID"
// @Failure 404 {object} echo.HTTPError "User not found"
// @Failure 500 {object} echo.HTTPError "Server error"
// @Router /users/{id} [put]
func (c *UserController) Update(ctx echo.Context) error {
	// 文字列のIDをUUIDに変換
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id format")
	}

	req := new(dto.UpdateUserRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := c.userService.Update(id, req.Name, req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// DBモデルからDTOのレスポンス型に変換
	response := &dto.UserResponse{
		ID:        user.ID.String(),
		GithubID:  user.GithubID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusOK, response)
}

// @Summary Delete user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID" format(uuid)
// @Success 204 "No Content"
// @Failure 400 {object} echo.HTTPError "Invalid ID format"
// @Failure 404 {object} echo.HTTPError "User not found"
// @Failure 500 {object} echo.HTTPError "Server error"
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
