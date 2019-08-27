package copilot

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// Version is the current app version
const Version = "1.0.0"

var (
	redisClient  *redis.Client
	redisOptions redis.Options
)

func init() {
	err := SetDefaultConfig()
	if err != nil {
		exitWithMessage(err.Error())
	}
	// Redis
	opts, _ := redis.ParseURL(Conf.cacheURL)
	redisClient = redis.NewClient(opts)
}

func finished() {
	redisClient.Close()
}

// Start is the main func to run the cmd that launches the cabsight REST api server
func Start() {
	var err error
	defer finished()

	// Health checks
	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting server...")
	router := gin.Default()
	s := NewServer(redisClient)
	s.Start(router)
}

// PrintVersion returns the the current app version
func PrintVersion() {
	fmt.Println(fmt.Sprintf("Copilot v%s", Version))
	os.Exit(1)
}

func exitWithMessage(message string) {
	fmt.Println("Error:", message)
	os.Exit(1)
}
