package main

import (
	"reflect"
	"fmt"
)

type T struct {
	A int
	B string
}

func main(){

	t := &T{1,"test string"}

	//反射对象类型
	tType := reflect.TypeOf(t)
	fmt.Println("t type is ",tType)

	//反射对象的值信息
	tValue := reflect.ValueOf(t).Elem()
	fmt.Println("tValue value is ",tValue)

	tValueType := tValue.Type()
	for i:=0;i<tValue.NumField();i++{
		value := tValue.Field(i)
		//fmt.Println(i,":",tValueType.Field(i).Name,value.Type(),value.Interface())
		fmt.Printf("%d:%s %s = %v\n",i,tValueType.Field(i).Name,value.Type(),value.Interface())
	}

}
