package service

import (
	"encoding/json"
	"net/http"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

type usr struct {
	Name     string
	Status   uint8
	Role     string
	Password string
	Phone    string
	Mail     string
	Intro    string
}

// TODO: add op code to guarantee that User already been checked via system
func RegisterUser(c *gin.Context) {
	ut := &usr{}
	b, err := c.GetRawData()
	ep(err)
	err = json.Unmarshal(b, ut)
	ep(err)
	err = models.RetriveMySqlDbAccessModel().CreateUser(ut.Name, ut.Password, ut.Phone, ut.Mail)
	ep(err)

	c.JSON(http.StatusOK, "create usr")
	c.Abort()
}

func GetUserInfo(c *gin.Context) {
	// should retrieve user from session ...
	// c.Get("name")
	m, err := models.RetriveMySqlDbAccessModel().GetUserInfo("jim")
	ep(err)

	c.JSON(http.StatusOK, m)
	c.Abort()
}

func UpdateUserInfos(c *gin.Context) {
	ut := make(map[string]interface{})
	b, err := c.GetRawData()
	ep(err)
	err = json.Unmarshal(b, &ut)
	ep(err)
	err = models.RetriveMySqlDbAccessModel().UpdateUserInfos(ut)
	ep(err)

	c.JSON(http.StatusOK, "update success")
	c.Abort()
}

func ChangeUserStatus(c *gin.Context) {
	ut := &usr{}
	b, err := c.GetRawData()
	ep(err)
	err = json.Unmarshal(b, ut)
	ep(err)
	err = models.RetriveMySqlDbAccessModel().ChangeUserStatus(ut.Name, ut.Status)
	ep(err)

	c.JSON(http.StatusOK, "change success")
	c.Abort()
}

func ChangeUserRole(c *gin.Context) {
	ut := &usr{}
	b, err := c.GetRawData()
	ep(err)
	err = json.Unmarshal(b, ut)
	ep(err)
	err = models.RetriveMySqlDbAccessModel().UpdateUserRole(ut.Name, ut.Role)
	ep(err)

	c.JSON(http.StatusOK, "change success")
	c.Abort()
}

// UnregisterUser : normal account can't use this
func UnregisterUser(c *gin.Context) {
	ut := &usr{}
	b, err := c.GetRawData()
	ep(err)
	err = json.Unmarshal(b, ut)
	ep(err)
	err = models.RetriveMySqlDbAccessModel().DeleteUser(ut.Name)
	ep(err)

	c.JSON(http.StatusOK, "delete success")
	c.Abort()
}
