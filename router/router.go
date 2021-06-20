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

		// auth ?
		ac.POST("/verify", nil) // 確認帳號可以使用，並且回傳一個驗證碼，驗證碼來自 authenicator
		ac.POST("/login", nil)  // 該使用者透過登入後得到的 Set-Cookie 才能看見 html page

		// normal - account
		ac.POST("/register", service.RegisterUser) // 註冊一個 user 帳號
		ac.GET("/info", service.GetUserInfo)       // 回傳 user 資訊
		ac.PUT("/update", service.UpdateUserInfos) // 更新 user 資訊

		// admin - account
		ac.PUT("/role", service.ChangeUserRole)          // 更新 user
		ac.PUT("/status", service.ChangeUserStatus)      // 註消該 user 帳號
		ac.DELETE("/unregister", service.UnregisterUser) // 註消該 user 帳號

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
