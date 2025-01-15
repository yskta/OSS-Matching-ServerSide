package controller

import (
	"OSS-Matching-ServerSide/internal/service"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type JobPostingController struct {
	jobPostingService service.JobPostingService
}

func NewJobPostingController(jps service.JobPostingService) *JobPostingController {
	return &JobPostingController{
		jobPostingService: jps,
	}
}

type CreateJobPostingRequest struct {
	ProjectID   uuid.UUID `json:"project_id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" validate:"required"`
	Deadline    time.Time `json:"deadline"`
}

func (c *JobPostingController) Create(ctx echo.Context) error {
	req := new(CreateJobPostingRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	posting, err := c.jobPostingService.Create(req.ProjectID, req.Title, req.Description, req.Status, req.Deadline)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, posting)
}
