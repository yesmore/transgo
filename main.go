package main

import (
	"os"
	"os/exec"
	"os/signal"

	server "transgo/server"
	browser "transgo/server/constant"
)

var port string = "23333"

func main() {

	go server.RunServer(port)
	cmd := runBrowser()
	chSignal := listenToInterrupt()

	// 阻塞等待中断信号
	select {
	case <-chSignal:
		println("Shutdown App...")
		cmd.Process.Kill()
	}
}

func runBrowser() *exec.Cmd {
	browserPath := browser.Chrome
	cmd := exec.Command(browserPath, "--app=http://127.0.0.1:"+port+"/static/index.html")
	cmd.Start()
	return cmd
}

// 监听中断信号
func listenToInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
