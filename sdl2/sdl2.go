package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
	}

}
