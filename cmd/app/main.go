package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"OSS-Matching-ServerSide/internal/config"
	"OSS-Matching-ServerSide/internal/controller"
	"OSS-Matching-ServerSide/internal/repository"
	"OSS-Matching-ServerSide/internal/router"
	"OSS-Matching-ServerSide/internal/service"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// configの読み込み
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Database config loaded: host=%s, port=%s, dbname=%s", cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName)

	// db接続
	db, err := connectToDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// DB接続確認
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Ping確認
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")

	// リポジトリの初期化
	userRepo := repository.NewUserRepository()
	projectRepo := repository.NewProjectRepository()
	projectContributorRepo := repository.NewProjectContributorRepository()
	jobPostingRepo := repository.NewJobPostingRepository()
	jobApplicationRepo := repository.NewJobApplicationRepository()
	chatMessageRepo := repository.NewChatMessageRepository()
	requiredSkillRepo := repository.NewRequiredSkillRepository()
	userSkillRepo := repository.NewUserSkillRepository()

	// サービスの初期化
	userService := service.NewUserService(db, userRepo)
	projectService := service.NewProjectService(db, projectRepo)
	projectContributorService := service.NewProjectContributorService(db, projectContributorRepo)
	jobPostingService := service.NewJobPostingService(db, jobPostingRepo)
	jobApplicationService := service.NewJobApplicationService(db, jobApplicationRepo)
	chatMessageService := service.NewChatMessageService(db, chatMessageRepo)
	requiredSkillService := service.NewRequiredSkillService(db, requiredSkillRepo)
	userSkillService := service.NewUserSkillService(db, userSkillRepo)

	// コントローラーの初期化
	controllers := &router.Controllers{
		User:               controller.NewUserController(userService),
		Project:            controller.NewProjectController(projectService),
		ProjectContributor: controller.NewProjectContributorController(projectContributorService),
		JobPosting:         controller.NewJobPostingController(jobPostingService),
		JobApplication:     controller.NewJobApplicationController(jobApplicationService),
		ChatMessage:        controller.NewChatMessageController(chatMessageService),
		RequiredSkill:      controller.NewRequiredSkillController(requiredSkillService),
		UserSkill:          controller.NewUserSkillController(userSkillService),
	}
	// Echoインスタンスを作成
	e := router.NewRouter(controllers)

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートハンドラー
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	//サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}

func connectToDB(cfg *config.Config) (*gorm.DB, error) {
	params := []string{
		"host=" + cfg.DB.Host,
		"port=" + cfg.DB.Port,
		"user=" + cfg.DB.User,
		"password=" + cfg.DB.Password,
		"dbname=" + cfg.DB.DBName,
		"sslmode=" + cfg.DB.SSLMode,
	}
	dsn := strings.Join(params, " ")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
