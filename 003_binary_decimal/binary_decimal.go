// converts numbers from binary to decimal and vice versa

package main

import (
	"fmt"
	"strconv"
)

func main(){
	for{
		fmt.Println("Select converting mode number:")
		fmt.Println("1 - decimal -> binary")
		fmt.Println("2 - binary -> decimal")
		fmt.Println("3 - just exit")

		var mode int

		fmt.Scan(&mode)

		var userInput string

		switch mode  {
			case 1:
				fmt.Println("Enter decimal:")
				fmt.Scan(&userInput)
								

				parsedDecimal, err := strconv.ParseInt(userInput, 10, 64)
				if err != nil {
					fmt.Printf("Parsing error: %v\n", err)
					continue
				}
				fmt.Println(strconv.FormatInt(parsedDecimal, 2))
			case 2:
				fmt.Println("Enter binary:")
				fmt.Scan(&userInput)
				
				parsedBinary, err := strconv.ParseInt(userInput, 2, 64)
				if err != nil {
					fmt.Printf("Parsing error: %v\n", err)
					continue
				}
				fmt.Println(strconv.FormatInt(parsedBinary, 10))
			case 3:
				fmt.Println("Bye!")
				return
			default:
				fmt.Println("Enter mode")
		}
	}	
}