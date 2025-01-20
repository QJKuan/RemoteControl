package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
	"os/exec"
	"strings"
)

func VideoRoute(r *gin.Engine) {
	// 增大音量
	r.GET("/volumeUp", func(c *gin.Context) {
		err := robotgo.KeyTap(robotgo.AudioVolUp)
		if err != nil {
			fmt.Println("增大音量异常 : ", err)
		}
	})

	// 调小声音
	r.GET("/volumeDown", func(c *gin.Context) {
		err := robotgo.KeyTap(robotgo.AudioVolDown)
		if err != nil {
			fmt.Println("减小音量异常 : ", err)
		}
	})

	// 暂停
	r.GET("/pause", func(c *gin.Context) {
		err := robotgo.KeyTap(robotgo.AudioPause)
		if err != nil {
			fmt.Println("暂停异常 : ", err)
		}
	})

	// 静音
	r.GET("/mute", func(c *gin.Context) {
		err := robotgo.KeyTap(robotgo.AudioMute)
		if err != nil {
			fmt.Println("静音异常 : ", err)
		}
	})

	// 快进
	r.GET("/next", func(c *gin.Context) {
		err := robotgo.KeyTap(robotgo.Right)
		if err != nil {
			fmt.Println("下一页异常 : ", err)
		}
	})

	// 后退
	r.GET("/previous", func(c *gin.Context) {
		err := robotgo.KeyTap(robotgo.Left)
		if err != nil {
			fmt.Println("上一页异常 : ", err)
		}
	})

	// 上一集
	r.GET("/lastLevel", func(c *gin.Context) {
		cmd := exec.Command("tasklist")
		// 获取命令的输出
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("上一集异常 : ", err)
			return
		}

		if strings.Contains(string(output), "PotPlayer") {
			err = robotgo.KeyTap(robotgo.Pageup)
			return
		}

		err = robotgo.KeyTap("[")
		if err != nil {
			fmt.Println("上一集异常 : ", err)
		}
	})

	// 下一集
	r.GET("/nextLevel", func(c *gin.Context) {
		cmd := exec.Command("tasklist")
		// 获取命令的输出
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("下一集异常 : ", err)
			return
		}

		if strings.Contains(string(output), "PotPlayer") {
			err = robotgo.KeyTap(robotgo.Pagedown)
			return
		}
		err = robotgo.KeyTap("]")
		if err != nil {
			fmt.Println("下一集异常 : ", err)
		}
	})
}
