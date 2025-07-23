package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"life-tracker/backend/internal/habits"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	} // Load environment variables from .env file
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
	habitsRepo := habits.NewRepository(pool)
	// Ensure the habits and habit_completions tables exist
	if err := habitsRepo.EnsureHabitsTable(ctx); err != nil {
		log.Fatalf("Failed to ensure habits table: %v", err)
	}
	if err := habitsRepo.EnsureHabitCompletionsTable(ctx); err != nil {
		log.Fatalf("Failed to ensure habit_completions table: %v", err)
	}
	habitsHandler := habits.NewHandler(habitsRepo)

	// Endpoints for habits
	r.GET("/habits", habitsHandler.ListHabits)
	r.POST("/habits", habitsHandler.CreateHabit)

	// Endpoints for habit completions
	r.GET("/habits/:id/completions", habitsHandler.ListHabitCompletions)
	r.POST("/habits/:id/completions", habitsHandler.MarkHabitDone)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
