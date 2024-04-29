package puzzle

func GenPuzzle() map[int][4]int {

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

/*
func (s Sodoku) GenBlank () Sodoku {
    boxs := make(map[int][4]int)

    for i := 0; i < 81; i++ {
        row := i / 9
        col := i % 9
        sqr := (row / 3) * 3 + (col / 3)
        val := 0

        boxs[i] = [4]int{row,col,sqr,val}
    }
    return Sodoku{boxs}
}
*/
