package main

import (
	"log"

	"github.com/dtaing11/Game/window"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalf("Failed to initialize GLFW: %v", err)
	}
	defer glfw.Terminate()

	win := window.CreateWindow(800, 600, "GLFW Event Listener")
	window.SetupCallbacks(win)

	for !win.ShouldClose() {
		glfw.PollEvents()
		win.SwapBuffers()
	}
}
