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

// @Summary Create a new chat message
// @Description Create a new chat message with job application ID, sender ID, and content
// @Tags chat_messages
// @Accept json
// @Produce json
// @Param request body dto.CreateChatMessageRequest true "Chat message creation request"
// @Success 201 {object} dto.ChatMessageResponse
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 500 {object} echo.HTTPError "Server error"
// @Router /chat_messages [post]
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
