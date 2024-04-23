package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
)

func main () {
    temp := board()
	a := app.New()
	w := a.NewWindow("Hello")

    // creating ui elements
    rect  := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 255, A: 255})
	hello := widget.NewLabel(temp)

    // adding to window
    gridContainer := container.New(layout.NewGridLayout(2), rect, hello)
    buttContainer := container.New(layout.NewGridLayout(9))
    mainContainer := container.New(layout.NewGridLayout(3), gridContainer, buttContainer)
    
	w.SetContent(mainContainer)

    // execution
	w.ShowAndRun()
    tidy()
}

func board() string {
    var retval string = "";
    boxs := make([]int8, 81)
    r, c, s := make(map[int]string),make(map[int]string),make(map[int]string)

    for i := 0; i < len(boxs); i++ {
        row := i / 9
        col := i % 9
        sqr := (row / 3) * 3 + (col / 3)

        r[row] = fmt.Sprintf("%s,%d", r[row], i)
        c[col] = fmt.Sprintf("%s,%d", c[col], i)
        s[sqr] = fmt.Sprintf("%s,%d", s[sqr], i)
    }

    for i := 0; i < 9; i++ {
        retval += fmt.Sprintf("%2s\n", s[i][1:])
    }
    //retval += fmt.Sprintln("%m", r)
    return retval;
}

func tidy() {
    fmt.Println("exited")
}
