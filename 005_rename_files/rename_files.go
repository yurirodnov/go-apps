package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

var NEW_NAME_LENGTH = 20;
var NEW_NAME_SYMBOLS_SET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func createRandomName() string {
		var newName strings.Builder;		

		for i := 0; i < NEW_NAME_LENGTH; i++ {
			randIndex := rand.Intn(len(NEW_NAME_SYMBOLS_SET))
			newName.WriteByte(NEW_NAME_SYMBOLS_SET[randIndex])
		}

		return newName.String()
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

	fmt.Printf("Prcessing files in %s\n", path)

	namesMap := make(map[string]string)
	
	for _, file := range filesList {
		fmt.Printf("Renaming file %s \n", file.Name())
		
		
		fileExt := filepath.Ext(file.Name())
		newName := createRandomName();
		oldPath := filepath.Join(path, file.Name())
		newPath := filepath.Join(path, newName + fileExt)


		namesMap[file.Name()] = newName
		

		err := os.Rename(oldPath, newPath)

		if(err != nil){
			log.Fatal(err)
		}

		fmt.Printf("Renamed: %s -> %s \n", file.Name(), newName + fileExt)
		
		// fmt.Println(oldPath)
		// fmt.Println(newPath)
		// fmt.Println("-----")
	}


	


	

	




}