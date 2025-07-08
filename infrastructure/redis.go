// package declaration
package infrastructure

// imports
import (
	"context"

	"github.com/redis/go-redis/v9"
)

// data
var (
	client_options *redis.Options = &redis.Options{
		Addr:     "localhost:6379", // <-- TODO: Environment variable
		Password: "",               // <-- TODO: Password, also an environment variable ;)
		DB:       0,
	}
	ctx = context.Background()
	rdb *redis.Client
)

func init() {

	// construct a redis client
	rdb = redis.NewClient(client_options)
}
