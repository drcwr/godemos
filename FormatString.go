package main

import (
"encoding/json"
"fmt"
"reflect"
"strconv"
)

func FormatString(iface interface{}) string {
    switch val := iface.(type) {
        case []byte:
        return string(val)
    }
    v := reflect.ValueOf(iface)
    fmt.Println(v.Kind())
    switch v.Kind() {
        case reflect.Invalid:
            return ""
        case reflect.Bool:
            return strconv.FormatBool(v.Bool())
        case reflect.String:
            return v.String()
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            return strconv.FormatInt(v.Int(), 10)
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
            return strconv.FormatUint(v.Uint(), 10)
        case reflect.Float64:
            return strconv.FormatFloat(v.Float(), 'f', -1, 64)
        case reflect.Float32:
            return strconv.FormatFloat(v.Float(), 'f', -1, 32)
        case reflect.Ptr:
            b, err := json.Marshal(v.Interface())
            if err != nil {
                return "nil"
            }
            return string(b)
    }
    return fmt.Sprintf("%v", iface)
}

func TestFormatString() {
    n := 100.001
    s := FormatString(n)
    fmt.Println(s)
    if s != "100.001" {
        fmt.Printf("%s", s)
    }
}

func TestMultiFormatString() {
    list := map[string]interface{}{
        // "10": 10,
        // "100": "100",
        // "100.001": 100.001,
        // "hello": "hello",
        // "[1 2 3]": []int{1, 2, 3},
        "[1 2 3]": []int{1, 2, 3},
        // "true": true,
        // "map[name:jason]": map[string]interface{}{"name": "jason"},
    }
    for k, v := range list {
        s := FormatString(v)
        fmt.Println("v",v,"s",s)
        if s != k {
            fmt.Printf("Error: %v to %s,but %s", v, k, s)
        }
    }

}

func main(){
    // TestFormatString()
    TestMultiFormatString()
}
