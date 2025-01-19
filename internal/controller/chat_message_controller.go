package controller

import (
	"OSS-Matching-ServerSide/internal/controller/dto"
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

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

func (c *ChatMessageController) Create(ctx echo.Context) error {
	req := new(dto.CreateChatMessageRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message, err := c.chatMessageService.Create(req.JobApplicationID, req.SenderID, req.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &dto.ChatMessageResponse{
		ID:               message.ID.String(),
		JobApplicationID: message.JobApplicationID.String(),
		SenderID:         message.SenderID.String(),
		Content:          message.Content,
		CreatedAt:        message.CreatedAt.Time,
	}

	return ctx.JSON(http.StatusCreated, response)
}
