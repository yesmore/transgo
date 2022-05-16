#! /bin/bash

echo '#### init dev ####'

export GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go get github.com/silenceper/gowatch
go get -u github.com/gin-gonic/gin

echo '#### over ####'