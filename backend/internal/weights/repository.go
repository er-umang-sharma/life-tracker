package weights

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

func (r *Repository) EnsureTables(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS workout_plans (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS workout_exercises (
			id SERIAL PRIMARY KEY,
			plan_id INTEGER REFERENCES workout_plans(id) ON DELETE CASCADE,
			name TEXT NOT NULL,
			rep_range TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS workout_logs (
			id SERIAL PRIMARY KEY,
			plan_id INTEGER REFERENCES workout_plans(id) ON DELETE CASCADE,
			date DATE NOT NULL
		);
		CREATE TABLE IF NOT EXISTS exercise_logs (
			id SERIAL PRIMARY KEY,
			log_id INTEGER REFERENCES workout_logs(id) ON DELETE CASCADE,
			exercise_id INTEGER REFERENCES workout_exercises(id) ON DELETE CASCADE,
			weight DOUBLE PRECISION,
			reps INTEGER
		);
	`)
	return err
}

// Add more CRUD methods as needed for plans, logs, etc.
