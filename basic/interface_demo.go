package basic

import "fmt"

type newEr interface {
	New()
}

type testInterface interface {
	newEr

	// Done() <-chan struct{}
}

type kkTest struct {
	Name string
}

func NewTest() interface{} {

	return kkTest{Name: "1212"}

}

func DemoInterface() {

	kk := NewTest()

	i, ok := kk.(kkTest)

	fmt.Println("demo interface", i, ok)

	// ch := i.Done()

	// fmt.Println(ch)

}
