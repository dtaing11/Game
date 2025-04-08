package engine

import (
	"log"
	"runtime"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)


type Game interface {
	Init()
	Update(dt float32)
	Draw()
	Dispose()
}

type Config interface{
	Width, Height int
	Title string
	Vsync bool
	FixedUpdateHz int
}


func Run (game Game, cfg Config){
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatal("Failed to initialize GLFW: %v ", err)
	}
	defer glfw.Terminate()

	win, err := glfw.CreateWindow(cfg.Width, cfg.Height, cfg.Title, nil, nil)
	if err != nil {
		log.Fatal("Failed to create window: %v", err)
	}
	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatal("Failed to initialize OpenGL: %v", err)
	}

	if cfg.Vsync {
		glfw.SwapInterval(1)
	}else {
		glfw.SwapInterval(0)
	}
	game.Init()


	var(
		fixedDT := time.Duration(0)
		accum := time.Duration(0)
		prev := time.Now()
		now := time.Now()

	)
	if cfg.FixedUpdateHz > 0 {
		fixedDT = time.Second / time.Duration(cfg.FixedUpdateHz)
	}

	for !win.ShouldClose() {
		now := time.Now()
		delta := now.Sub(prev)
		prev = now

		glfw.PollEvents()

		if fixedDT > 0 {
			accum += delta
			for accum >= fixedDT {
				game.Update(float32(fixedDT.Seconds()))
				accum -= fixedDT
			}
		} else {
			game.Update(float32(delta.Seconds()))
		}

		game.Draw()
		win.SwapBuffers()
	}

	game.Dispose()
}



}
