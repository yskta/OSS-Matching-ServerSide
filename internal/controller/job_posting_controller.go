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

// @Summary Create a new job posting
// @Description Create a new job posting with project ID, title, description, status, and deadline
// @Tags job_postings
// @Accept json
// @Produce json
// @Param request body dto.CreateJobPostingRequest true "Job posting creation request"
// @Success 201 {object} dto.JobPostingResponse
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 500 {object} echo.HTTPError "Server error"
// @Router /job_postings [post]
func (c *JobPostingController) Create(ctx echo.Context) error {
	req := new(dto.CreateJobPostingRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	posting, err := c.jobPostingService.Create(req.ProjectID, req.Title, req.Description, req.Status, req.Deadline)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &dto.JobPostingResponse{
		ID:          posting.ID.String(),
		ProjectID:   posting.ProjectID.String(),
		Title:       posting.Title,
		Description: posting.Description.String,
		Status:      posting.Status.String(),
		Deadline:    posting.Deadline.Time,
		CreatedAt:   posting.CreatedAt.Time,
		UpdatedAt:   posting.UpdatedAt.Time,
	}

	return ctx.JSON(http.StatusCreated, response)
}
