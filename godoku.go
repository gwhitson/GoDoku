package main

import(
    "fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main () {
    temp := board()
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel(temp)
	w.SetContent(container.NewVBox(
		hello,
	))

	w.ShowAndRun()
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
    /*
func main () {
    boxs := make([]int8, 81)
    r, c, s := make(map[int]string),make(map[int]string),make(map[int]string)
    rows := make(map[int]int)
    cols := make(map[int]int)
    sqrs := make(map[int]int)

    for i := range len(boxs){
        rows[i] = i / 9
        cols[i] = i % 9
        sqrs[i] = ((rows[i] / 3) * 3) + (cols[i] / 3)

        r[1] = "test"
        c[1] = "test"
        s[1] = "test"
    }

    for i := range 81 {
        if i % 9 == 0 {
            fmt.Println()
        }
        fmt.Printf("%d, ", sqrs[i])
    }
}
    */
