package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"os"
	"strconv"
)

// IdLabel 标记题号
var IdLabel *widget.Label
var ChooseGroup *widget.RadioGroup
var filepath = "..\\需审查题目\\入学考试题库.txt"
var DeriveFilePath = "..\\已审查题目\\导出试题.txt"

//10.5 20.40
//问题: 无法显示删除的题目

//绑定数据

// AllEntry 存储可写入口
var AllEntry = make([]*widget.Entry, 0)

// SetWindow 应用窗体设计
func (Myapp *Config) SetWindow() {
	//1.将界面堆好
	Myapp.Window.SetIcon(SetIcon())
	Myapp.Window.Resize(fyne.Size{
		Width:  1000,
		Height: 700,
	})
	Myapp.Window.CenterOnScreen()
	Myapp.Window.SetFixedSize(true)

}

// SetWidget 应用小部件设计
func (Myapp *Config) SetWidget() {
	conSearch := WidgetSearch()
	//画线
	conLine := WidgetLine()
	//提示词
	conTip := WidgetTip()
	//下一题/上一题
	conTool := WidgetChoose()
	//标签
	TeamLabel := widget.NewLabel("---------->校易班技术部出品<----------")
	TeamLabel.Move(fyne.NewPos(360, 620))
	//主体
	conItem := WidgetTabItem()
	//列表(大改)
	var ConList *fyne.Container
	if Myapp.ShowPop == true {
		ConList = PopAModWidgetList(Myapp.PopTopicId)
	} else if Myapp.ShowModify == true {
		ConList = PopAModWidgetList(Myapp.ModifyTopicId)
	} else {
		ConList = WidgetList()
	}
	Myapp.Window.SetContent(container.New(nil, conSearch, conTip, conLine, conItem, TeamLabel, conTool, ConList))
}

// SetMenu 应用菜单设计
func (Myapp *Config) SetMenu() {
	//功能
	Content := fyne.NewMenuItem("显示所有题目", func() {
		if Myapp.TopicLink == nil {
			OpenAllDiaLose()
			return
		}
		Myapp.ShowPop = false
		Myapp.ShowModify = false
		Myapp.SetWidget()
	})
	PopContent := fyne.NewMenuItem("显示删除选项", func() {
		//找题目
		Myapp.LoadPopID()
		//判断是否有内容
		if len(Myapp.PopTopicId) == 0 {
			OpenPopDiaLose()
			return
		}
		//显示内容
		Myapp.ShowPop = true
		Myapp.ShowModify = false
		OpenPopDiaSuccess()
		Myapp.SetWidget()
	})
	ModifyContent := fyne.NewMenuItem("显示修改选项", func() {
		//找题目
		Myapp.LoadModifyID()
		//判断是否有内容
		if len(Myapp.ModifyTopicId) == 0 {
			OpenModifyDiaLose()
			return
		}
		//显示内容
		Myapp.ShowModify = true
		Myapp.ShowPop = false
		OpenModifyDiaSuccess()
		Myapp.SetWidget()
	})
	DeriveFile := fyne.NewMenuItem("导出", func() {
		count := DeriveData(Myapp.TopicLink)
		SetDeriveFileDia(count)
	})

	AboutMe := fyne.NewMenuItem("关于我", func() {
		AboutMeDia()
	})

	Function := fyne.NewMenu("功能", Content, PopContent, ModifyContent, DeriveFile)
	About := fyne.NewMenu("联系", AboutMe)

	MainMenu := fyne.NewMainMenu(Function, About)

	//添加
	Myapp.Window.SetMainMenu(MainMenu)
}

// WidgetSearch 搜索框
func WidgetSearch() *fyne.Container {
	SearchText := ""
	//搜索设计
	SearchEntry := widget.NewEntry()
	SearchEntry.SetPlaceHolder("填写搜索题目序号!")
	SearchEntry.Resize(fyne.NewSize(200, 30))
	SearchEntry.Move(fyne.NewPos(0, 0))
	SearchEntry.OnChanged = func(s string) {
		SearchText = s
	}
	SearchButton := widget.NewButtonWithIcon("", theme.SearchIcon(), func() {
		IdNum, _ := strconv.Atoi(SearchText)
		if IdNum <= 0 || IdNum > len(Myapp.TopicLink) {
			SetSearchDia()
			return
		}
		//赋值
		Myapp.Id = IdNum
		//修正
		Myapp.Id -= 1
		//更改内容
		Myapp.SetEntryText()
		//初始化删除和修改
		Myapp.InitPopAndModify()
	})
	SearchButton.Resize(fyne.NewSize(40, 30))
	SearchButton.Move(fyne.NewPos(200, 0))

	con1 := container.New(nil, SearchEntry, SearchButton)
	con1.Resize(fyne.NewSize(1000, 30))

	return con1
}

