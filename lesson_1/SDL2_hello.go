package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//打包成main

//调用sdl官网golang部分下载的包，安装位置自定，@https://github.com/veandco/go-sdl2
func main() {
	//初始化SDL系统，错误时将结果传入err变量
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {

		panic(err)
	}
	//函数尾退出SDL2
	defer sdl.Quit()

	//初始化SDL_ttf系统，错误时将结果传入err变量
	if err := ttf.Init(); err != nil {

		panic(err)
	}
	//函数尾退出SDL2_ttf
	defer ttf.Quit()

	/*
	*@func CreateWindow(title string, x, y, w, h int32, flags uint32) (*Window, error)
	*创建一个窗口标题为"hello"的800*600的窗口给var window,状态为展示
	 */
	window, err := sdl.CreateWindow("hello", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	//var err不为空时提示错误
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	//绑定窗口的surface为var surface
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	//绑定字体
	font, err := ttf.OpenFont("../../../veandco/go-sdl2/.go-sdl2-examples/assets/test.ttf", 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	defer font.Close()

	//surface背景填充为0，填充大小为默认
	surface.FillRect(nil, 0)

	//设置一个矩形
	rect := sdl.Rect{0, 0, 200, 200}
	//在矩形内填充为0xffff0000颜色
	surface.FillRect(&rect, 0xffff0000)

	//创建一个font到var solid
	solid, err := font.RenderUTF8Solid("Hello, World!", sdl.Color{0, 0, 255, 255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to render text: %s\n", err)
	}
	defer solid.Free()

	//solid贴在surface上
	if err := solid.Blit(nil, surface, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to put text on window surface: %s\n", err)
	}
	//刷新窗口surface
	window.UpdateSurface()

	//停留在SDL2的界面中
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
