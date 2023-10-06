package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

// TopicStruct 题目结构定义
type TopicStruct struct {
	//题目序号
	Id int
	//分数
	Point string
	//题目
	Topic string
	//选项
	Choice []string
	//删除
	IsPop bool
	//修改
	IsModify bool
}

func CreatLink() []*TopicStruct {
	TopicLink := make([]*TopicStruct, 0)
	return TopicLink
}

// Init_Topic 创建题目
func Init_Topic(id int, Point string, Topic string, Choice []string, IsPop bool, IsModify bool) *TopicStruct {
	return &TopicStruct{
		Id:       id,
		Point:    Point,
		Topic:    Topic,
		Choice:   Choice,
		IsPop:    IsPop,
		IsModify: IsModify,
	}
}

// OpenData 导入数据
func OpenData(filename string) []*TopicStruct {
	//创建存储结构
	TopicLink := CreatLink()
	//题目Id
	var NumId = 1
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()
	//扫描器
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	//取出数据
	for {
		//获取数据
		topic, end := ReadFile(scanner, NumId)
		if end == false {
			break
		}
		NumId++
		TopicLink = append(TopicLink, topic)
	}
	return TopicLink
}

// ReadFile 读取文件数据
func ReadFile(scanner *bufio.Scanner, id int) (*TopicStruct, bool) {
	//创建题目变量
	var Point string
	var Topic string
	Choice := make([]string, 4)
	var count int
	var flag bool
	for scanner.Scan() {
		if scanner.Text() == "" {
			flag = true
			break
		}
		switch count {
		case 0:
			Point = scanner.Text()
		case 1:
			Topic = scanner.Text()
		case 2, 3, 4, 5:
			Choice[count-2] = scanner.Text()
		}
		count++
		flag = false
	}
	if flag == true {
		topic := Init_Topic(id, Point, Topic, Choice, false, false)
		return topic, true
	} else {
		return nil, false
	}
}

// DeriveData 导出数据
func DeriveData(TopicLink []*TopicStruct) int {
	var count int
	//创建文件
	openFile, err := os.OpenFile(DeriveFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("错误:", err)
		return 0
	}
	defer openFile.Close()
	//遍历链表
	for i := range TopicLink {
		if TopicLink[i].IsPop == false {
			count++
			var buf bytes.Buffer
			buf.WriteString(TopicLink[i].Point + "\n")
			buf.WriteString(TopicLink[i].Topic + "\n")
			for j := range TopicLink[i].Choice {
				if TopicLink[i].Choice[j] != "" {
					buf.WriteString(TopicLink[i].Choice[j] + "\n")
				}
			}
			buf.WriteString("\n")
			_, err := openFile.WriteString(buf.String())
			if err != nil {
				log.Println("错误:", err)
			}
		}
	}
	return count
}
