package widgets

import (
	/*
		    "os"
			"fyne.io/fyne/v2/app"
			"fyne.io/fyne/v2/widget"
			"fyne.io/fyne/v2/layout"
	"fmt"
	"godoku/puzzle"
	"image/color"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
	*/

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)



type SudokuSquare struct {
    widget.BaseWidget
    Text        *canvas.Text
    Square      *canvas.Rectangle
}

func (s *SudokuSquare) CreateRenderer() SudokuSquareRenderer {
    return SudokuSquareRenderer{s}
}
