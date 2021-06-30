// package main

// import "fmt"

// func main() {
//     var countryCapitalMap map[string]string /*创建集合 */
//     countryCapitalMap = make(map[string]string)

//     /* map插入key - value对,各个国家对应的首都 */
//     countryCapitalMap [ "France" ] = "Paris"
//     countryCapitalMap [ "Italy" ] = "罗马"
//     countryCapitalMap [ "Japan" ] = "东京"
//     countryCapitalMap [ "India " ] = "新德里"

//     /*使用键输出地图值 */ for country := range countryCapitalMap {
//         fmt.Println(country, "首都是", countryCapitalMap [country])
//     }

//     /*查看元素在集合中是否存在 */
//     captial, ok := countryCapitalMap [ "美国" ] /*如果确定是真实的,则存在,否则不存在 */
//     /*fmt.Println(captial) */
//     /*fmt.Println(ok) */
//     if (ok) {
//         fmt.Println("美国的首都是", captial)
//     } else {
//         fmt.Println("美国的首都不存在")
//     }
// }

// package main

// import "fmt"

// func main() {

// 	var countryCapitalMap map[string]string = map[string]string{"": ""}
// 	// countryCapitalMap = make(map[string]string)

// 	countryCapitalMap["France"] = "Paris"
// 	countryCapitalMap["Italy"] = "Roma"
// 	countryCapitalMap["Japan"] = "Tokyo"
// 	countryCapitalMap["India"] = "New Delhi"

// 	for country := range countryCapitalMap {
// 		fmt.Printf("%s,capital is %s\n", country, countryCapitalMap[country])
// 	}

// 	country := "America"
// 	captial, ok := countryCapitalMap[country]
// 	if ok {
// 		fmt.Printf("%s is %s\n", country, captial)
// 	} else {
// 		fmt.Printf("%s is not record\n", country)
// 	}

// }

package main

import "fmt"

func main() {
	// var countryCapital map[string]string
	// countryCapital = make(map[string]string)
	// countryCapital["America"] = "H"
	var numstr map[int]string = map[int]string{8: "eight"}
	_ = numstr
	countryCapital := map[string]string{"America": "HH"}

	for country := range countryCapital {
		fmt.Printf("%s capital is %s\n", country, countryCapital[country])
	}
}


package

import

const

type T struct {

}

func main() {
    var m  map[string]string
    m = make(map[string]string)
    m = map[string]string{"k":"v"}
    defer fmt.Println("ha")
    
    for k := range m
    return
}