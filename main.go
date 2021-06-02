package main

import (
	"fmt"
	"os"

	PortScanner "PortScanner/pkg"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

const (
	height, width = 600, 600
	title         = "Port Scanner"
)

var (
	ops op.Ops
)

func main() {
	go func() {
		//Create window
		w := app.NewWindow(app.Size(unit.Dp(width), unit.Dp(height)), app.Title(title))
		if err := loop(w); err != nil {
			fmt.Println(err)
		}
		os.Exit(1)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			PortScanner.DrawUI(gtx, th)

			e.Frame(gtx.Ops)
		}
	}
}
