package main

import (
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE, proxyServerTE, proxyUserTE, proxyPasswordTE *walk.TextEdit

	var lbl *walk.Label

	MainWindow{
		Title:   "MOS Upload by cURL helper",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			Label{Text: "Proxy Server"},
			TextEdit{AssignTo: &proxyServerTE},
			HSplitter{
				Children: []Widget{
					Label{Text: "Proxy user"},
					Label{Text: "Proxy password"},
				},
			},
			HSplitter{
				Children: []Widget{
					TextEdit{Text: "iz", AssignTo: &proxyUserTE},
					TextEdit{Text: "pwd", AssignTo: &proxyPasswordTE},
				},
			},
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
					lbl.SetText(strings.ToUpper(inTE.Text()))
				},
			},
			Label{
				AssignTo: &lbl,
				Text:     "TEST",
			},
		},
	}.Run()
}
