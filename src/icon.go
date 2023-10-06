package main

import (
	"fyne.io/fyne/v2"
	"io/ioutil"
	"log"
)

func SetIcon() *fyne.StaticResource {
	//读取图片
	iconContent, err := ioutil.ReadFile("image\\all.ico")
	if err != nil {
		log.Println(err)
	}
	resource := fyne.StaticResource{
		StaticName:    "欧振贤",
		StaticContent: iconContent,
	}
	return &resource
}
