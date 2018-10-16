package main

import "fmt"
import "os"

func main() {
	userFile := "mujibur.txt"
	err := os.Remove(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	fmt.Println(userFile + " file was removed.")
}
