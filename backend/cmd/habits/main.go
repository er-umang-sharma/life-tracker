package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"life-tracker/backend/internal/habits"

)

func main() {
	godotenv.Load() // Load environment variables from .env file
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
	habitsHandler := habits.NewHandler(habitsRepo)

	r.GET("/habits", habitsHandler.ListHabits)
	r.POST("/habits", habitsHandler.CreateHabit)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