// WidgetLine 画线布局
func WidgetLine() *fyne.Container {
	lineTop := Myapp.LineStyle(-3, 35, 1000, 35, 3, color.Black)
	lineLeft := Myapp.LineStyle(-3, 35, -3, 1000, 3, color.Black)
	lineBottom := Myapp.LineStyle(-3, 673, 1000, 673, 3, color.Black)
	lineRight := Myapp.LineStyle(995, 35, 995, 673, 3, color.Black)

	conLine := container.New(nil, lineTop, lineLeft, lineBottom, lineRight)

	return conLine
}

// WidgetTip 添加提示词
func WidgetTip() *fyne.Container {
	IdLabel := widget.NewLabel("题目序号")
	ContentLabel := widget.NewLabel("题目内容")
	IdLabel.Move(fyne.NewPos(-5, 36))
	ContentLabel.Move(fyne.NewPos(480, 36))
	conId := container.New(nil, IdLabel, ContentLabel)

	return conId
}

// WidgetChoose 上下一题
func WidgetChoose() *fyne.Container {
	PopIndex := 0
	ModifyIndex := 0

	LabelHead := widget.NewLabel("上一题")
	toolHead := widget.NewToolbar(
		//向前
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {
			if Myapp.ShowPop == true {
				PopIndex -= 1
				if PopIndex < 0 {
					PopIndex = 0
					return
				}
				Myapp.Id = Myapp.PopTopicId[PopIndex] - 1
			} else if Myapp.ShowModify == true {
				ModifyIndex -= 1
				if ModifyIndex < 0 {
					ModifyIndex = 0
					return
				}
				Myapp.Id = Myapp.ModifyTopicId[ModifyIndex] - 1
			} else {
				//记录和判断Id
				Myapp.Id -= 1
				if Myapp.Id < 0 {
					Myapp.Id = 0
					return
				}
			}
			//更改内容
			Myapp.SetEntryText()
			//初始化删除和修改
			Myapp.InitPopAndModify()
		}),
	)
	HeadContent := container.New(layout.NewHBoxLayout(), LabelHead, toolHead)

	LabelTail := widget.NewLabel("下一题")
	toolTail := widget.NewToolbar(
		//向后
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {
			if Myapp.ShowPop == true {
				PopIndex += 1
				if PopIndex > len(Myapp.PopTopicId)-1 {
					PopIndex = len(Myapp.PopTopicId) - 1
					return
				}
				Myapp.Id = Myapp.PopTopicId[PopIndex] - 1
			} else if Myapp.ShowModify == true {
				ModifyIndex += 1
				if ModifyIndex > len(Myapp.ModifyTopicId)-1 {
					ModifyIndex = len(Myapp.ModifyTopicId) - 1
					return
				}
				Myapp.Id = Myapp.ModifyTopicId[ModifyIndex] - 1
			} else {
				Myapp.Id += 1
				if Myapp.Id >= len(Myapp.TopicLink) {
					Myapp.Id = len(Myapp.TopicLink)
					return
				}
			}
			//更改内容
			Myapp.SetEntryText()
			//初始化删除和修改
			Myapp.InitPopAndModify()
		}),
	)
	TailContent := container.New(layout.NewHBoxLayout(), LabelTail, toolTail)

	conTool := container.New(layout.NewVBoxLayout(), HeadContent, TailContent)
	conTool.Resize(fyne.NewSize(50, 150))
	conTool.Move(fyne.NewPos(800, 300))

	return conTool
}

// WidgetBool 删除和修改操作
func WidgetBool() *fyne.Container {
	FuncLabel := widget.NewLabel("  操作卡  ")
	IdLabel = widget.NewLabel("当前题目序号: ")
	FuncLabel.Resize(fyne.NewSize(60, 30))
	IdLabel.Resize(fyne.NewSize(60, 30))
	//单选组//删除//修改
	choose := []string{"删除该题", "修改"}
	ChooseGroup = widget.NewRadioGroup(choose, func(s string) {
		Myapp.SetPopAndModify(choose, s)
	})
	ChooseGroup.Resize(fyne.NewSize(60, 60))

	conCheck := container.New(layout.NewVBoxLayout(), FuncLabel, ChooseGroup, IdLabel)
	conCheck.Move(fyne.NewPos(730, 100))
	conCheck.Resize(fyne.NewSize(60, 120))

	return conCheck
}

