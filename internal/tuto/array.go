package tuto

import (
	"fmt"
)

func createSlice() {
	var a [5]string
	//var b [0]string
	var c []string
	//var b = make([]string, 0)
	var d = make([]string, 0)
	e := &[5]string{"a", "b", "c"}

	//make tableau illimité

	fmt.Println(a) //[nil,nil,nil,nil,nil]
	a[2] = "Toto"
	fmt.Println(a) //[nil,nil,nil,nil,nil]
	//b[1] = "Toto" // impossible de créer la valeur
	c[2] = "Toto"  // [nil,nil,"Toto"]
	d[2] = "Toto"  // [nil,nil,"Toto"]
	fmt.Println(e) // ["a","b","c",nil,nil]

	//equivalent of the push in javascript
	d = append(d, "ok")
	fmt.Println(d)

	data := []string{"A", "B", "C", "D"}
	data = append(data, "Append Item")
	fmt.Println(data) // [A B C D]

	data2 := []string{"A", "B", "C", "D"}
	data2 = append([]string{"Prepend Item"}, data2...)
	//data2 = append([]string{"Prepend Item"}, "A", "B", "C", "D")
	fmt.Println(data2)
	// [Prepend Item A B C D]
}

func createMap() {

	var a map[int]string
	b := make(map[int]string, 5)
	c := make(map[int]string)

	fmt.Println(a, b, c, len(b))
}
