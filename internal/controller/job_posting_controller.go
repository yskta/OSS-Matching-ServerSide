package controller

import (
	"OSS-Matching-ServerSide/internal/controller/dto"
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

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

func (c *JobPostingController) Create(ctx echo.Context) error {
	req := new(dto.CreateJobPostingRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	posting, err := c.jobPostingService.Create(req.ProjectID, req.Title, req.Description, req.Status, req.Deadline)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, posting)
}
