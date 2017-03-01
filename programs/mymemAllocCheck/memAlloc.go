// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package main

import "fmt"

type TestStruct struct {
	id     int64
	serial int64
	leaves int32
	// serial1 int64
	// serial2 int64
	// serial3 int64
}

// func (t TestStruct) String() string {
// 	return fmt.Sprintf("%g - %g", t.id, t.serial)
// }

func main() {
	algOne()

}
func algOne() {
	var j []TestStruct
	//Populate
	popJstruct(&j)

}

func popJstruct(j *[]TestStruct) {
	v := *j
	for i := 0; i < 30; i++ {
		v = append(v, TestStruct{
			int64(i), int64(i + 1), int32(i + 2), //int64(i), //int64(i), //int64(i),
		})
		fmt.Printf("%+v\n", v[i])
	}
	// callPrinter(&v)
}

// func callPrinter(k *[]TestStruct) {
// 	v := *k
// 	for _, v := range v {
// 		// fmt.Printf("%v\n", v)
// 		println(v.String())
// 	}

// }
