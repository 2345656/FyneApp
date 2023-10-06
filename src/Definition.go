package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// LineStyle 线条样式
func (Myapp *Config) LineStyle(hx float32, hy float32, tx float32, ty float32, StrokeWidth float32, color color.Color) *canvas.Line {
	return &canvas.Line{
		Position1:   fyne.Position{X: hx, Y: hy},
		Position2:   fyne.Position{X: tx, Y: ty},
		StrokeWidth: StrokeWidth,
		StrokeColor: color,
	}
}

// SetContent 输入框样式
func (Myapp *Config) SetContent(text string, tx float32, ty float32, tw float32, th float32, Ew float32, Eh float32, MultiLine bool) (*widget.Label, *widget.Entry) {
	point := widget.NewLabel(text)
	point.Move(fyne.NewPos(tx, ty))
	point.Resize(fyne.NewSize(tw, th))

	pointEntry := widget.NewMultiLineEntry()
	//实现文本自动换行
	pointEntry.Wrapping = fyne.TextWrapWord
	//pointEntry.SetPlaceHolder("文本内容")
	pointEntry.Resize(fyne.NewSize(Ew, Eh))
	pointEntry.Move(fyne.NewPos(tx+tw, ty))

	return point, pointEntry
}
