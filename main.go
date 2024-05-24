package main

import (
	/*
	   "godoku/GameBoard"
	*/
	"fmt"
	"godoku/widgets"
	"godoku/puzzle"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main () {
	a := app.New()
	w := a.NewWindow("GoDoku")

    one := widget.NewButton("1", func () {fmt.Println("1")})
    two := widget.NewButton("2", func () {fmt.Println("2")})
    thr := widget.NewButton("3", func () {fmt.Println("3")})
    fou := widget.NewButton("4", func () {fmt.Println("4")})
    fiv := widget.NewButton("5", func () {fmt.Println("5")})
    six := widget.NewButton("6", func () {fmt.Println("6")})
    sev := widget.NewButton("7", func () {fmt.Println("7")})
    eig := widget.NewButton("8", func () {fmt.Println("8")})
    nin := widget.NewButton("9", func () {fmt.Println("9")})

    // adding to window
    p := puzzle.New()
    p.Print()
    
    // Super crude way to gen puzzle by a file
    dat,err := os.ReadFile("puzzle.txt")
    if err != nil {
        panic(err)
    }
    
    var t [81]int

    for i := 0; i < 81; i++ {
        t[i] = int(dat[i * 2]) - 48
    }

    p.SetPuzzle(&t)

    puz := puzzle.GenPuzzle()
    puzzle.Solve(puz)
    
    controls := container.New(layout.NewCenterLayout(), widget.NewLabel("top controls area"))
    innerCont := container.New(layout.NewGridLayout(9), one, two, thr, fou, fiv, six, sev, eig, nin)

    gameBoard := widgets.GameBoard{p}

    content := container.NewBorder(controls,innerCont,gameBoard.CreateGrid(),nil,nil)
    w.Resize(fyne.NewSize(600, 1080))
    w.SetContent(content)
    w.ShowAndRun()
    tidy()
}

func tidy() {
    fmt.Println("exited")
}
