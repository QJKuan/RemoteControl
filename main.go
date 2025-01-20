package main

import (
	"RemoteControl/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	server.VideoRoute(r)

	// 获取 IP
	ip := server.GetLocalIP()

	// 加载静态资源
	r.LoadHTMLGlob("resource/*")

	if ip == "" {
		fmt.Println("无法获取当前 IP")
	}
	r.GET("/", func(c *gin.Context) {
		// 使用 c.HTML 方法返回 HTML 页面
		c.HTML(http.StatusOK, "control.html", gin.H{
			"ServiceIP": ip,
		})
	})

	r.Run(":9090")
}
