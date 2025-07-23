package weights

type WorkoutPlan struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Exercises []WorkoutExercise `json:"exercises"`
}

type WorkoutExercise struct {
	ID       int    `json:"id"`
	PlanID   int    `json:"plan_id"`
	Name     string `json:"name"`
	RepRange string `json:"rep_range"`
}

type WorkoutLog struct {
	ID           int           `json:"id"`
	PlanID       int           `json:"plan_id"`
	Date         string        `json:"date"`
	ExerciseLogs []ExerciseLog `json:"exercise_logs"`
}

type ExerciseLog struct {
	ID         int     `json:"id"`
	LogID      int     `json:"log_id"`
	ExerciseID int     `json:"exercise_id"`
	Weight     float64 `json:"weight"`
	Reps       int     `json:"reps"`
}
