package main

import (
	"Tasks/findian"
	"Tasks/slice"

	//"Tasks/makejson"
	"fmt"
)

func Slicetest() {

	y := slice.Slice()
	fmt.Println(y)
}

func Findiantest() {

	x := findian.Findian()
	if x {
		fmt.Printf("Found!")
	} else {

		fmt.Printf("Not Found!")
	}
}

/*func makeJsonTest() {
	x := makejson.makeJSON()
	fmt.Print(x)
}*/
func main() {
	//Findiantest()
	//Slicetest()
	//makeJsonTest()

}
