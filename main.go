package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/thanarat/qa-backend/entity"

	answerDelivery "github.com/thanarat/qa-backend/feature/answer/delivery"
	answerRepo "github.com/thanarat/qa-backend/feature/answer/repository"
	answerUsecase "github.com/thanarat/qa-backend/feature/answer/usecase"

	categoryDelivery "github.com/thanarat/qa-backend/feature/category/delivery"
	categoryRepo "github.com/thanarat/qa-backend/feature/category/repository"
	categoryUsecase "github.com/thanarat/qa-backend/feature/category/usecase"

	questionDelivery "github.com/thanarat/qa-backend/feature/question/delivery"
	questionRepo "github.com/thanarat/qa-backend/feature/question/repository"
	questionUsecase "github.com/thanarat/qa-backend/feature/question/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	// Load .env file
	// if err := godotenv.Load(); err != nil {
	// 	return nil, fmt.Errorf("error loading .env file: %v", err)
	// }

	// Get database connection details from environment variables
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	// Create DSN string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	fmt.Println(dsn)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	fmt.Println("DBConnection: Success")

	// Auto migrate the schema
	err = db.AutoMigrate(&entity.Category{}, &entity.Question{}, &entity.Answer{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
}

func main() {
	// Initialize database connection
	db, err := initDB()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Configure CORS
	corsOrigin := os.Getenv("CORS_ALLOWED_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "http://localhost:3000" // Default fallback
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{corsOrigin},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	fmt.Println("Echo init")
	e.Logger.Info("Echo started")
	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	// Initialize handlers
	questionDelivery.NewHandler(e.Group("/question"),
		questionUsecase.NewQuestionUsecase(
			questionRepo.NewQuestionRepo(db),
			categoryRepo.NewCategoryRepo(db)))

	categoryDelivery.NewHandler(e.Group("/category"),
		categoryUsecase.NewCategoryUsecase(
			categoryRepo.NewCategoryRepo(db)))

	answerDelivery.NewHandler(e.Group("/answer"),
		answerUsecase.NewAnswerUsecase(
			answerRepo.NewAnswerRepo(db)))

	// Create default category if none exists
	var count int64
	db.Model(&entity.Category{}).Count(&count)
	if count == 0 {
		defaultCategory := entity.Category{
			Category: "Programming",
		}
		db.Create(&defaultCategory)
		fmt.Println("Created default category")
	}

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
