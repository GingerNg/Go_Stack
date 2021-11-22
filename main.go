package main

import (
	"Go_Stack/basic"
	"Go_Stack/tools"
	"Go_Stack/utils"
	"fmt"
)

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")

	// ***************************** basic ****************
	basic.DemoInterface()

	//循环语句

	basic.DemoMap()
	basic.DemoSlice()
	basic.DemoStruct()
	basic.DemoGoroutine()

	utils.DemoString()

	fmt.Println("Hello, World!")
	//utils.DemoEs()

	//tools.ESMain()

	basic.DemoSwitch()

	var md = map[string]interface{}{}
	md["1"] = "2"
	//println(md["1"].(type))
	basic.DemoJson()

	basic.HttpDemo()

	// ****************************** tools **********************
	tools.DemoSqlite()
}
