package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

/*
	生成二维码
*/
func QrcodesController(c *gin.Context) {
	// 获取参数
	if content := c.Query("content"); content != "" {
		png, err := qrcode.Encode(content, qrcode.Medium, 256)
		if err != nil {
			log.Fatal(err)
		}
		// content="http://ip:23333/static/downloads?type=text&url=http://ip:23333/uploads/filename.txt"
		// /api/v1/qrcodes?content=[content]
		c.Data(http.StatusOK, "image/png", png) // 展示二维码图片
	} else {
		c.Status(http.StatusPreconditionRequired)
	}
}
