package main

import (
	"fmt"
	"os"
)

func makeDir(dirPath string) {
	//os.Mkdir(dirPath, 0777)
	os.MkdirAll(dirPath, 0777)
}

func main() {
	var name string
	fmt.Println("Enter name of project: ")
	fmt.Scanln(&name)
	dirPath := "./" + name

	err := os.Mkdir(dirPath, 0755) // Create the directory with read-write-execute permissions

	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

}
