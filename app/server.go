package copilot

import (
	"fmt"
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
	apiv1.GET("/contact", s.GetContactByID)
	apiv1.POST("/contact/:id", s.CreateContact)
	apiv1.PUT("/contact/:id", s.UpdateContactByID)

	err := router.Run(fmt.Sprintf(":%s", Conf.port))
	if err != nil {
		fmt.Println("Cannot start server:", err)
		os.Exit(1)
	}
}
