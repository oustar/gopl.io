package main

import (
	"fmt"
	"reflect"
)

type bailout struct {
	name string
}

func retvalue(a string) {
	panic(bailout{a})
}
func main() {

	defer func() {
		p := recover()
		v := reflect.ValueOf(p)
		switch v.Kind() {
		case reflect.Struct:
			fmt.Println(v.Field(0))
		}
	}()

	retvalue("hello ")
	fmt.Println("world")

}
