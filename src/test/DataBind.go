package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"time"
)

func main() {
	os.Setenv("FYNE_FONT", "fonts/HanYiFangSongJian-1.ttf")
	myWindow := app.New()

	window := myWindow.NewWindow("数据绑定")
	window.Resize(fyne.Size{
		Height: 400,
		Width:  500,
	})
	//数据绑定
	str := "hello"
	num := 18
	bindString := binding.BindString(&str)
	bindInt := binding.BindInt(&num)
	log.Println(bindString.Get())
	log.Println(bindInt.Get())

	//绑定简单小部件(文字切换效果,利用通道和协程完成)
	newString := binding.NewString()
	newString.Set("效果演示")
	dialogue := []string{
		"大家好!",
		"我叫欧振贤",
		"这个是我写的小功能",
		"这个可以将对话一句句的呈现出来",
		"利用Go语言的协程和通信完成",
		"感谢大家的支持",
	}

	done := make(chan struct{})

	data := widget.NewLabelWithData(newString)
	go func() {
		for i := 0; i < len(dialogue); i++ {
			time.Sleep(time.Second)
			newString.Set(dialogue[i])
		}
		done <- struct{}{}
	}()
	//接收结束信息
	go func() {
		<-done
		newString.Set("感谢大家收看")
	}()
	//双向绑定
	s2 := binding.NewString()
	s2.Set("Hello!")
	withData := widget.NewLabelWithData(s2)
	entryWithData := widget.NewEntryWithData(s2)

	window.SetContent(data)
	window.SetContent(container.NewVBox(withData, entryWithData))

	//List数据
	//第一步先绑定数据
	list := binding.BindStringList(&[]string{
		"a",
		"b",
		"c",
	})
	//将绑定数据添加到组件中
	listWithData := widget.NewListWithData(list,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(item binding.DataItem, object fyne.CanvasObject) {
			object.(*widget.Label).Bind(item.(binding.String))
		})
	button := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Append %d", list.Length()+1)
		list.Append(val)
	})
	//数据添加到容器中
	window.SetContent(container.NewBorder(nil, button, nil, nil, listWithData))
	window.ShowAndRun()
}
