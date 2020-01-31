package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Movedst interface {
	up() error
	dn() error
	lt() error
	rt() error
	xuprt() error
	xuplt() error
	xdnrt() error
	xdnlt() error
}

func dn(a *sdl.Rect) error {
	a.Y++
	return nil
}
func up(a *sdl.Rect) error {
	a.Y--
	return nil
}
func lt(a *sdl.Rect) error {
	a.X--
	return nil
}
func rt(a *sdl.Rect) error {
	a.X++
	return nil
}
func xuprt(a *sdl.Rect) error {
	a.X++
	a.Y--
	return nil
}
func xuplt(a *sdl.Rect) error {
	a.X--
	a.Y--
	return nil
}
func xdnrt(a *sdl.Rect) error {
	a.X++
	a.Y++
	return nil
}
func xdnlt(a *sdl.Rect) error {
	a.X--
	a.Y++
	return nil
}
func Boot(r Movedst) error {
	return r.up()
}

func main() {
	//初始化sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {

		panic(err)
	}
	defer sdl.Quit()
	moveli := true
	moveri := 0
	//创建窗口,渲染器
	window, render, err := sdl.CreateWindowAndRenderer(800, 600 /* sdl.WINDOW_FULLSCREEN_DESKTOP|*/, sdl.RENDERER_ACCELERATED)
	//var err不为空时提示错误
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer render.Destroy()

	render.Clear()
	//src := sdl.Rect{0, 0, 800, 600}
	dst := sdl.Rect{X: 0, Y: 0, W: 100, H: 75}
	image1, err := sdl.LoadBMP("../../../veandco/go-sdl2/.go-sdl2-examples/assets/test.bmp")
	defer image1.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s\n", err)
	}
	image2, err := img.Load("../../../veandco/go-sdl2/.go-sdl2-examples/assets/test.png")
	defer image2.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
	}

	texture1, err := render.CreateTextureFromSurface(image2)
	dst.X = 100
	dst.Y = 100
	dst.W = image2.W / 10
	dst.H = image2.H / 10

	wsw, wsh := window.GetSize()
	//事件模块
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
				if (t.Keysym.Sym) == sdl.K_ESCAPE {
					sdl.Quit()
				}
				if (t.Keysym.Sym) == sdl.K_1 {
					moveli = false
				}
				if (t.Keysym.Sym) == sdl.K_2 {
					moveli = true
				}
				break
			}

		}
		//清理屏幕
		render.Clear()
		sdl.Delay(5)
		if moveli {
			if moveri == 0 {
				xdnrt(&dst)
			}
			if moveri == 1 {
				xuprt(&dst)
			}
			if moveri == 2 {
				xuplt(&dst)
			}
			if moveri == 3 {
				xdnlt(&dst)
			}
		}
		//dst=@sdl.Rect
		if dst.X+dst.W >= wsw {
			lt(&dst) //dst.X-=1
			if moveri == 0 {
				moveri = 3
			}
			if moveri == 1 {
				moveri = 2
			}
		} else if dst.X <= 0 {
			rt(&dst) //dst.X+=1
			if moveri == 2 {
				moveri = 1
			}
			if moveri == 3 {
				moveri = 0
			}
		}
		if dst.Y+dst.H >= wsh {
			up(&dst)
			if moveri == 0 {
				moveri = 1
			}
			if moveri == 3 {
				moveri = 2
			}
		} else if dst.Y <= 0 {
			dn(&dst)
			if moveri == 2 {
				moveri = 3
			}
			if moveri == 1 {
				moveri = 0
			}
		}

		//更新窗口表面
		render.Copy(texture1, nil, &dst)
		render.Present()
	}
}
