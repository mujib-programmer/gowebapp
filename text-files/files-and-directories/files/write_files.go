package main

import "fmt"
import "os"

func main() {
	userFile := "mujibur.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}
}
