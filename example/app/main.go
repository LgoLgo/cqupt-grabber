package main

import (
	"github.com/LgoLgo/cqupt-grabber/client"
)

func main() {
	a := client.InitApp()
	w := client.InitWindow(a, 800, 600)

	client.InitCTLs(a, w)
	w.ShowAndRun()
}
