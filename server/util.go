package server

import (
	"fmt"
	"net"
)

// GetLocalIP 获取本地可用 IP
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	ips := make([]string, 0)
	for _, addr := range addrs {
		// 检查IP地址是否为IPv4并且不是回环地址
		ip, ok := addr.(*net.IPNet)
		if ok && ip.IP.To4() != nil && !ip.IP.IsLoopback() && ip.IP.IsPrivate() {
			//fmt.Println(ip.IP.String())
			ips = append(ips, ip.IP.String())
		}
	}
	if 0 == len(ips) {
		return ""
	} else if 1 == len(ips) {
		fmt.Println("============== 手机浏览器访问以下地址即可打开遥控器: ==============")
		fmt.Println("")
		fmt.Println(ips[0] + ":9090")
		fmt.Println("")
		fmt.Println("=====================================================================")
		return ips[0]
	} else {
		fmt.Println("================== 识别 IP 存在多个, 如下: ==================")
		for _, ip := range ips {
			fmt.Println(ip)
		}
		fmt.Println("")
		fmt.Println("使用存在问题，可查看使用文档手动修改 IP 再启动")
		fmt.Println("==========================================================")
		return ""
	}
}
