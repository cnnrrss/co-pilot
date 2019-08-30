package copilot

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// Server ...
type Server struct {
	cache *RedisCache
}

// NewServer is a constructor for the API
func NewServer(cache *redis.Client) *Server {
	redisCache := NewRedisCache(cache)
	return &Server{
		cache: redisCache,
	}
}

// Start the API server
func (s *Server) Start(router *gin.Engine) {
	router.GET("/ping", s.ping)

	apiv1 := router.Group("/api")
	apiv1.GET("/contact/:id", s.GetContactByID)
	apiv1.POST("/contact/", s.UpsertContact)
	apiv1.PUT("/contact/", s.UpsertContact)

	if Conf.clean == true {
		log.Println("Flushing the cache..")
		s.cache.CleanCache()
	}

	err := router.Run(fmt.Sprintf(":%s", Conf.port))
	if err != nil {
		fmt.Println("Cannot start server:", err)
		os.Exit(1)
	}
}
