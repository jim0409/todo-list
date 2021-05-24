package router

import (
	"todo-list/service"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	authrized := r.Group("/")

	authrized.Use()

	r1 := authrized.Group("/")
	{
		r1.POST("/note/:id", service.CreateNotes)
		r1.GET("/notes", service.ReadNotes) // pagenation .. ?
		r1.PUT("/note/:id", service.UpdateNotes)
		r1.DELETE("/note/:id", service.DeleteNotes)
	}
}
