package basic

import "fmt"

func DemoMap() {
	//map
	///* 声明变量，默认 map 是 nil */
	//var map_variable map[key_data_type]value_data_type
	///* 使用 make 函数 */
	//map_variable := make(map[key_data_type]value_data_type)

	demoMap := make(map[string]int) // var demoMap map[string]int   demoMap = make(map[string]int)

	demoMap["first"] = 1
	demoMap["second"] = 2
	demoMap["three"] = 3
	for k, v := range demoMap {
		fmt.Println(k, "首都是", demoMap[k])
		fmt.Println(k, "首都是", v)
	}
	/*查看元素在集合中是否存在 */
	capital, ok := demoMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	delete(demoMap, "France")

	var demoMapp map[string]FirstStruct
	demoMapp = make(map[string]FirstStruct)
	firstStruct := FirstStruct{2, "fs", "fs"}
	demoMapp["1"] = firstStruct
	d, ko := demoMapp["1"]
	if ko {
		fmt.Println("American 的首都是", d) //  {2 fs fs}
	} else {
		fmt.Println("American 的首都不存在")
	}

}

//func main(){
//	Demo()
//}
