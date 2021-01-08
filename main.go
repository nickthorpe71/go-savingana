package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 400
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, screenWidth, screenHeight)
	if err != nil {
		panic(err)
	}
	defer tex.Destroy()

	pixels := make([]byte, screenWidth*screenHeight*4)

	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			setPixel(x, y, color{byte(x % 255), byte(y % 255), 0}, pixels)
		}
	}

	tex.Update(nil, pixels, screenWidth*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		sdl.Delay(16)
	}

}
