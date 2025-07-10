// package declaration
package application

// imports
import (
	"fit/infrastructure"
	"fit/model"
	"time"
)

// type definitions
type ExercisePlan struct {
	Date      string     `json:"date"`
	Exercises []Exercise `json:"exercises"`
	T         []string   `json:"timestamps"`
	Ti        int        `json:"-"`
	T1        int        `json:"-"`
}

type Exercise struct {
	Exercise string `json:"exercise"`
	Sets     int    `json:"sets"`
	Reps     int    `json:"reps"`
	T0       int    `json:"-"`
	T1       int    `json:"-"`
	Ts       int    `json:"t"`
	Weight   int    `json:"w"`
}

// data
var (
	ExercisePlans []ExercisePlan
)

// function definitions
func NewExercisePlan(exercise_plan_record *model.ExercisePlanRecord) ExercisePlan {

	var exercise_plan ExercisePlan = ExercisePlan{
		Date:      time.Now().Format("2006-01-02"),
		Exercises: make([]Exercise, len(exercise_plan_record.Exercises)),
	}

	for i := 0; i < len(exercise_plan_record.Exercises); i++ {
		var exercise Exercise = Exercise{
			Ts: int(time.Now().Unix()),
		}
		exercise_plan.Exercises[i] = exercise
	}

	// success
	return exercise_plan
}

func CreateExercisePlans() {

	// load the exercise plan records from redis
	infrastructure.LoadExercisePlanRecords()

	// construct the exercise plans
	for i := 0; i < len(model.ExercisePlanRecords); i++ {

		// construct the english menu
		ExercisePlans = append(ExercisePlans, NewExercisePlan(&model.ExercisePlanRecords[i]))
	}
}
