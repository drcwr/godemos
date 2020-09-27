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

	tType = reflect.TypeOf(t.A)
	fmt.Println("t.A type is ",tType)

	tType2 := reflect.TypeOf(tType)
	fmt.Println("tType type is ",tType2)

	//反射对象的值信息
	tValue := reflect.ValueOf(t).Elem()
	fmt.Println("tValue value is ",tValue)

	tValueType := tValue.Type()
	for i:=0;i<tValue.NumField();i++{
		value := tValue.Field(i)
		//fmt.Println(i,":",tValueType.Field(i).Name,value.Type(),value.Interface())
		fmt.Printf("%d:%s %s = %v\n",i,tValueType.Field(i).Name,value.Type(),value.Interface())
	}

	num0,err := reterr()
	fmt.Printf("\nreterr ret err:%#v,%d\n",err,num0)

	et := reflect.TypeOf(err)
	fmt.Println("reterr err type is ",et)
	if err !=nil {
		fmt.Printf("reterr err != nil, err:%#v,nil:%#v\n",err,nil)
	} else {
		fmt.Printf("reterr err == nil, err:%#v,nil:%#v\n",err,nil)
	}



	num,err := test()

	et = reflect.TypeOf(err)
	fmt.Println("\ntest err type is ",et)

	fmt.Printf("test ret err:%#v,%d\n",err,num)
	if err !=nil {
		fmt.Printf("main err != nil, err:%#v,nil:%#v\n",err,nil)
	} else {
		fmt.Printf("main err == nil, err:%#v,nil:%#v\n",err,nil)
	}



	num2,err2 := test()
	fmt.Printf("\n22 test ret err:%#v,%d\n",err2,num2)
	et = reflect.TypeOf(err2)
	fmt.Println("err2 type is ",et)
	if err2 != nil {
		fmt.Printf("main err2!=nil, err2:%#v,nil:%#v\n",err2,nil)
	} else {
		fmt.Printf("main err2==nil err2:%#v,nil:%#v\n",err2,nil)
	}

	if err2 != err {
		fmt.Printf("err main err2 != err, err2:%#v,err:%#v\n",err2,err)
	} else {
		fmt.Printf("err main err2 == err, err2:%#v,err:%#v\n",err2,err)
	}

	if err != nil {
		fmt.Printf("main err != nil, err:%#v,nil:%#v\n",err,nil)
	} else {
		fmt.Printf("main err == nil, err:%#v,nil:%#v\n",err,nil)
	}

	var i *int
	fmt.Printf("\ntest ret i:%#v\n",i)
	if i !=nil {
		fmt.Printf("main i!=nil, err:%#v,nil:%#v\n",i,nil)
	} else {
		fmt.Printf("test i==nil err:%#v,nil:%#v\n",i,nil)
	}

	var tp *T
	fmt.Printf("\ntest ret tp:%#v\n",tp)
	if tp !=nil {
		fmt.Printf("main tp!=nil, tp:%#v,nil:%#v\n",tp,nil)
	} else {
		fmt.Printf("test tp==nil tp:%#v,nil:%#v\n",tp,nil)
	}



	if err2 == nil {
		if err2 == err {
			if err == nil {
				fmt.Printf("========= err == nil, err:%#v, err2:%#v,nil:%#v\n",err,err2,nil)
			} else {
				fmt.Printf("========= err != nil, err:%#v, err2:%#v,nil:%#v\n",err,err2,nil)
			}
		}
	}

	var err3 error
	err3 = err2
	// err2 = err3
	if err3 == nil {
		if err3 == err {
			if err == nil {
				fmt.Printf("3========= err == nil, err:%#v, err2:%#v,nil:%#v\n",err,err2,nil)
			} else {
				fmt.Printf("3========= err != nil, err:%#v, err2:%#v,nil:%#v\n",err,err2,nil)
			}
		}
	} else {
		fmt.Printf("========= err3 != nil, err:%#v, err3:%#v,nil:%#v\n",err,err3,nil)
	}

	if err2 == err3 {
		fmt.Printf("23========= err2 == err3, err3:%#v, err2:%#v,nil:%#v\n",err3,err2,nil)
	} else {
		fmt.Printf("23========= err2 != err3, err3:%#v, err2:%#v,nil:%#v\n",err3,err2,nil)
	}

	err4 := err2
@  }

func (t *T) Error() string {
	fmt.Printf("T Error\n")
	return "T error"
}

func test()(int,*T) {
	var ret *T

	return 2,ret
}

func reterr()(int,error) {
	return 3,nil
}