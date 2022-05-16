package main

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//go:embed app/dist/*
var FS embed.FS

func main() {
	go func() {
		gin.SetMode(gin.ReleaseMode)
		router := gin.Default()

		staticFiles, _ := fs.Sub(FS, "app/dist")

		router.POST("/api/v1/texts", TextsController)
		router.StaticFS("/static", http.FS(staticFiles))
		// Guard
		router.NoRoute(func(c *gin.Context) {
			path := c.Request.URL.Path
			if strings.HasPrefix(path, "/static/") {
				reader, err := staticFiles.Open("index.html")
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()

				stat, err := reader.Stat() // Statistics FileInfo: file length
				if err != nil {
					log.Fatal(err)
				}
				c.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
			} else {
				c.Status(http.StatusNotFound)
			}
		})
		router.Run(":8080")
	}()

	browserPath := "D:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(browserPath, "--app=http://127.0.0.1:8080/static")
	cmd.Start()

	// 监听 Ctrl+C 中断
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	// 阻塞等待中断信号
	select {
	case <-chSignal:
		println("shutdown spwan...")
		cmd.Process.Kill()
	}
}

func TextsController(c *gin.Context) {
	var json struct {
		Raw string
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
