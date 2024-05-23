package main

import (
	"fmt"
    "os"
    "godoku/puzzle"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type GameBoard struct {
    Data        puzzle.Sodoku
    Controls    *fyne.Container
    Numbuts     *fyne.Container
    SqrEntries  map[int]fyne.CanvasObject
    Gamesqrs    *fyne.Container
}


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

    fmt.Printf("%v\n", t)

    p.SetPuzzle(&t)

    puz := puzzle.GenPuzzle()
    puzzle.Solve(puz)
    
    controls := container.New(layout.NewCenterLayout(), widget.NewLabel("top controls area"))
    board, entries := genBoard(p)
    innerCont := container.New(layout.NewGridLayout(9), one, two, thr, fou, fiv, six, sev, eig, nin)

    gameBoard := GameBoard{p, controls, innerCont, entries, board}

    content := container.NewBorder(gameBoard.Controls, gameBoard.Numbuts, gameBoard.Gamesqrs, nil, nil)
    w.Resize(fyne.NewSize(600, 1080))
    w.SetContent(content)
    w.ShowAndRun()
    tidy()
}

func genBoard (puz puzzle.Sodoku) (*fyne.Container, map[int]fyne.CanvasObject) {
    entries := make(map[int]fyne.CanvasObject)

    for x := range 81 {
        entries[x] = widget.NewRichTextWithText(fmt.Sprint(puz.Board[x][3]))
        fmt.Println(puz.Board[x])
    }

    one := container.New(layout.NewGridLayout(9), entries[0],entries[1],entries[2],entries[3],entries[4],entries[5],entries[6],entries[7],entries[8]) 
    two := container.New(layout.NewGridLayout(9), entries[9],entries[10],entries[11],entries[12],entries[13],entries[14],entries[15],entries[16],entries[17])
    thr := container.New(layout.NewGridLayout(9), entries[18],entries[19],entries[20],entries[21],entries[22],entries[23],entries[24],entries[25],entries[26])
    fou := container.New(layout.NewGridLayout(9), entries[27],entries[28],entries[29],entries[30],entries[31],entries[32],entries[33],entries[34],entries[35])
    fiv := container.New(layout.NewGridLayout(9), entries[36],entries[37],entries[38],entries[39],entries[40],entries[41],entries[42],entries[43],entries[44])
    six := container.New(layout.NewGridLayout(9), entries[45],entries[46],entries[47],entries[48],entries[49],entries[50],entries[51],entries[52],entries[53])
    sev := container.New(layout.NewGridLayout(9), entries[54],entries[55],entries[56],entries[57],entries[58],entries[59],entries[60],entries[61],entries[62])
    eig := container.New(layout.NewGridLayout(9), entries[63],entries[64],entries[65],entries[66],entries[67],entries[68],entries[69],entries[70],entries[71])
    nin := container.New(layout.NewGridLayout(9), entries[72],entries[73],entries[74],entries[75],entries[76],entries[77],entries[78],entries[79],entries[80])

    board := container.New(layout.NewVBoxLayout(), one, two, thr, fou, fiv, six, sev, eig, nin)
    return board, entries
}

func tidy() {
    fmt.Println("exited")
}
