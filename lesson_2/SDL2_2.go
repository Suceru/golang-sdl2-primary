package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	//初始化sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {

		panic(err)
	}
	defer sdl.Quit()
	//创建窗口
	window, err := sdl.CreateWindow("lesson2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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
	//src := sdl.Rect{0, 0, 800, 600}
	dst := sdl.Rect{0, 0, 400, 300}
	//surface背景填充为0，填充大小为默认
	surface.FillRect(nil, 0)
	image1, err := sdl.LoadBMP("../../../veandco/go-sdl2/.go-sdl2-examples/assets/test.bmp")
	if err := image1.BlitScaled(nil /*&src*/, surface, &dst); err != nil {

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
