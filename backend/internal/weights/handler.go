package weights

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) CreateWorkoutPlan(c *gin.Context) {
	var req struct {
		Name      string `json:"name"`
		Exercises []struct {
			Name     string `json:"name"`
			RepRange string `json:"rep_range"`
		} `json:"exercises"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// TODO: Call repo to create plan and exercises
	c.Status(http.StatusCreated)
}

func (h *Handler) ListWorkoutPlans(c *gin.Context) {
	// TODO: Call repo to list plans
	c.JSON(http.StatusOK, []string{})
}

func (h *Handler) LogWorkout(c *gin.Context) {
	var req struct {
		PlanID       int    `json:"plan_id"`
		Date         string `json:"date"`
		ExerciseLogs []struct {
			ExerciseID int     `json:"exercise_id"`
			Weight     float64 `json:"weight"`
			Reps       int     `json:"reps"`
		} `json:"exercise_logs"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// TODO: Call repo to log workout
	c.Status(http.StatusCreated)
}

func (h *Handler) ListWorkoutLogs(c *gin.Context) {
	// TODO: Call repo to list logs
	c.JSON(http.StatusOK, []string{})
}
