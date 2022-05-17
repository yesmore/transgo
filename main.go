package main

import (
	"os"
	"os/exec"
	"os/signal"

	"transgo/server"
	browser "transgo/server/constant"
)

func main() {
	port := "23333"

	// 启动gin服务
	go func() {
		server.Run(port)
	}()

	// 启动浏览器
	browserPath := browser.Chrome
	cmd := exec.Command(browserPath, "--app=http://127.0.0.1:"+port+"/static/index.html")
	cmd.Start()

	// 监听 Ctrl+C 中断
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	// 阻塞等待中断信号
	select {
	case <-chSignal:
		println("Shutdown App...")
		cmd.Process.Kill()
	}
}
