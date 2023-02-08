package main

import (
	"technical-test-skyshi/app/router"
	"technical-test-skyshi/di"
)

func main() {
	server := di.InitializedEchoServer()
	server.Logger.Fatal(server.Start(router.HostURL))
}
