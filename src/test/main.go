package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"time"
)

func main() {
	os.Setenv("FYNE_FONT", "fonts/HanYiFangSongJian-1.ttf")
	myApp := app.New()

	myWindow := myApp.NewWindow("Box Layout")
	myWindow.Resize(fyne.Size{
		Width:  400,
		Height: 500,
	})

	//标签
	label := widget.NewLabel("你好")
	//按钮
	button := widget.NewButton("点击", func() {
		log.Println("tapped")
	})
	//创建文本
	entry := widget.NewEntry()
	//设置新大小(不起作用
	entry.Resize(fyne.Size{
		Width:  100,
		Height: 30,
	})
	//设置提示值
	entry.SetPlaceHolder("hello world")
	//实时更新数据
	entry.OnChanged = func(s string) {
		log.Println(s)
	}
	//可以收集提交信息
	entry.OnSubmitted = func(s string) {
		log.Println(s)
	}
	//追踪光标
	//entry.OnCursorChanged = func() {
	//	log.Println("change")
	//}
	//设置默认值
	entry.Append("0000")
	//启动多行
	entry.MultiLine = true
	//查看是否启动多行(有返回值
	entry.AcceptsTab()
	//刷新
	entry.Refresh()
	//不知道
	//eventOne := fyne.KeyEvent{
	//	Name: fyne.KeyA,
	//	Physical: fyne.HardwareKey{
	//		ScanCode: 64,
	//	},
	//}
	//entry.KeyUp(&eventOne)
	//返回当前位置(左上角坐标
	position := entry.Position()
	log.Println(position)
	//判断是否可见
	visible := entry.Visible()
	log.Println(visible)
	//返回当前选中文字
	text := entry.SelectedText()
	log.Println(text)

	//创建密码框
	psw := widget.NewPasswordEntry()

	//选项
	//单选
	check1 := widget.NewCheck("A", func(b bool) {
		log.Println(b)
	})
	//多选
	group := widget.NewCheckGroup([]string{"A", "B"}, func(strings []string) {
		log.Println(strings)
	})
	//下拉单选
	newSelect := widget.NewSelect([]string{"A", "B", "C"}, func(s string) {
		log.Println(s)
	})

	//表单
	newEntry := widget.NewEntry()
	newEntry.SetPlaceHolder("姓名")

	lineEntry := widget.NewMultiLineEntry()
	lineEntry.SetPlaceHolder("自我介绍")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Text",
				Widget: newEntry,
			},
		},
		//获取提交信息
		OnSubmit: func() {
			log.Println("Form submitted:", newEntry.Text)
			log.Println("multiline:", lineEntry.Text)
		},
	}
	//外部添加表单组件
	form.Append("AreaText", lineEntry)

	//进度条
	//静态
	bar := widget.NewProgressBar()
	//动态(不太会用
	infinite := widget.NewProgressBarInfinite()

	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			//改变值
			time.Sleep(time.Millisecond * 250)
			bar.SetValue(i)
		}
	}()

	//工具栏
	toolbar := widget.NewToolbar(
		//添加按键
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {

		}),
		//分隔符
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {

		}),
		widget.NewToolbarAction(theme.ContentClearIcon(), func() {

		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {

		}),
		//空格
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {

		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
		}),
		//向前
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {

		}),
		//向后
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {

		}),
	)
	//列表
	var data = []string{"a", "b", "c"}
	list := widget.NewList(
		//返回列表大小
		func() int {
			return len(data)
		},
		//部件样式
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		//使用它来设置准备显示的内容。
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(data[id])
		},
	)

	myWindow.SetContent(label)
	myWindow.SetContent(button)
	myWindow.SetContent(entry)
	myWindow.SetContent(psw)
	myWindow.SetContent(form)
	myWindow.SetContent(container.NewVBox(bar, infinite))
	myWindow.SetContent(toolbar)
	myWindow.SetContent(list)
	myWindow.SetContent(container.NewVBox(check1, group, newSelect))
	myWindow.ShowAndRun()
}
