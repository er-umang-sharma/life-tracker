package habits

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) ListHabits(ctx context.Context) ([]Habit, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, name FROM habits")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	habits := []Habit{}
	for rows.Next() {
		var h Habit
		if err := rows.Scan(&h.ID, &h.Name); err != nil {
			return nil, err
		}
		habits = append(habits, h)
	}
	return habits, nil
}

func (r *Repository) CreateHabit(ctx context.Context, name string) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO habits (name) VALUES ($1)", name)
	return err
}
