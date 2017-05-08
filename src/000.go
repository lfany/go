package main

import (
	"fmt"
	//"path/filepath"
	//"os"
	//"strings"
	//"path/filepath"
)

var  sentinels = []string {
	"aa",
	"bb",
	"cc",
}

var str interface{} = `// abcde
// hello
// world`

func main(){
	//for _, i := range sentinels {
	//	fmt.Println(i)
	//}

	//for _, i := range os.Environ() {
	//	a := strings.Split(i, "=")
	//
	//	for _, k := range a {
	//		//if k == "GOPATH" {
	//			fmt.Println(k)
	//
	//		//}
	//
	//	}
	//}

	//err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
	//	if ext := filepath.Ext(path); ext != ".go" && ext != "*.xx" {
	//		return nil
	//	}
	//	return nil
	//})
	//
	//if err != nil{
	//	fmt.Println("err", err)
	//}

	//fmt.Println(len(str))

	switch str.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	}

	//fmt.Println()

}