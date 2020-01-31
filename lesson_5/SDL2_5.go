package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	//初始化sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {

		panic(err)
	}
	defer sdl.Quit()

	//创建窗口
	window, err := sdl.CreateWindow("lesson5", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	//var err不为空时提示错误
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	render, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer render.Destroy()
	render.Clear()
	//src := sdl.Rect{0, 0, 800, 600}
	dst := sdl.Rect{0, 0, 100, 75}
	image1, err := sdl.LoadBMP("../../../veandco/go-sdl2/.go-sdl2-examples/assets/test.bmp")
	defer image1.Free()
	texture1, err := render.CreateTextureFromSurface(image1)

	dst.X = 0
	dst.Y = 100
	image2, err := img.Load("../../../veandco/go-sdl2/.go-sdl2-examples/assets/test.png")
	defer image2.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
	}
	dst.X = 100
	dst.Y = 100
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			//如果事件种类为退出事件时，打印Quit，并设置for循环为假
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break

			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.State)
				dst.X = t.X - (dst.W / 2)
				dst.Y = t.Y - (dst.H / 2)
				break
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
				if t.Button == 3 {
					dst.X += 10
				}
				if t.Button == 1 {
					dst.X -= 10
				}
				break

			case *sdl.KeyboardEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				if (t.Keysym.Sym) == sdl.K_w {
					dst.Y -= 10
				}
				if (t.Keysym.Sym) == sdl.K_s {
					dst.Y += 10
				}
				if (t.Keysym.Sym) == sdl.K_a {
					dst.X -= 10
				}
				if (t.Keysym.Sym) == sdl.K_d {
					dst.X += 10
				}
				break
			}

		}
		render.Clear()
		//更新窗口表面
		render.Copy(texture1, nil, &dst)
		render.Present()
	}
}
