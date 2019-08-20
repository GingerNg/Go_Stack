package utils

import (
	"fmt"
	"strings"
)

func DemoString() {
	s := "http://192.168.1.205:8001/convert/pdf"
	ss := strings.Split(s, "/")
	for _, v := range ss {
		println(v)
	}

	str := "XBodyContentX"
	content := str[1 : len(str)-1]
	fmt.Println(content)

}
