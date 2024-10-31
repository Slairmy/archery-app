package main

import (
	"BasicArchery/table"
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

	myApp := app.New()
	// 明亮模式
	myApp.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
	// Dark模式
	// myApp.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
	myWindow := myApp.NewWindow("Archery")

	// 一个巨大的基础容器,所有的东西都往里面放
	baseContentContainer := container.NewWithoutLayout()

	// 灰色的顶部
	// 上半部分单独带阴影的容器
	grayBackground := canvas.NewRectangle(color.RGBA{
		R: 200,
		G: 200,
		B: 200,
		A: 255,
	})
	grayBackground.Resize(fyne.NewSize(1086, 30))
	baseContentContainer.Add(grayBackground)
	baseContentContainer.Add(widget.NewLabel("SQL查询"))

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
	baseContentContainer.Add(sqlContentContainer)

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
	// 添加进来下拉选择器
	baseContentContainer.Add(selectContainer)

	// 提前定义一个表格
	var resultTable fyne.CanvasObject

	// 查询按钮
	selectButton := widget.NewButton("SQL查询", func() {
		// 渲染之前先删除
		// baseContentContainer.Remove(resultTable)
		// 渲染一个结果表格
		resultTable = table.MakeQueryResultTable()
		resultTable.Resize(fyne.NewSize(786, 400))
		// 设置位置
		resultTable.Move(fyne.NewPos(30, 600))
		// 添加查询结果表格
		baseContentContainer.Add(resultTable)
		// 重新刷新布局
		//baseContentContainer.Refresh()
	})
	selectButton.Resize(fyne.NewSize(90, 40))
	selectButton.Move(fyne.NewPos(936, 230))

	baseContentContainer.Add(selectButton)

	myWindow.SetContent(baseContentContainer)
	myWindow.Resize(fyne.NewSize(1186, 748))
	myWindow.ShowAndRun()
}
