package controller

import (
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ChatMessageController struct {
	chatMessageService service.ChatMessageService
}

func NewChatMessageController(cms service.ChatMessageService) *ChatMessageController {
	return &ChatMessageController{
		chatMessageService: cms,
	}
}

type CreateChatMessageRequest struct {
	JobApplicationID uuid.UUID `json:"job_application_id" validate:"required"`
	SenderID         uuid.UUID `json:"sender_id" validate:"required"`
	Content          string    `json:"content" validate:"required"`
}

func (c *ChatMessageController) Create(ctx echo.Context) error {
	req := new(CreateChatMessageRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message, err := c.chatMessageService.Create(req.JobApplicationID, req.SenderID, req.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, message)
}