// InitPopAndModify 初始化
func (Myapp *Config) InitPopAndModify() {
	if Myapp.TopicLink[Myapp.Id].IsPop == true {
		ChooseGroup.SetSelected("删除该题")
	} else if Myapp.TopicLink[Myapp.Id].IsModify == true {
		ChooseGroup.SetSelected("修改")
	} else {
		ChooseGroup.SetSelected("")
	}
}

// WidgetTabItem 内容布局
func WidgetTabItem() *fyne.Container {
	conCheck := WidgetBool()
	//主体
	point, pointEntry := Myapp.SetContent("分数: ", 120, 40+40, 50, 30, 100, 30, false)
	content, contentEntry := Myapp.SetContent("题目: ", 120, 85+40, 50, 30, 500, 150, true)
	SelectOne, SelectOneEntry := Myapp.SetContent("选项A: ", 120, 250+40, 50, 30, 500, 30, false)
	SelectTwo, SelectTwoEntry := Myapp.SetContent("选项B: ", 120, 295+40, 50, 30, 500, 30, false)
	SelectThree, SelectThreeEntry := Myapp.SetContent("选项C: ", 120, 340+40, 50, 30, 500, 30, false)
	SelectFour, SelectFourEntry := Myapp.SetContent("选项D: ", 120, 385+40, 50, 30, 500, 30, false)

	tabItem := container.New(nil, point, pointEntry, content, contentEntry, SelectOne,
		SelectOneEntry, SelectTwo, SelectTwoEntry,
		SelectThree, SelectThreeEntry, SelectFour, SelectFourEntry, conCheck)

	//清空AllEntry
	AllEntry = make([]*widget.Entry, 0)
	//添加entry
	AllEntry = append(AllEntry, pointEntry)
	AllEntry = append(AllEntry, contentEntry)
	AllEntry = append(AllEntry, SelectOneEntry)
	AllEntry = append(AllEntry, SelectTwoEntry)
	AllEntry = append(AllEntry, SelectThreeEntry)
	AllEntry = append(AllEntry, SelectFourEntry)

	return tabItem
}

// WidgetList 列表布局
func WidgetList() *fyne.Container {
	data := make([]string, 0)
	for i := 1; i <= len(Myapp.TopicLink); i++ {
		data = append(data, strconv.Itoa(i))
	}
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("List", func() {})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Button).SetText(data[i])
			o.(*widget.Button).OnTapped = func() {
				Myapp.Id = i
				//更改内容
				Myapp.SetEntryText()
				//初始化删除和修改
				Myapp.InitPopAndModify()
			}
		})
	//设置
	list.Resize(fyne.NewSize(100, 590))

	conList := container.New(nil, list)
	conList.Resize(fyne.NewSize(100, 600))
	conList.Move(fyne.NewPos(0, 80))

	return conList
}

// PopAModWidgetList 删除和修改列表布局
func PopAModWidgetList(DataId []int) *fyne.Container {
	data := make([]string, 0)
	for _, i := range DataId {
		data = append(data, strconv.Itoa(i))
	}
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("Pop", func() {})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Button).SetText(data[i])
			o.(*widget.Button).OnTapped = func() {
				Myapp.Id = DataId[i] - 1
				//更改内容
				Myapp.SetEntryText()
				//初始化删除和修改
				Myapp.InitPopAndModify()
			}
		})
	//设置
	list.Resize(fyne.NewSize(100, 590))

	conList := container.New(nil, list)
	conList.Resize(fyne.NewSize(100, 600))
	conList.Move(fyne.NewPos(0, 80))

	return conList
}

// SetEntryText 写入内容
func (Myapp *Config) SetEntryText() {

	IdText := fmt.Sprintf("当前题目序号: %d", Myapp.Id+1)
	IdLabel.SetText(IdText)

	Myapp.EnWrite = false
	//清空
	AllEntry[0].SetText("")
	AllEntry[1].SetText("")
	for j := 0; j < 4; j++ {
		AllEntry[2+j].SetText("")
	}
	//写入内容
	AllEntry[0].SetText(Myapp.TopicLink[Myapp.Id].Point)
	AllEntry[1].SetText(Myapp.TopicLink[Myapp.Id].Topic)
	for i := 0; i < len(Myapp.TopicLink[Myapp.Id].Choice); i++ {
		AllEntry[2+i].SetText(Myapp.TopicLink[Myapp.Id].Choice[i])
	}
	//判断是否修改内容
	Myapp.EnWrite = true
	//写回去
	Myapp.EntryBackLink()
}

