package controllers

import (
	"net/http"

	"path/filepath"
	config "transgo/server/config"

	"github.com/gin-gonic/gin"
)

/*
	文件上传
	生成文件下载链接
*/
func UploadsController(c *gin.Context) {
	if path := c.Param("path"); path != "" {
		target := filepath.Join(config.UploadsDir, path) // 拼接上传文件路径（uploads/）
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary") // 转二进制文件
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.Header("Content-Type", "application/octet-stream") // 二进制流 支持任意文件格式
		c.File(target)                                       // 向前端发送文件
	} else {
		c.Status(http.StatusNotFound)
	}
}
