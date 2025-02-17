package window

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func CreateWindow (width, height int , title string ) *glfw.Window {
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
		if err != nil {
			log.Fatalf("Failed to create GLFW window: %v", err)
		}

	window.MakeContextCurrent()
	return window;
}
