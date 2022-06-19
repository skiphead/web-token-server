package main

import (
	"runtime"
	"web-token-server/internal/server"
)

func main() {
	runtime.GOMAXPROCS(4)
	Server.Run()
}
