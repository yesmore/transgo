package controllers

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
	文件上传
*/
func FilesController(c *gin.Context) {
	file, err := c.FormFile("raw") // 自动读取用户上传文件
	if err != nil {
		log.Fatal(err)
	}

	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Dir(exe)
	if err != nil {
		log.Fatal(err)
	}

	filename := uuid.New().String()
	uploads := filepath.Join(dir, "uploads")
	err = os.MkdirAll(uploads, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	fullpath := path.Join("uploads", filename+filepath.Ext(file.Filename))
	fileErr := c.SaveUploadedFile(file, filepath.Join(dir, fullpath)) // 保存文件
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath})
}
