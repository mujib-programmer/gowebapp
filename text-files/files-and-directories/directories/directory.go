package main

import "fmt"
import "os"

func main() {
	os.Mkdir("mujibur", 0777)
	os.MkdirAll("mujibur/test1/test2", 0777)
	err := os.Remove("mujibur")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("mujibur")
}
