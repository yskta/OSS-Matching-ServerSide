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

// @Summary Create a new job application
// @Description Create a new job application with job posting ID, user ID, and message
// @Tags job_applications
// @Accept json
// @Produce json
// @Param request body dto.CreateJobApplicationRequest true "Job application creation request"
// @Success 201 {object} dto.JobApplicationResponse
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 500 {object} echo.HTTPError "Server error"
// @Router /job_applications [post]
func (c *JobApplicationController) Create(ctx echo.Context) error {
	req := new(dto.CreateJobApplicationRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	application, err := c.jobApplicationService.Create(req.JobPostingID, req.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &dto.JobApplicationResponse{
		ID:           application.ID.String(),
		JobPostingID: application.JobPostingID.String(),
		UserID:       application.UserID.String(),
		Status:       application.Status.String(),
		CreatedAt:    application.CreatedAt.Time,
		UpdatedAt:    application.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusCreated, response)
}
