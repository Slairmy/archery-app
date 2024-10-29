package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	// 初始化一个桌面应用
	a := app.New()

	// 各种组件测试
	// box

	// 添加一个box
	//text1 := canvas.NewText("slairmy", color.White)
	//text2 := canvas.NewText("Emma", color.White)
	//text3 := canvas.NewText("(right)", color.White)
	//// text5 := canvas.NewText("(----where----)", color.White)
	//
	//content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)
	//
	//text4 := canvas.NewText("centered", color.White)
	//centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	//
	//text6 := canvas.NewText("第三行", color.White)
	//content3 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text6, layout.NewSpacer(), layout.NewSpacer())
	//
	//w.SetContent(container.New(layout.NewVBoxLayout(), content, centered, content3))

	// Grid layout
	// 窗口title
	//w := a.NewWindow("Grid layout")
	//text1 := canvas.NewText("slairmy", color.White)
	//text2 := canvas.NewText("Emma", color.White)
	//text3 := canvas.NewText("(right)", color.White)
	//
	//// grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
	//// grid := container.New(layout.NewGridLayoutWithColumns(2), text1, text2, text3)
	//grid := container.New(layout.NewGridLayoutWithRows(3), text1, text2, text3)
	// w.SetContent(grid)
	// Grid Wrap
	//w := a.NewWindow("Grid Wrap")
	//text1 := canvas.NewText("1", color.White)
	//text2 := canvas.NewText("2", color.White)
	//text3 := canvas.NewText("3", color.White)
	//
	//gridWrap := container.New(layout.NewGridWrapLayout(fyne.NewSize(500, 200)), text1, text2, text3)
	//w.SetContent(gridWrap)
	//// 重新设置尺寸
	//w.Resize(fyne.NewSize(700, 200))

	// Border
	//w := a.NewWindow("Border")
	//top := canvas.NewText("top bar", color.White)
	//left := canvas.NewText("left", color.White)
	//right := canvas.NewText("right", color.White)
	//bottom := canvas.NewText("bottom", color.White)
	//middle := canvas.NewText("content", color.White)
	//
	//content := container.NewBorder(top, bottom, left, right, middle)
	//w.SetContent(content)

	// Form
	//w := a.NewWindow("Form layout")
	//label1 := widget.NewLabel("Label 1")
	//value1 := widget.NewLabel("Value")
	//label2 := widget.NewLabel("Label 2")
	//value2 := widget.NewLabel("Something")
	//
	//content := container.New(layout.NewFormLayout(), label1, value1, label2, value2)
	//w.SetContent(content)

	// 控件演示,比如按钮,复选框啥的

	// 表单组件
	//w := a.NewWindow("表单组件")
	//
	//// 单行文本输入栏
	//entry := widget.NewEntry()
	//// 多行文本输入栏
	//textArea := widget.NewMultiLineEntry()
	//
	//form := &widget.Form{
	//	// 表单有哪几个项目
	//	Items: []*widget.FormItem{
	//		{Text: "Entry", Widget: entry},
	//	},
	//	OnSubmit: func() {
	//
	//		// 点击之后需要获取各个输入框的值
	//
	//		log.Println("Form submitted", entry.Text)
	//		log.Println("multiline", textArea.Text)
	//		w.Close()
	//	},
	//}
	//
	//// 给表单追加项目
	//form.Append("Text", textArea)
	//w.SetContent(form)

	w := a.NewWindow("toolBar")

	toolBar := widget.NewToolbar(widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
		log.Println("New Document")
	}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := container.NewBorder(toolBar, nil, nil, nil, widget.NewLabel("content"))
	w.SetContent(content)
	w.Resize(fyne.NewSize(700, 700))
	w.ShowAndRun()
}
