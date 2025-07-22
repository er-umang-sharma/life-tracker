package habits

// HabitCompletion tracks when a habit was marked as done
type HabitCompletion struct {
	ID      int   `json:"id"`
	HabitID int   `json:"habit_id"`
	Date    string `json:"date"` // ISO date string, e.g., 2025-07-22
}

type FrequencyType string

const (
	FrequencyDaily    FrequencyType = "daily"
	FrequencyWeekly   FrequencyType = "weekly"
	FrequencyWeekdays FrequencyType = "weekdays"
	FrequencyWeekends FrequencyType = "weekends"
	FrequencyCustom   FrequencyType = "custom"
)

type Habit struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Frequency FrequencyType `json:"frequency"` // daily, weekly, weekdays, weekends, custom
	Reminder  string        `json:"reminder"`  // e.g., 08:00 AM
	LastDone  string        `json:"last_done"` // ISO date string, e.g., 2025-07-22
	// For custom frequency, CustomDays can contain any combination of weekdays, e.g., ["Monday", "Tuesday", "Thursday"]
	CustomDays []string `json:"custom_days,omitempty"`
}