// EntryBackLink 将修改内容写回TopicLink中
func (Myapp *Config) EntryBackLink() {
	AllEntry[0].OnChanged = func(ModifyValue string) {
		if Myapp.EnWrite == true {
			Myapp.TopicLink[Myapp.Id].Point = ModifyValue
			//修改按钮
			SetModify()
		}
	}
	AllEntry[1].OnChanged = func(ModifyValue string) {
		if Myapp.EnWrite == true {
			Myapp.TopicLink[Myapp.Id].Topic = ModifyValue
			//修改按钮
			SetModify()
		}
	}
	AllEntry[2].OnChanged = func(ModifyValue string) {
		if Myapp.EnWrite == true {
			Myapp.TopicLink[Myapp.Id].Choice[0] = ModifyValue
			//修改按钮
			SetModify()
		}
	}
	AllEntry[3].OnChanged = func(ModifyValue string) {
		if Myapp.EnWrite == true {
			Myapp.TopicLink[Myapp.Id].Choice[1] = ModifyValue
			//修改按钮
			SetModify()
		}
	}
	AllEntry[4].OnChanged = func(ModifyValue string) {
		if Myapp.EnWrite == true {
			Myapp.TopicLink[Myapp.Id].Choice[2] = ModifyValue
			//修改按钮
			SetModify()
		}
	}
	AllEntry[5].OnChanged = func(ModifyValue string) {
		if Myapp.EnWrite == true {
			Myapp.TopicLink[Myapp.Id].Choice[3] = ModifyValue
			//修改按钮
			SetModify()
		}
	}
}

// SetPopAndModify 修改判断和标签
func (Myapp *Config) SetPopAndModify(choose []string, str string) {
	//设置
	if choose[0] == str {
		Myapp.TopicLink[Myapp.Id].IsPop = true
		Myapp.TopicLink[Myapp.Id].IsModify = false
	} else if choose[1] == str {
		Myapp.TopicLink[Myapp.Id].IsPop = false
		Myapp.TopicLink[Myapp.Id].IsModify = true
	} else {
		Myapp.TopicLink[Myapp.Id].IsPop = false
		Myapp.TopicLink[Myapp.Id].IsModify = false
	}
}

// SetSearchDia 制作查询警告
func SetSearchDia() {
	ErrText := fmt.Sprintf("查询的范围超过0 ~ %d的范围", len(Myapp.TopicLink))
	ErrDia := dialog.NewInformation("查询错误", ErrText, Myapp.Window)
	ErrDia.Resize(fyne.NewSize(100, 100))
	ErrDia.Show()
}

// SetExitDia 制作退出警告
func SetExitDia() {
	Exit := "你确定要退出程序吗?程序退出前一定要导出数据,否则所有修改的数据会丢失!"
	ExitDia := dialog.NewConfirm("退出", Exit, func(result bool) {
		if result {
			Myapp.App.Quit()
		}
	}, Myapp.Window)
	ExitDia.Resize(fyne.NewSize(100, 100))
	ExitDia.Show()
}

// OpenFileDiaSuccess 导入成功提醒
func OpenFileDiaSuccess() {
	ErrText := fmt.Sprintf("成功导入所有题目,题目的数量为%d", len(Myapp.TopicLink))
	ErrDia := dialog.NewInformation("导入成功", ErrText, Myapp.Window)
	ErrDia.Resize(fyne.NewSize(100, 100))
	ErrDia.Show()
}

// OpenFileDiaLose 导入失败提醒(文件无法找到
func OpenFileDiaLose() {
	ErrText := "导入题目失败,请检查需要审查题目文件夹下是否存在'入学考试题库.txt'文件"
	ErrDia := dialog.NewInformation("导入失败", ErrText, Myapp.Window)
	ErrDia.Resize(fyne.NewSize(100, 100))
	ErrDia.Show()
}

