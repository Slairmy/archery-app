package table

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// MakeQueryResultTable 使用 _表示参数不会在函数体内使用
func MakeQueryResultTable() fyne.CanvasObject {

	t := widget.NewTableWithHeaders(
		func() (int, int) { return 5, 10 },
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell 000, 000")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			// 表的渲染
			label := cell.(*widget.Label)
			switch id.Col {
			case 0:
				label.SetText("A longer Cell")
			default:
				label.SetText(fmt.Sprintf("Cell %d,%d", id.Row+1, id.Col+1))
			}
		})
	t.SetColumnWidth(0, 102)
	t.SetRowHeight(2, 50)

	return t
}
