package main

import (
	"log"
	"runtime"

	"github.com/dtaing11/Game/window"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	/*

		locks the current Goroutine to the current OS thread, meaning the Go runtime
		scheduler will not move this Goroutine to a different OS thread during its lifetime.
		This can make the application more stable, particularly when interacting with system-level resources or
		external libraries (such as C code) that require running on a specific OS thread. By preventing the Go runtime
		from migrating the Goroutine,
		you reduce the risk of race conditions or crashes due to unexpected thread switching.
	*/
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalf("Failed to initialize GLFW: %v", err)
	}
	defer glfw.Terminate()

	win := window.CreateWindow(800, 600, "GLFW Event Listener")

	window.SetupCallBack(win)

	for !win.ShouldClose() {

		win.SetAspectRatio(16, 9)
		glfw.PollEvents()
		win.SwapBuffers()

	}
}
