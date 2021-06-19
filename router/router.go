package router

import (
	"net/http"
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
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func ApiRouter(r *gin.Engine) {
	// https://stackoverflow.com/questions/29418478/go-gin-framework-cors
	// middleware need to be implement before Group
	r.Use(CORSMiddleware())
	// r.Use(service.CheckAuth())

	version := r.Group("/v1")

	ac := version.Group("/usr")
	{
		ac.POST("/verify", nil) // 確認帳號可以使用，並且回傳一個驗證碼，驗證碼來自 authenicator
		ac.POST("/regist", nil) // 透過 admin 註冊一個 user 帳號，回傳 註冊使用者成功與否
		ac.POST("/login", nil)  // 該使用者透過登入後得到的 Set-Cookie 才能看見 html page
	}

	no := version.Group("/note")
	{
		no.Use(service.CheckAuth())
		no.POST("/add", service.CreateNotes)

		no.GET("", service.ReadNoteByPage)       // 顯示第 n 頁 .. 應該用 querystring 顯示
		no.GET("/lists", service.ReadAllNotes)   // get all notes .. 花費太多效能跟傳輸，應該做分頁
		no.GET("/totalpages", service.CountPage) // 依據每頁分頁的數目，回傳總共頁數

		no.PUT("/update/:id", service.UpdateNotes)
		no.DELETE("/delete/:id", service.DeleteNotes)
	}
}
