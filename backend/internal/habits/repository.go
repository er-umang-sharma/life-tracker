package habits

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// EnsureHabitsTable creates the habits table if it does not exist
func (r *Repository) EnsureHabitsTable(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS habits (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			frequency TEXT NOT NULL,
			reminder TEXT,
			last_done DATE,
			custom_days TEXT[]
		)
	`)
	return err
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) ListHabits(ctx context.Context) ([]Habit, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, name, frequency, reminder, last_done FROM habits")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	habits := []Habit{}
	for rows.Next() {
		var h Habit
		if err := rows.Scan(&h.ID, &h.Name, &h.Frequency, &h.Reminder, &h.LastDone); err != nil {
			return nil, err
		}
		habits = append(habits, h)
	}
	return habits, nil
}

func (r *Repository) CreateHabit(ctx context.Context, name, frequency, reminder string) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO habits (name, frequency, reminder) VALUES ($1, $2, $3)", name, frequency, reminder)
	return err
}

// EnsureHabitCompletionsTable creates the habit_completions table if it does not exist
func (r *Repository) EnsureHabitCompletionsTable(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS habit_completions (
			id SERIAL PRIMARY KEY,
			habit_id INTEGER NOT NULL REFERENCES habits(id) ON DELETE CASCADE,
			date DATE NOT NULL,
			UNIQUE(habit_id, date)
		)
	`)
	return err
}

// MarkHabitDoneToday inserts a completion record for today if not already present
func (r *Repository) MarkHabitDoneToday(ctx context.Context, habitID int) error {
	// Only allow marking for today, and only once
	_, err := r.pool.Exec(ctx, `
		INSERT INTO habit_completions (habit_id, date)
		VALUES ($1, CURRENT_DATE)
		ON CONFLICT (habit_id, date) DO NOTHING
	`, habitID)
	return err
}

// ListHabitCompletions returns all completion dates for a habit
func (r *Repository) ListHabitCompletions(ctx context.Context, habitID string) ([]HabitCompletion, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, habit_id, date FROM habit_completions WHERE habit_id = $1 ORDER BY date`, habitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var completions []HabitCompletion
	for rows.Next() {
		var hc HabitCompletion
		if err := rows.Scan(&hc.ID, &hc.HabitID, &hc.Date); err != nil {
			return nil, err
		}
		completions = append(completions, hc)
	}
	return completions, nil
}
