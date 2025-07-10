// package declaration
package model

// type definitions
type ExercisePlanRecord struct {
	Name      string           `json:"name"`
	Exercises []ExerciseRecord `json:"exercises"`
}

type ExerciseRecord struct {
	Exercise string `json:"exercise"`
	Sets     int    `json:"sets"`
	Reps     int    `json:"reps"`
}

// data
var (
	ExercisePlanRecords []ExercisePlanRecord
)
