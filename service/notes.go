package service

import (
	"encoding/json"
	"strconv"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

func ep(err error) {
	if err != nil {
		panic(err)
	}
}

type body struct {
	Title   string
	Content string
}

func CreateNotes(c *gin.Context) {
	ct := &body{}
	b, err := c.GetRawData()
	ep(err)
	err = json.Unmarshal(b, ct)
	ep(err)
	err = models.RetriveMySqlDbAccessModel().CreateNotes(ct.Title, ct.Content)
	ep(err)

	c.JSON(200, "create post")
	c.Abort()
}

func ReadAllNotes(c *gin.Context) {
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
		err = models.RetriveMySqlDbAccessModel().UpdateNotes(id, ct.Title, ct.Content)
		ep(err)
	}

	c.JSON(200, "update post")
	c.Abort()
}

func DeleteNotes(c *gin.Context) {
	id := c.Param("id")
	err := models.RetriveMySqlDbAccessModel().DeleteNote(id)
	ep(err)
	c.JSON(200, "delete post")
	c.Abort()
}

func ReadNoteByPage(c *gin.Context) {
	page := c.DefaultQuery("page", "0")   // default page 0
	limit := c.DefaultQuery("limit", "5") // default limit 5

	pageInt, err := strconv.Atoi(page)
	ep(err)

	limitInt, err := strconv.Atoi(limit)
	ep(err)

	m, err := models.RetriveMySqlDbAccessModel().ReadNoteByPage(pageInt, limitInt)
	ep(err)

	c.JSON(200, m)
	c.Abort()
}

func CountPage(c *gin.Context) {
	limit := c.DefaultQuery("limit", "5") // default limit 5
	pageUint, err := strconv.Atoi(limit)
	ep(err)
	page, err := models.RetriveMySqlDbAccessModel().CountPage(int64(pageUint))
	ep(err)
	c.JSON(200, page)
}
