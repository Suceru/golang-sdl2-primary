package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
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
type watersize struct {
	X     int32
	Y     int32
	R     int32
	speed int32
}

func drawwater(size *watersize, rend *sdl.Renderer, watert *sdl.Texture) {
	dst := sdl.Rect{X: size.X, Y: size.Y, W: size.R, H: size.R}
	rend.Copy(watert, nil, &dst)
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
	moveli = false
	updatali := true
	moveri := 0
	//创建窗口,渲染器
	window, render, err := sdl.CreateWindowAndRenderer(WinW, WinH /*sdl.WINDOW_FULLSCREEN_DESKTOP|*/, sdl.RENDERER_ACCELERATED)
	//var err不为空时提示错误
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer render.Destroy()
	//初始化SDL_ttf系统，错误时将结果传入err变量
	if err := ttf.Init(); err != nil {

		panic(err)
	}
	//函数尾退出SDL2_ttf
	defer ttf.Quit()
	font, err := ttf.OpenFont("../../../veandco/go-sdl2/.go-sdl2-examples/assets/test.ttf", 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	defer font.Close()
	//创建一个font到var solid
	solid, err := font.RenderUTF8Solid("Press 1 to turn on the displascment,`   Press 2 to turn off the displascment,`   Press 3 to turn on the afterimage,`   Press 4 to turn off the afterimage.`    ", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to render text: %s\n", err)
	}
	defer solid.Free()
	var waternum int = 50
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
	wateri, err := sdl.LoadBMP("../sucai/water.bmp")
	defer image3.Free()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s\n", err)
	}

	wsw, wsh := window.GetSize()

	//Set image RGB and remove
	image2.SetColorKey(true, sdl.MapRGB(image2.Format, 255, 0, 0))
	image3.SetColorKey(true, sdl.MapRGB(image3.Format, 255, 255, 255))
	wateri.SetColorKey(true, sdl.MapRGB(wateri.Format, 255, 0, 0))
	texture1, err := render.CreateTextureFromSurface(image1)
	texture2, err := render.CreateTextureFromSurface(image2)
	texture3, err := render.CreateTextureFromSurface(image3)
	texture4, err := render.CreateTextureFromSurface(solid)
	watert, err := render.CreateTextureFromSurface(wateri)
	//set blend mode
	texture1.SetBlendMode(sdl.BLENDMODE_BLEND)

	//prompt
	src.X = 0
	src.Y = 0
	src.W = (solid.W / 4)
	src.H = solid.H
	dst.X = 0
	dst.Y = 0
	dst.W = solid.W / 8
	dst.H = solid.H / 2
	src.X = (solid.W / 4) * 0
	dst.Y = (solid.H / 2) * 0
	render.Copy(texture4, &src, &dst)
	src.X = (solid.W / 4) * 1
	dst.Y = (solid.H / 2) * 1
	render.Copy(texture4, &src, &dst)
	src.X = (solid.W / 4) * 2
	dst.Y = (solid.H / 2) * 2
	render.Copy(texture4, &src, &dst)
	src.X = (solid.W / 4) * 3
	dst.Y = (solid.H / 2) * 3
	render.Copy(texture4, &src, &dst)
	render.Present()
	sdl.Delay(2000)
	println("Quit1")

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
	println("Quit2")
	wz := make([]watersize, waternum)
	println("Quit3")
	for i := 0; i < waternum; i++ {
		wz[i].X = int32(rand.Intn(int(WinW)))
		wz[i].Y = int32(rand.Intn(int(WinH)))
		wz[i].R = int32(rand.Intn(5))
		wz[i].speed = int32(rand.Intn(int(10)))
	}
	println("Quit4")
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
				} else if (t.Keysym.Sym) == sdl.K_1 {
					moveli = true
				} else if (t.Keysym.Sym) == sdl.K_2 {
					moveli = false
				} else if (t.Keysym.Sym) == sdl.K_3 {
					updatali = false
				} else if (t.Keysym.Sym) == sdl.K_4 {
					updatali = true
				}
				break
			}

		}
		//清理屏幕
		if updatali == true {
			render.Clear()
		}
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
		for i := 0; i < waternum; i++ {
			drawwater(&wz[i], render, watert)
			wz[i].Y = wz[i].Y + wz[i].speed
			if wz[i].Y > WinH/8*7 {
				if wz[i].Y >= WinH-int32(rand.Intn(int(WinH))) {
					wz[i].Y = 0
				}
			} else if wz[i].Y >= WinH {
				wz[i].Y = 0
			}
		}
		//add lightning
		if lighto1 <= 4 || (lighto1 <= 10 && lighto1 >= 7) || lighto1 >= 247 {
			lghtg.X = (image3.H / 2) * timeev
			render.Copy(texture3, nil, &lghtg)

		}

		render.Copy(texture2, &src, &dst)
		render.Present()
		sdl.Delay(30)
	}
}
