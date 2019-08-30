package copilot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cnnrrss/co-pilot/api"
	"github.com/gin-gonic/gin"
)

// GetContactByID takes the ID from the url param and returns the contact from the cache or API
func (s *Server) GetContactByID(c *gin.Context) {
	var err error
	var contact Contact

	id := c.Param("id")
	if id == "" {
		response(c, http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
		return
	}

	// Ignore error if not found int cache
	cacheResp, _ := s.cache.GetContactByID(id)
	if len(cacheResp) > 0 {
		err = json.Unmarshal(cacheResp, &contact)
		if err != nil {
			log.Printf("Could not unmarshal json from cache %s\n", err)
			return
		}

		response(c, http.StatusOK, fmt.Sprintf("contact from cache"), contact)
		return
	}

	resp, err := api.GetContact(id)
	if err != nil {
		response(c, http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response(c, http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
		return
	}
	json.Unmarshal(data, &contact)

	err = s.cache.SetContactByID(id, contact)
	if err != nil {
		log.Printf("Could not update cache %s\n", err)
	}

	response(c, http.StatusOK, fmt.Sprintf("contact from api"), contact)
	return
}

// UpsertContact sends a request to Create/Update a contact to the API
func (s *Server) UpsertContact(c *gin.Context) {
	var contact Contact

	reqData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response(c, http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
		return
	}

	resp, err := api.CreateContact(reqData)
	if err != nil {
		response(c, http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
		return
	}

	respData, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respData))
	if err != nil {
		response(c, http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
		return
	}
	json.Unmarshal(respData, &contact)

	// invalidate cache if contact existed
	s.cache.DeleteContactByID(contact.ID)
	response(c, http.StatusOK, fmt.Sprintf("contact %s successfully created", contact.ID), nil)
	return
}

func response(c *gin.Context, httpCode int, message string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code":    httpCode,
		"message": message,
		"data":    data,
	})
}

func (s *Server) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
