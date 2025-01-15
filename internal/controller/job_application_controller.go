package controller

import (
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type JobApplicationController struct {
	jobApplicationService service.JobApplicationService
}

func NewJobApplicationController(jas service.JobApplicationService) *JobApplicationController {
	return &JobApplicationController{
		jobApplicationService: jas,
	}
}

type CreateJobApplicationRequest struct {
	JobPostingID uuid.UUID `json:"job_posting_id" validate:"required"`
	UserID       uuid.UUID `json:"user_id" validate:"required"`
}

func (c *JobApplicationController) Create(ctx echo.Context) error {
	req := new(CreateJobApplicationRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	application, err := c.jobApplicationService.Create(req.JobPostingID, req.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, application)
}
