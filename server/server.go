package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"transgo/server/controllers"

	"github.com/gin-gonic/gin"
)

//go:embed app/dist/*
var FS embed.FS

func Run(port string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// 静态文件
	staticFiles, _ := fs.Sub(FS, "app/dist")
	router.StaticFS("/static", http.FS(staticFiles))

	// Routers
	router.GET("/api/v1/addresses", controllers.AddressesController)
	router.GET("/api/v1/uploads/:path", controllers.UploadsController)
	router.GET("/api/v1/qrcodes", controllers.QrcodesController)
	router.POST("/api/v1/files", controllers.FilesController)
	router.POST("/api/v1/txt", controllers.TextsController)

	// Render
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
	router.Run(":" + port)
}
