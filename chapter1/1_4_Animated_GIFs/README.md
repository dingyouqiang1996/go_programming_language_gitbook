# 1.4 Animated GIFs
- lissajous.go
```go
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

```
- 变量声明在包则可在其他包可见，如果声明在函数里则只在函数可见
- 常量只能是数值、字符串或布尔值
- 表达式 `[]color.Color{...}` 和 `gif.GIF{...}` 是复合类型字面量
- 结构体是一组值, 这些值名为 `字段`, 他们组合在一个对象里, 可以视为一个单元
  - 没有被赋值的字段会被初始化为零值
- 结构体的字段通过 `.` 来访问
- lissajous程序外层循环64次, 每次生成一个新的图像, 并将它添加到动画中
  - 所有像素会被设置为白色, 然后用黑色绘制一个有规律的轨迹
  - 传递到内层循环生成一个新的图像，并设置一些像素为黑色
  - 结果通过内置函数append追加
