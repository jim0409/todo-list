package service

import (
	"encoding/json"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

func ep(err error) {
	if err != nil {
		panic(err)
	}
}

type body struct {
	Content string
}

func CreateNotes(c *gin.Context) {
	id := c.Param("id")
	ct := &body{}
	b, err := c.GetRawData()
	ep(err)
	err = json.Unmarshal(b, ct)
	ep(err)
	err = models.RetriveMySqlDbAccessModel().CreateNotes(id, ct.Content)
	ep(err)

	c.JSON(200, "create post")
	c.Abort()
}

func ReadNotes(c *gin.Context) {
	m, err := models.RetriveMySqlDbAccessModel().ReadAllNotes()
	ep(err)
	c.JSON(200, m)
	c.Abort()
}

func UpdateNotes(c *gin.Context) {
	id := c.Param("id")
	ct := &body{}
	b, err := c.GetRawData()
	ep(err)
	err = json.Unmarshal(b, ct)
	ep(err)
	if ct != nil {
		err = models.RetriveMySqlDbAccessModel().UpdateNotes(id, ct.Content)
		ep(err)
	}

	c.JSON(200, "update post")
	c.Abort()
}

func DeleteNotes(c *gin.Context) {
	id := c.Param("id")
	err := models.RetriveMySqlDbAccessModel().DeleteNotes(id)
	ep(err)
	c.JSON(200, "delete post")
	c.Abort()
}
