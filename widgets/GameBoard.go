package widgets

import (
	/*
		    "os"
			"fyne.io/fyne/v2/app"
			"fyne.io/fyne/v2/widget"
			"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	*/
	"fmt"
	"godoku/puzzle"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type GameBoard struct {
    Data        puzzle.Sodoku
}

func (b GameBoard) CreateGrid () *fyne.Container {
    grid := container.NewGridWithColumns(9)

    for y := 0; y < 9; y++ {
        for x := 0; x < 9; x++ {
            fg := canvas.NewText(fmt.Sprintf("%d",b.Data.Board[(y * 9) + x][3]), color.White)
            grid.Add(fg)
        }
    }
    return grid
}
