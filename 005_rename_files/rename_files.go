package main

import (
	"fmt"
	"math/rand"
	"os"
)



var NEW_NAME_LENGTH = 20;
var NEW_NAME_SYMBOLS_SET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func createRandomName() string {
		var newName string;		

		for i := 0; i < NEW_NAME_LENGTH; i++ {
			randIndex := rand.Intn(len(NEW_NAME_SYMBOLS_SET))
			newName += string(NEW_NAME_SYMBOLS_SET[randIndex])
		}

		return newName
}

func main() {	
	if len(os.Args) < 2{
		fmt.Println("Need path-dir argument!")
		return
	};

	path := os.Args[1];

	fileInfo, err := os.Stat(path);
	if(err != nil){
		fmt.Printf("%s is not a path", path);
		return
	}	

	if(!fileInfo.IsDir()){
		fmt.Printf("%s is not a directory", fileInfo)
		return	
	}

	filesList, err := os.ReadDir(path)
	if(err != nil){
		fmt.Printf("Error while reading %s", path)
	}

	namesMap := make(map[string]string)

	for _, file := range filesList {
		fmt.Println(file)
		newName := createRandomName();
		namesMap[file.Name()] = newName
	}


	fmt.Println(namesMap)


	

	




}