package puzzle

import (
    "fmt"
)

type Sodoku struct {
    Board map[int][4]int
}

func New () Sodoku {
    m := make(map [int][4]int)
    for i := 0; i < 81; i++ {
        row := i / 9
        col := i % 9
        sqr := (row / 3) * 3 + (col / 3)
        val := 0

        m[i] = [4]int{row,col,sqr,val}
    }
    return Sodoku{m}
}

func (s Sodoku) GetBoard () map[int][4]int {
    return s.Board
}

func (s Sodoku) SetPuzzle (a *[81]int) Sodoku {
    for i := 0; i < 81; i++ {
        row := i / 9
        col := i % 9
        sqr := (row / 3) * 3 + (col / 3)
        s.Board[i] = [4]int{row, col, sqr, a[i]}
        //fmt.Printf("%v\n", [4]int{row, col, sqr, a[i]})
    }

    return s
}

func (s Sodoku) Print () {
    for i := 0; i < 81; i++ {
        fmt.Printf("%v", s.Board[i][3])
        if i % 9 == 8 {
            fmt.Println()
        }
    }
}
