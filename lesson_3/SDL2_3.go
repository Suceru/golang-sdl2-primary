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
	window, err := sdl.CreateWindow("lesson3", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	//var err不为空时提示错误
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	//创建窗口表面
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
	image1, err := sdl.LoadBMP("C:/Users/fidax/go/src/github.com/veandco/go-sdl2/.go-sdl2-examples/assets/test.bmp")
	defer image1.Free()
	for i := 0; i < 10; i++ {
		if err := image1.BlitScaled(nil, surface, &dst); err != nil {
		}
		dst.X += 75
	}
	dst.X = 0
	dst.Y = 100
	image2, err := img.Load("C:/Users/fidax/go/src/github.com/veandco/go-sdl2/.go-sdl2-examples/assets/test.png")
	defer image2.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
	}
	for i := 0; i < 10; i++ {
		if err := image2.BlitScaled(nil, surface, &dst); err != nil {
		}
		dst.X += 75
	}
	//更新窗口表面
	window.UpdateSurface()

	//暂停
	running := true
	for running {
		//将事件传递给var event，事件不为空时运行for循环，运行完后再次获取事件
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			//如果事件种类为退出事件时，打印Quit，并设置for循环为假
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
