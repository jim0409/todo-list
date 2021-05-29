package service

import (
	"encoding/json"
	"fmt"
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

func ReadNoteByPage(c *gin.Context) {
	page, ok := c.GetQuery("page")
	if !ok {
		page = "0"
	}
	fmt.Println(page)

	limit, ok := c.GetQuery("limit")
	if !ok {
		limit = "5"
	}
	fmt.Println(limit)

	pageInt, err := strconv.Atoi(page)
	ep(err)

	limitInt, err := strconv.Atoi(limit)
	ep(err)

	m, err := models.RetriveMySqlDbAccessModel().ReadNoteByPage(pageInt, limitInt)
	ep(err)

	for i, j := range m {
		fmt.Println(i)
		for ii, jj := range j.(map[string]string) {
			fmt.Println(ii)
			fmt.Println(jj)
		}
	}

	c.JSON(200, m)
	c.Abort()
}
