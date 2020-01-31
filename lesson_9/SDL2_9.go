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
	window, render, err := sdl.CreateWindowAndRenderer(WinW, WinH /*sdl.WINDOW_FULLSCREEN_DESKTOP|*/, sdl.RENDERER_ACCELERATED)
	//var err不为空时提示错误
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer render.Destroy()

	render.Clear()
	src := sdl.Rect{X: 0, Y: 0, W: 800, H: 600}
	dst := sdl.Rect{X: 0, Y: 0, W: 100, H: 75}
	lghtg := sdl.Rect{X: 0, Y: 0, W: 100, H: 100}
	image1, err := img.Load("../sucai/timg4.jpg")
	defer image1.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load jpg: %s\n", err)
	}
	image2, err := img.Load("../sucai/Aegi_02_8.png")
	defer image2.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
	}
	image3, err := img.Load("../sucai/lightning.jpg")
	defer image3.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load jpg: %s\n", err)
	}

	wsw, wsh := window.GetSize()

	//Set image RGB and remove
	image2.SetColorKey(true, sdl.MapRGB(image2.Format, 255, 0, 0))
	image3.SetColorKey(true, sdl.MapRGB(image3.Format, 255, 255, 255))
	texture1, err := render.CreateTextureFromSurface(image1)
	texture2, err := render.CreateTextureFromSurface(image2)
	texture3, err := render.CreateTextureFromSurface(image3)
	//set blend mode
	texture1.SetBlendMode(sdl.BLENDMODE_BLEND)
	src.X = (image2.W / 4) * 0
	src.Y = (image2.H / 4) * 0
	src.W = image2.W / 4
	src.H = image2.H / 4
	dst.X = 100
	dst.Y = 100
	dst.W = image2.W / 4
	dst.H = image2.H / 4
	lghtg.X = 100
	lghtg.Y = 100
	lghtg.W = image3.W / 4
	lghtg.H = image3.H / 4

	//事件模块
	running := true
	var timeev int32 = 0
	var lighto1 uint8 = 255
	for running {
		lighto1 = lighto1 - 5
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
		texture1.SetAlphaMod(lighto1)
		render.Copy(texture1, nil, nil)
		//add lightning
		if lighto1 <= 4 || (lighto1 <= 10 && lighto1 >= 7) || lighto1 >= 247 {
			lghtg.X = (image3.H / 2) * timeev
			render.Copy(texture3, nil, &lghtg)

		}

		render.Copy(texture2, &src, &dst)
		render.Present()
		sdl.Delay(50)
	}
}
