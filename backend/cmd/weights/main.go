package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"life-tracker/backend/internal/weights"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	r := gin.Default()

	repo := weights.NewRepository(pool)
	if err := repo.EnsureTables(ctx); err != nil {
		log.Fatalf("Failed to ensure weights tables: %v", err)
	}
	handler := weights.NewHandler(repo)

	r.POST("/workout-plans", handler.CreateWorkoutPlan)
	r.GET("/workout-plans", handler.ListWorkoutPlans)
	r.POST("/workout-logs", handler.LogWorkout)
	r.GET("/workout-logs", handler.ListWorkoutLogs)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run weights service: %v", err)
	}
}
