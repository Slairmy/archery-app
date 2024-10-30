package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

func main() {
	// 初始化一个桌面应用
	// a := app.New()

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

	//w := a.NewWindow("toolBar")
	//
	//toolBar := widget.NewToolbar(widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
	//	log.Println("New Document")
	//}),
	//	widget.NewToolbarSpacer(),
	//	widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	//	widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
	//	widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
	//	widget.NewToolbarSpacer(),
	//	widget.NewToolbarAction(theme.HelpIcon(), func() {
	//		log.Println("Display help")
	//	}),
	//)
	//
	//content := container.NewBorder(toolBar, nil, nil, nil, widget.NewLabel("content"))
	//w.SetContent(content)
	//w.Resize(fyne.NewSize(700, 700))
	//w.ShowAndRun()

	myApp := app.New()
	// 明亮模式
	myApp.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
	myWindow := myApp.NewWindow("Archery")

	// 创建一个 SQL 多行输入框
	sqlContentEntry := widget.NewMultiLineEntry()
	sqlContentEntry.SetPlaceHolder("输入 SQL 查询...")
	sqlContentEntry.Resize(fyne.NewSize(886, 500))

	// sql容器显示边框好定位
	border := canvas.NewRectangle(color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	})
	border.SetMinSize(fyne.NewSize(400, 300))
	border.StrokeColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	border.StrokeWidth = 2
	sqlContentContainer := container.NewWithoutLayout(border, sqlContentEntry)
	sqlContentContainer.Resize(fyne.NewSize(786, 500))
	// 坐标位置
	sqlContentContainer.Move(fyne.NewPos(30, 100))

	// 下拉选择组件暂时2+1个
	// 实例选择器
	instanceSelect := widget.NewSelect([]string{"实例1", "实例2", "实例3"}, func(s string) {
		log.Println(fmt.Sprintf("已经选择实例: %s", s))
	})
	instanceSelect.PlaceHolder = "选择一个实例"

	// 数据库下拉选择器
	databaseSelect := widget.NewSelect([]string{"数据库1", "数据库2", "数据库3"}, func(s string) {
		log.Println(fmt.Sprintf("已经选择数据库: %s", s))
	})
	databaseSelect.PlaceHolder = "选择一个数据库"

	// 查询数量限制器
	limitSelect := widget.NewSelect([]string{"100", "500", "1000"}, func(s string) {
		log.Println(fmt.Sprintf("已经选择查询限制: %s", s))
	})
	limitSelect.SetSelected("100")

	selectContainer := container.NewVBox(instanceSelect, databaseSelect, limitSelect)
	selectContainer.Resize(fyne.NewSize(150, 300))
	// 相对位置
	selectContainer.Move(fyne.NewPos(936, 100))

	// 查询按钮
	selectButton := widget.NewButton("SQL查询", func() {
		sqlContent := sqlContentEntry.Text
		instanceSelectContent := instanceSelect.Selected
		databaseSelectContent := databaseSelect.Selected
		limitSelectContent := limitSelect.Selected

		log.Println("输出各个组件选中的值")
		log.Println(sqlContent, instanceSelectContent, databaseSelectContent, limitSelectContent)
	})
	selectButton.Resize(fyne.NewSize(90, 40))
	selectButton.Move(fyne.NewPos(936, 230))

	// 上半部分单独带阴影的容器
	grayBackground := canvas.NewRectangle(color.RGBA{
		R: 200,
		G: 200,
		B: 200,
		A: 255,
	})
	grayBackground.Resize(fyne.NewSize(1086, 30))

	// 大的括号容器 -- 里面需要包含几个容器
	content := container.NewWithoutLayout(
		grayBackground,
		widget.NewLabel("SQL查询"),
		sqlContentContainer,
		selectContainer,
		selectButton,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(1186, 748))
	myWindow.ShowAndRun()
}
