package habits

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) ListHabits(c *gin.Context) {
	habits, err := h.repo.ListHabits(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habits)
}

func (h *Handler) CreateHabit(c *gin.Context) {
	var req struct {
		Name      string `json:"name"`
		Frequency string `json:"frequency"`
		Reminder  string `json:"reminder"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Validation: name and frequency must not be empty, frequency must be valid
	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}
	validFrequencies := map[string]bool{
		string(FrequencyDaily):    true,
		string(FrequencyWeekly):   true,
		string(FrequencyWeekdays): true,
		string(FrequencyWeekends): true,
		string(FrequencyCustom):   true,
	}
	if req.Frequency == "" || !validFrequencies[req.Frequency] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing frequency"})
		return
	}
	if err := h.repo.CreateHabit(context.Background(), req.Name, req.Frequency, req.Reminder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

// MarkHabitDone marks a habit as done for the current day (cannot mark for past/future)
func (h *Handler) MarkHabitDone(c *gin.Context) {
	var req struct {
		HabitID int `json:"habit_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Only allow marking for today
	if err := h.repo.MarkHabitDoneToday(context.Background(), req.HabitID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// ListHabitCompletions returns all completion dates for a habit
func (h *Handler) ListHabitCompletions(c *gin.Context) {
	id := c.Param("id")
	completions, err := h.repo.ListHabitCompletions(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, completions)
}
