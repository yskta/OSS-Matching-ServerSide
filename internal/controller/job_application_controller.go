package controller

import (
	"OSS-Matching-ServerSide/internal/controller/dto"
	"OSS-Matching-ServerSide/internal/service"
	"net/http"

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

func (c *JobApplicationController) Create(ctx echo.Context) error {
	req := new(dto.CreateJobApplicationRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	application, err := c.jobApplicationService.Create(req.JobPostingID, req.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, application)
}
