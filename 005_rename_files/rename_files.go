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
		fmt.Printf("%s is not a path", path)
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

	fmt.Printf("Processing files in %s\n", path)

	usedNames := make(map[string]bool)
	
	for _, file := range filesList {
		if file.IsDir(){
			continue
		}

		fmt.Printf("Renaming file %s \n", file.Name())		
		
		fileExt := filepath.Ext(file.Name())

		var newName string;
		for{
			candidate := createRandomName() + fileExt
			
			if !usedNames[candidate]{
				newName = candidate
				break
			}
		}
		
		oldPath := filepath.Join(path, file.Name())
		newPath := filepath.Join(path, newName)

		usedNames[newName] = true		

		err := os.Rename(oldPath, newPath)

		if(err != nil){
			log.Fatal(err)
		}

		fmt.Printf("Renamed: %s -> %s \n", file.Name(), newName)
		
	}
}