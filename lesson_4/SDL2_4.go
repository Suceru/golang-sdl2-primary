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
	window, err := sdl.CreateWindow("lesson4", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	//var err不为空时提示错误
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	//获得窗口表面
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	//释放表面
	defer surface.Free()

	//src := sdl.Rect{0, 0, 800, 600}
	dst := sdl.Rect{0, 0, 100, 75}
	//surface背景填充为0，填充大小为默认
	surface.FillRect(nil, 0xff00ffff)
	image1, err := sdl.LoadBMP("../../../veandco/go-sdl2/.go-sdl2-examples/assets/test.bmp")
	defer image1.Free()
	for i := 0; i < 10; i++ {
		if err := image1.BlitScaled(nil, surface, &dst); err != nil {
		}
		dst.X += 75
	}
	dst.X = 0
	dst.Y = 100
	image2, err := img.Load("../../../veandco/go-sdl2/.go-sdl2-examples/assets/test.png")
	defer image2.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
	}
	for i := 0; i < 10; i++ {
		if err := image2.BlitScaled(nil, surface, &dst); err != nil {
		}
		dst.X += 75
	}
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			//如果事件种类为退出事件时，打印Quit，并设置for循环为假
			switch t := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
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
		surface.FillRect(nil, 0xff00ffff)
		//更新窗口表面
		if err := image2.BlitScaled(nil, surface, &dst); err != nil {
		}
		window.UpdateSurface()
	}
}
