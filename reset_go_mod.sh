#!/bin/bash
rm go.mod go.sum

go mod init godoku

go get fyne.io/fyne/v2/app
go get fyne.io/fyne/v2/canvas
go get fyne.io/fyne/v2/container
go get fyne.io/fyne/v2/layout
go get fyne.io/fyne/v2/widget

go mod tidy
