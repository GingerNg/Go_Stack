package basic

import (
	"errors"
	"fmt"
)

/* 定义接口 */
type FirstInterface interface {
	method1() (int, error)
	method2(int) (int, error)
	method3()
	method4(int) (int, error)
}

type FirstStruct struct {
	i    int
	ii   string
	name string
}

func (firstStruct FirstStruct) method1() (int, error) {
	return 1, nil
}

func (firstStruct FirstStruct) method2(s int) (int, error) {
	s = firstStruct.i + s
	println(firstStruct.name)
	println("method2", s)
	return s, nil
}

func (firstStruct FirstStruct) method4(s int) (int, error) {
	if firstStruct.i > s {
		return s, errors.New("input too big error")
	} else {
		return s, nil
	}
}

func DemoStruct() {
	//Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
	firstStruct := FirstStruct{2, "fs", "fs"}

	println(fmt.Sprint("firstStruct:", firstStruct))
	println(fmt.Sprintf("firstStruct: %+v", firstStruct))

	_, err := firstStruct.method1()
	if err != nil {
		print(err)
	}
	_, err = firstStruct.method2(1)
	if err != nil {
		print(err)
	}

	fs := new(FirstStruct) // default i=0  name=""
	_, err = fs.method2(10)
	if err != nil {
		print(err)
	}

	println("method4")
	fs1 := FirstStruct{2, "fs1", "fs1"}
	_, err = fs1.method4(3)
	if err != nil {
		print(err)
	}
	_, err = fs1.method4(0)
	if err != nil {
		print(err.Error())
	}

}
