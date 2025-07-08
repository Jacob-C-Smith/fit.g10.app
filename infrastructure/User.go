// package declaration
package infrastructure

// imports
import (
	"encoding/json"

	"fit/model"
)

func LoadUserRecords() (err error) {

	// load the user keys
	keys, _, _ := rdb.Scan(ctx, 0, "user:*", -1).Result()

	// get the users
	values, _ := rdb.MGet(ctx, keys...).Result()

	// construct
	for _, v := range values {

		// initialized data
		var item model.UserRecord

		// reflect the user
		err = json.Unmarshal([]byte(v.(string)), &item)

		// add the user to the list of users
		model.UserRecords = append(model.UserRecords, item)
	}

	// done
	return
}
