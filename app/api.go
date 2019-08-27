package copilot

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MsgFlags maps an error code to a msg
var MsgFlags = map[int]string{
	http.StatusOK:         "ok",
	http.StatusBadRequest: "invalid params",
}

// Response is a response struct for uniform responses from the API server
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// GetContactByID ...
func (s *Server) GetContactByID(c *gin.Context) {
	var err error

	id := c.Query("id")
	if id == "" {
		response(c, http.StatusBadRequest, err)
	}

	// Don't handle error from cache at this point in time
	contact, _ := s.cache.GetContactByID(id)
	if len(contact) == 0 {
		// TODO here is where the call to Autopilot API will go
		// if no results retrieved from cache, waiting for API access

		// TODO if we retrieved data from the cache, we should make
		// a call to set the retrieved contact in cache
	}

	response(c, http.StatusOK, fmt.Sprintf("%v", contact))
}

// CreateContact ...
func (s *Server) CreateContact(c *gin.Context) {
	var contact Contact
	c.BindJSON(&contact)
	s.cache.SetContactByID(contact.contact_id, contact.FirstName)
	response(c, http.StatusOK, fmt.Sprintf("contact %s successfully created", contact.FirstName))
}

// UpdateContactByID ...
func (s *Server) UpdateContactByID(c *gin.Context) {
	response(c, http.StatusOK, fmt.Sprintf("updated contact"))
}

func response(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  getMsg(httpCode),
		Data: data,
	})
	return
}

func getMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[http.StatusBadRequest]
}

func (s *Server) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
