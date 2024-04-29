package main

import (
	"fmt"
    "os"
    "godoku/puzzle"
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
    // rect  := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 255, A: 255})

    // adding to window
    p := puzzle.New()
    //p.GenBlank()
    p.Print()
    
    // Super crude way to gen puzzle by a file
    dat,err := os.ReadFile("puzzle.txt")
    if err != nil {
        panic(err)
    }
    
    var t [81]int
    //var s = t[:]

    for i := 0; i < 81; i++ {
        t[i] = int(dat[i * 2]) - 48
        //s = append(s, int(dat[i * 2]) - 48)
    }

    fmt.Printf("%v\n", t)

    p.SetPuzzle(&t)

    puz := puzzle.GenPuzzle()
    puzzle.Solve(puz)
    
    w.SetContent(genBoard(p.GetBoard()))
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
