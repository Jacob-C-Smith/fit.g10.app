// package declaration
package infrastructure

// imports
import (
	"encoding/json"
	"fmt"

	"fit/model"
)

func LoadExercisePlanRecords() (err error) {

	// load the exercise plan keys
	keys, _, _ := rdb.Scan(ctx, 0, "exercise_plan:*", -1).Result()

	// get the exercise plans
	values, _ := rdb.MGet(ctx, keys...).Result()

	// construct
	for _, v := range values {

		// initialized data
		var item model.ExercisePlanRecord

		// reflect the exercise plan
		err = json.Unmarshal([]byte(v.(string)), &item)

		// add the exercise plan to the list of exercise plans
		model.ExercisePlanRecords = append(model.ExercisePlanRecords, item)
	}

	fmt.Printf("Loaded %d exercise plan records\n", len(model.ExercisePlanRecords))

	// done
	return
}
