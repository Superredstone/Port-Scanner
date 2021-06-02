package PortScanner

import (
	"fmt"
	"image/color"
	"strconv"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var (
	border = widget.Border{Color: color.NRGBA{A: 0xff}, CornerRadius: unit.Dp(5), Width: unit.Px(1.5)}

	IP = "localhost"

	PortParagraphTxt string
	progress         float32

	//Widgets
	DeviceIP = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	ScanButton = new(widget.Clickable)
	list       = &layout.List{
		Axis: layout.Vertical,
	}
)

type (
	D = layout.Dimensions
	C = layout.Context
)

func DrawUI(gtx layout.Context, th *material.Theme) layout.Dimensions {

	widgets := []layout.Widget{
		material.H6(th, "Port Scanner").Layout,

		//Device ip form
		func(gtx C) D {
			e := material.Editor(th, DeviceIP, "Device IP: ")

			IP = e.Editor.Text()

			return border.Layout(gtx, func(gtx C) D {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, e.Layout)
			})
		},
		//Scan button

		func(gtx C) D {
			in := layout.UniformInset(unit.Dp(8))
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						for ScanButton.Clicked() {
							InitializePortParagraph()

							//Run scan on ports
							go func() {
								err := Scan(IP)
								if err != nil {
									fmt.Println(err)
								}
							}()
						}

						dims := material.Button(th, ScanButton, "Scan").Layout(gtx)
						pointer.CursorNameOp{Name: pointer.CursorPointer}.Add(gtx.Ops)
						return dims
					})
				}),
			)
		},

		material.ProgressBar(th, progress/65535).Layout,
		material.Label(th, unit.Dp(16), "Scanning port "+strconv.Itoa(int(progress))+" of 65535").Layout,

		func(gtx C) D {
			in := layout.UniformInset(unit.Dp(8))
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						PortParagraph := material.Label(th, unit.Dp(16), PortParagraphTxt).Layout(gtx)

						return PortParagraph
					})
				}),
			)
		},
	}

	//Return widgets
	return list.Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})
}

func InitializePortParagraph() {
	PortParagraphTxt = "Scanning " + IP
}
