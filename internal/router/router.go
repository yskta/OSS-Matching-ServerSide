package router

import (
	"OSS-Matching-ServerSide/internal/controller"

	"github.com/labstack/echo/v4"
)

type Controllers struct {
	User               *controller.UserController
	Project            *controller.ProjectController
	ProjectContributor *controller.ProjectContributorController
	JobPosting         *controller.JobPostingController
	JobApplication     *controller.JobApplicationController
	ChatMessage        *controller.ChatMessageController
	RequiredSkill      *controller.RequiredSkillController
	UserSkill          *controller.UserSkillController
}

func NewRouter(c *Controllers) *echo.Echo {
	e := echo.New()
	api := e.Group("/api")

	users := api.Group("/users")
	{
		users.POST("", c.User.Create)
		users.POST("/:id/skills", c.UserSkill.Create)
	}

	projects := api.Group("/projects")
	{
		projects.POST("", c.Project.Create)
		projects.POST("/:id/contributors", c.ProjectContributor.Create)
		projects.POST("/:id/job-postings", c.JobPosting.Create)
	}

	// 求人関連
	jobPostings := api.Group("/job-postings")
	{
		jobPostings.POST("/:id/applications", c.JobApplication.Create)
		jobPostings.POST("/:id/required-skills", c.RequiredSkill.Create)
	}

	// 応募関連
	applications := api.Group("/applications")
	{
		applications.POST("/:id/messages", c.ChatMessage.Create)
	}

	return e
}
