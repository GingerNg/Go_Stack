package basic

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"os"
)

type PersonInfo struct {
	Name    string
	age     int32
	Sex     bool
	Hobbies []string
}

func DemoJson() {
	personInfo := []PersonInfo{{"David", 30, true, []string{"跑步", "读书", "看电影"}}, {"Lee", 27, false, []string{"工作", "读书", "看电影"}}}
	dump("./person.json", personInfo)

	info := load("./person.json")

	var person []PersonInfo
	//将 map 转换为指定的结构体
	if err := mapstructure.Decode(info, &person); err != nil {
		println(err)
	}
	println(person)

	//res := info.([]PersonInfo)
	//println(res)
}

func load(filePath string) interface{} {
	filePtr, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return nil
	}
	defer filePtr.Close()

	//var person []PersonInfo
	var info interface{}

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())

	} else {
		fmt.Println("Decoder success")
		fmt.Println(info)
	}
	return info
}

func dump(filePath string, info interface{}) {

	// 创建文件
	filePtr, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)

	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("Encoder failed", err.Error())

	} else {
		fmt.Println("Encoder success")
	}

}
