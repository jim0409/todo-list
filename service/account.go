package service

import (
	"encoding/json"
	"net/http"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

type usr struct {
	Name       string `json:"name"`
	Status     uint8  `json:"status"`
	Role       string `json:"role"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Mail       string `json:"mail"`
	Intro      string `json:"intro"`
	VerifyCode string `json:"code"`
}

func LoginUser(c *gin.Context) {
	ut := &usr{}
	b, err := c.GetRawData()
	ep(err)
	err = json.Unmarshal(b, ut)
	ep(err)
	m, err := models.RetriveMySqlDbAccessModel().LoginUser(ut.Name, ut.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to login",
		})
		c.Abort()
		return
	}

	jwtcookie := EncryptJwt(m["name"].(string), m["role"].(string))
	// c.SetCookie("token", jwtcookie, 300, "/", "*", false, false) // set-cookie

	c.JSON(http.StatusOK, gin.H{
		"msg":   "login success",
		"token": jwtcookie,
	})

	c.Abort()
}

// TODO: add op code to guarantee that User already been checked via system
func RegisterUser(c *gin.Context) {
	ut := &usr{}
	b, err := c.GetRawData()
	ep(err)
	err = json.Unmarshal(b, ut)
	ep(err)

	// TOOD: 檢查 verified code 然後才註冊使用者，避免直接創建而導致 使用者重複
	err = models.RetriveMySqlDbAccessModel().CreateUser(ut.Name, ut.Password, ut.Phone, ut.Mail)
	ep(err)

	c.JSON(http.StatusOK, "create usr")
	c.Abort()
}

func GetUserInfo(c *gin.Context) {
	// should retrieve user from session ...
	name := c.GetString("name")
	m, err := models.RetriveMySqlDbAccessModel().GetUserInfo(name)
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