// SetDeriveFileDia 导出成功提醒
func SetDeriveFileDia(TopicNum int) {
	ErrText := fmt.Sprintf("成功导出所有题目,题目的数量为%d", TopicNum)
	ErrDia := dialog.NewInformation("导出成功", ErrText, Myapp.Window)
	ErrDia.Resize(fyne.NewSize(100, 100))
	ErrDia.Show()
}

// OpenPopDiaSuccess 导入成功提醒
func OpenPopDiaSuccess() {
	ErrText := fmt.Sprintf("成功找到所有需要删除题目,题目的数量为%d", len(Myapp.PopTopicId))
	ErrDia := dialog.NewInformation("删除题目", ErrText, Myapp.Window)
	ErrDia.Resize(fyne.NewSize(100, 100))
	ErrDia.Show()
}

// OpenPopDiaLose 导入失败提醒
func OpenPopDiaLose() {
	ErrText := fmt.Sprintf("无法找到任何删除题目")
	ErrDia := dialog.NewInformation("删除题目", ErrText, Myapp.Window)
	ErrDia.Resize(fyne.NewSize(100, 100))
	ErrDia.Show()
}

// OpenModifyDiaSuccess 导入成功提醒
func OpenModifyDiaSuccess() {
	ErrText := fmt.Sprintf("成功找到所有需要修改题目,题目的数量为%d", len(Myapp.ModifyTopicId))
	ErrDia := dialog.NewInformation("修改题目", ErrText, Myapp.Window)
	ErrDia.Resize(fyne.NewSize(100, 100))
	ErrDia.Show()
}

// OpenModifyDiaLose 导入失败提醒
func OpenModifyDiaLose() {
	ErrText := fmt.Sprintf("无法找到任何修改题目")
	ErrDia := dialog.NewInformation("修改题目", ErrText, Myapp.Window)
	ErrDia.Resize(fyne.NewSize(100, 100))
	ErrDia.Show()
}

// OpenAllDiaLose 导入失败提醒
func OpenAllDiaLose() {
	ErrText := fmt.Sprintf("无法找到任何题目")
	ErrDia := dialog.NewInformation("所有题目", ErrText, Myapp.Window)
	ErrDia.Resize(fyne.NewSize(100, 100))
	ErrDia.Show()
}

// AboutMeDia 关于我的联系方式
func AboutMeDia() {
	title := "本人联系方式"
	file, err := os.ReadFile("image/me.png")
	if err != nil {
		log.Println(err)
	}
	newIcon := widget.NewIcon(fyne.NewStaticResource("微信", file))
	newIcon.Resize(fyne.NewSize(300, 300))
	newIcon.Move(fyne.NewPos(30, 0))
	con := container.New(nil, newIcon)
	con.Resize(fyne.NewSize(400, 410))
	custom := dialog.NewCustom(title, "关闭", con, Myapp.Window)
	custom.Resize(fyne.NewSize(400, 410))
	custom.Show()
}

// SetModify 修改
func SetModify() {
	Myapp.TopicLink[Myapp.Id].IsPop = false
	Myapp.TopicLink[Myapp.Id].IsModify = true
	Myapp.InitPopAndModify()
}

// SetTip 操作提醒
func (Myapp *Config) SetTip() {
	TipButton := widget.NewButton("点击导入文件", func() {
		//文件读取
		Myapp.TopicLink = OpenData(filepath)
		if Myapp.TopicLink == nil {
			OpenFileDiaLose()
		} else {
			//导入成功提醒
			OpenFileDiaSuccess()
			//设置布局
			Myapp.SetWidget()
		}
	})
	TipButton.Resize(fyne.NewSize(100, 100))
	TipButton.Move(fyne.NewPos(450, 250))
	Myapp.Window.SetContent(container.New(nil, TipButton))
}

// LoadPopID 记录所有目前删除的题目ID
func (Myapp *Config) LoadPopID() {
	//清除以往数据
	Myapp.PopTopicId = make([]int, 0)
	//记录目前删除题目序号
	for i, Value := range Myapp.TopicLink {
		if Value.IsPop == true {
			Myapp.PopTopicId = append(Myapp.PopTopicId, i+1)
		}
	}
}

// LoadModifyID 记录所有目前修改的题目ID
func (Myapp *Config) LoadModifyID() {
	//清除以往数据
	Myapp.ModifyTopicId = make([]int, 0)
	//记录目前删除题目序号
	for i, Value := range Myapp.TopicLink {
		if Value.IsModify == true {
			Myapp.ModifyTopicId = append(Myapp.ModifyTopicId, i+1)
		}
	}
}
