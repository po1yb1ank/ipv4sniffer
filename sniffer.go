package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {

	var ret string
	var output, dest *walk.TextEdit
	MainWindow{
		Background: SystemColorBrush{Color: walk.SystemColor(1)},
		Title:   "IPSniff",
		Size: Size{500, 100},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Column: 1,
				Children: []Widget{
					TextLabel{
						Text:          "Destination",
						TextColor: walk.RGB(0,255,0),
						MaxSize: Size{
							Height: 10,
							Width:  30,
						},
					},
					TextEdit{AssignTo: &dest},
				},
			},
			PushButton{
				Text: "Sniff",
				OnClicked: func() {
					go func(){
						ret = Sniff()
					}()
					output.AppendText(ret)

				},
				MaxSize: Size{
					Width:  30,
					Height: 20,
				},
			},
			TextEdit{
				ReadOnly: true,
				AssignTo: &output,
			},
		},
	}.Run()

}
