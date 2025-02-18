package window

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func SetupCallBack(win *glfw.Window) {
	win.SetKeyCallback(keyCallback)
	win.SetMouseButtonCallback(mouseButtonCallback)
	win.SetCursorPosCallback(cursorPosCallback)
}

func keyCallback(win *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		fmt.Printf("Key Pressed: %v\n", key)
	} else if action == glfw.Release {
		fmt.Printf("Key Released: %v\n", key)
	}

}

func mouseButtonCallback(win *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		fmt.Printf("Mouse Button Pressed: %v\n", button)

	}

}

func cursorPosCallback(win *glfw.Window, xpos, ypos float64) {
	fmt.Printf("Mouse moved to: (%f, %f)\n", xpos, ypos)
}
