package service

import (
	"net/http"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)

/*
	TODO: refactor CheckPermission .. GetUserRols .. ShowRouter ..
	router implement CheckPermission is enough
*/
func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetString("role")
		obj := c.Request.RequestURI
		act := c.Request.Method

		ok, err := models.RetriveMySqlDbAccessModel().CheckPolicy(sub, obj, act)
		if err != nil || !ok {
			c.JSON(http.StatusForbidden, map[string]string{
				"msg": "forbidden",
			})
			c.Abort()
		}

		c.Next()
	}
}
