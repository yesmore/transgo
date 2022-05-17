package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
	文本上传
*/
func TextsController(c *gin.Context) {
	var json struct {
		Raw string `json:"raw"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// 1 获取 transgo.exe 所在目录
		exe, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		dir := filepath.Dir(exe) // 获取程序所在目录
		if err != nil {
			log.Fatal(err)
		}
		filename := uuid.New().String() // 随机生成文件名
		// 2 在exe文件所在目录创建uploads目录
		uploads := filepath.Join(dir, "uploads") // 拼接uploads路径
		err = os.MkdirAll(uploads, os.ModePerm)  // 注意文件权限
		if err != nil {
			log.Fatal(err)
		}
		fullpath := path.Join("uploads", filename+".txt")
		// 3 将文本保存为一个文件
		err = ioutil.WriteFile(filepath.Join(dir, fullpath), []byte(json.Raw), 0644)
		if err != nil {
			log.Fatal(err)
		}
		// 4 返回该文件的下载路径
		c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath})
	}
}
