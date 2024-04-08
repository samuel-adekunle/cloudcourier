package main

import (
	"fmt"
	"reflect"
)

func main() {

	type Value struct{}
	s := &Value{}
	v := reflect.TypeOf(s)
	fmt.Println(v.Elem().Kind(), "1")
	fmt.Println(v.Elem().Name(), "2")
}
