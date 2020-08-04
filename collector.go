package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var folders = [...]string {"/tmp/teste", "/tmp/teste2"}

func main() {

	listFolders := folders

	for i, s := range listFolders {

		fmt.Println("Scanning Folder")
		fmt.Println(i, s)

		var dirname = s

		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
			readAndCopy(dirname, file.Name())
		}
	}
}

func readAndCopy(folder string, file string) {
	fmt.Println(file)
	data, err := ioutil.ReadFile(folder + "/" + file)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	err = ioutil.WriteFile("dataBuffer/"+file, data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Copy file:", string(folder+"/"+file))
	removeFile(string(folder+"/"+file))
	return
}

func removeFile(file string){
	path := file
	err := os.Remove(path)

	fmt.Println("Removed file:", string(file))

	if err != nil {
		fmt.Println(err)
		return
	}
}
