package widgets

import (
	"fyne.io/fyne/v2"
)

type SudokuSquareRenderer struct {
    square      *SudokuSquare
}

func (r *SudokuSquareRenderer) Objects () []fyne.CanvasObject {
    return []fyne.CanvasObject{r.square.Square, r.square.Text}
}

func (r *SudokuSquareRenderer) Layout (s fyne.Size) {
}

func (r *SudokuSquareRenderer) MinSize () fyne.Size {
    return fyne.NewSize(0, 0)
}

func (r *SudokuSquareRenderer) Refresh () {
}

func (r *SudokuSquareRenderer) Destroy () {
    r.square.Square.Hide()
    r.square.Text.Hide()
}
