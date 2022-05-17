package controllers

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
  获取本机IP地址
*/
func AddressesController(c *gin.Context) {
	addrs, _ := net.InterfaceAddrs() // 获取所有网卡的地址
	var result []string
	// 遍历所有的地址
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		// address.(*net.IPNet) 表示 断言 address 为 net.IPNet 类型
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP.String())
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"addresses": result})
}
