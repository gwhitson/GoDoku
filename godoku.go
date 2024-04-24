package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	/*
	"fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/container"
	"image/color"
	"fyne.io/fyne/v2/canvas"
    */
)

func main () {
	a := app.New()
	w := a.NewWindow("Hello")

    // creating ui elements
//    rect  := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 255, A: 255})

    // adding to window
    puz := genPuzzle()
    
    w.SetContent(genBoard(puz))
    //w.SetContent(append(w.Content(), widget.NewLabel("test")))
    // execution
    w.ShowAndRun()
    tidy()
}

func genBoard (puz map[int][4]int) fyne.CanvasObject {
    board := widget.NewTable(
        // Length callback
        func() (int, int) {
            return 9,9
        },
        // CreateCell callback
        func() fyne.CanvasObject {
            //return widget.NewButton("0", func () {
            //    fmt.Println(" --implement")
            //})
            return widget.NewLabel("0")
        },
        // UpdateCell callback
        func(i widget.TableCellID, o fyne.CanvasObject) {
            //o.(*widget.Button).SetText(fmt.Sprintf("%v", rows[(i.Row * 9) + i.Col][3]))
            //o.(*widget.Button).OnTapped = func (o fyne.CanvasObject) {fmt.Println("anotehr test")}
            o.(*widget.Label).SetText(fmt.Sprintf("%v", puz[(i.Row * 9) + i.Col][3]))
        }) 

    return board
}

func tidy() {
    fmt.Println("exited")
}

func genPuzzle() map[int][4]int {

    boxs := make(map[int][4]int)

    for i := 0; i < 81; i++ {
        row := i / 9
        col := i % 9
        sqr := (row / 3) * 3 + (col / 3)
        val := 0

        boxs[i] = [4]int{row,col,sqr,val}
    }
    return boxs
}
