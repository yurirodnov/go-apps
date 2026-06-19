package main

import (
	"fmt"
	"os"
)

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

	fmt.Println("Ok")

	//fmt.Printf("Yor arg is %s", os.Args[1]);




}