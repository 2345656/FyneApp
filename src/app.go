package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"os"
)

type Config struct {
	App           fyne.App       //应用主体
	Window        fyne.Window    //应用界面
	Layout        fyne.Layout    //应用布局
	TopicLink     []*TopicStruct //题目存储
	Id            int            //题目ID
	EnWrite       bool           //是否写入
	ShowPop       bool           //只显示删除题目
	ShowModify    bool           //只显示修改题目
	PopTopicId    []int          //删除题目ID
	ModifyTopicId []int          //修改题目ID
}

var Myapp = &Config{
	PopTopicId:    make([]int, 0),
	ModifyTopicId: make([]int, 0),
}

func main() {
	//文体设计
	os.Setenv("FYNE_FONT", "fonts/HanYiFangSongJian-1.ttf")
	//os.Setenv("FYNE_FONT", "fonts/MaoKenWangXingYuan-2.ttf")
	//os.Setenv("FYNE_FONT", "fonts/HanYiCuFangSongJian-1.ttf")
	//程序窗口创建
	Myapp.App = app.New()
	Myapp.Window = Myapp.App.NewWindow("优课审题脚本")
	//退出提醒
	Myapp.Window.SetCloseIntercept(SetExitDia)
	//设置界面
	Myapp.SetWindow()
	//菜单设计
	Myapp.SetMenu()
	//主页选项
	Myapp.SetTip()
	//结果运行
	Myapp.Window.ShowAndRun()
}
