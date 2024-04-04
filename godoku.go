package main

import(
    "fmt"
)

func main () {
    boxs := make([]int8, 81)
    r, c, s := make(map[int]string),make(map[int]string),make(map[int]string)

    for i := range len(boxs){
        row := i / 9
        col := i % 9
        sqr := (row / 3) * 3 + (col / 3)

        r[row] = fmt.Sprintf("%s,%d", r[row], i)
        c[col] = fmt.Sprintf("%s,%d", c[col], i)
        s[sqr] = fmt.Sprintf("%s,%d", s[sqr], i)
    }

    for i := range 9 {
        fmt.Printf("%s\n", s[i][1:])
    }
    fmt.Println("%m", r)
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
