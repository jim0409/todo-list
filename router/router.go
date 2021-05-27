package router

import (
	"todo-list/service"

	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/questions/29418478/go-gin-framework-cors
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ApiRouter(r *gin.Engine) {
	// https://stackoverflow.com/questions/29418478/go-gin-framework-cors
	// middleware need to be implement before Group
	r.Use(CORSMiddleware())

	authrized := r.Group("/")
	r1 := authrized.Group("/")
	{
		r1.POST("/note/:id", service.CreateNotes)
		r1.GET("/notes", service.ReadNotes) // pagenation .. ?
		r1.PUT("/note/:id", service.UpdateNotes)
		r1.DELETE("/note/:id", service.DeleteNotes)
	}
}
