package main

import (
	"fmt"
	//"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"

    "fyne.io/fyne/v2/container"
	/*
		"fyne.io/fyne/v2/canvas"
		"fyne.io/fyne/v2/layout"*/
	"fyne.io/fyne/v2/widget"
)

func main () {
	a := app.New()
	w := a.NewWindow("Hello")

    // creating ui elements
//    rect  := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 255, A: 255})
	rows := genBoard()
    fmt.Println(rows)

    // adding to window
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
            o.(*widget.Label).SetText(fmt.Sprintf("%v", rows[(i.Row * 9) + i.Col][3]))
        })
    mainContainer := container.New(layout.NewStackLayout(), board)
    
    w.SetContent(mainContainer)

    // execution
	w.ShowAndRun()
    tidy()
}

func genBoard() map[int][4]int {

    boxs := make(map[int][4]int)

    for i := 0; i < 81; i++ {
        row := i / 9
        col := i % 9
        sqr := (row / 3) * 3 + (col / 3)
        val := 0

        boxs[i] = [4]int{row,col,sqr,val}
    }
    return boxs

    /*

    r, c, s := make(map[int]string),make(map[int]string),make(map[int]string)
        c[col] = fmt.Sprintf("%s,%02d", c[col], i)
        r[row] = fmt.Sprintf("%s,%02d", r[row], i)
        s[sqr] = fmt.Sprintf("%s,%02d", s[sqr], i)
    }

    for x := 0; x < 81; x++ {
        retval[x] = x + 1;
    }
    return retval;
    */
}

func tidy() {
    fmt.Println("exited")
}
