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

var WinW int32 = 800
var WinH int32 = 600

func main() {
	//初始化sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {

		panic(err)
	}
	defer sdl.Quit()
	moveli := true
	moveli = moveli
	moveri := 0
	//创建窗口,渲染器
	window, render, err := sdl.CreateWindowAndRenderer(WinW, WinH, sdl.WINDOW_FULLSCREEN_DESKTOP|sdl.RENDERER_ACCELERATED)
	//var err不为空时提示错误
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer render.Destroy()

	render.Clear()
	src := sdl.Rect{X: 0, Y: 0, W: 800, H: 600}
	dst := sdl.Rect{X: 0, Y: 0, W: 100, H: 75}
	//bg := sdl.Rect{X: 0, Y: 0, W: 100, H: 100}
	image1, err := img.Load("../sucai/timg4.jpg")
	defer image1.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s\n", err)
	}
	image2, err := img.Load("../sucai/Aegi_02.png")
	defer image2.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
	}
	wsw, wsh := window.GetSize()
	texture1, err := render.CreateTextureFromSurface(image1)
	texture2, err := render.CreateTextureFromSurface(image2)
	//texture3, err := render.CreateTextureFromSurface(image3)

	src.X = (image2.W / 4) * 0
	src.Y = (image2.H / 4) * 0
	src.W = image2.W / 4
	src.H = image2.H / 4
	dst.X = 100
	dst.Y = 100
	dst.W = image2.W / 4
	dst.H = image2.H / 4

	//事件模块
	running := true
	var timeev int32 = 0
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
				if (t.Keysym.Sym) == sdl.K_w && t.State == 1 {
					dst.Y -= 10
					src.Y = (image2.H / 4) * 3
					timeev = (timeev + 1) % 4
					src.X = (image2.W / 4) * timeev
				} else if (t.Keysym.Sym) == sdl.K_s && t.State == 1 {
					dst.Y += 10
					src.Y = (image2.H / 4) * 0
					timeev = (timeev + 1) % 4
					src.X = (image2.W / 4) * timeev
				} else if (t.Keysym.Sym) == sdl.K_a && t.State == 1 {
					dst.X -= 10
					src.Y = (image2.H / 4) * 1
					timeev = (timeev + 1) % 4
					src.X = (image2.W / 4) * timeev
				} else if (t.Keysym.Sym) == sdl.K_d && t.State == 1 {
					dst.X += 10
					src.Y = (image2.H / 4) * 2
					timeev = (timeev + 1) % 4
					src.X = (image2.W / 4) * timeev
				} else if t.State == 0 && t.Repeat == 0 {
					src.X = (image2.W / 4) * (src.Y / (image2.H / 4))
				} else if (t.Keysym.Sym) == sdl.K_ESCAPE {
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
		/*if moveli {
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
		*/
		//dst=@sdl.Rect,判断是否到边缘
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

		render.Copy(texture1, nil, nil)
		render.Copy(texture2, &src, &dst)
		render.Present()
		sdl.Delay(50)
	}
}
