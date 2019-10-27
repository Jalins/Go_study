package main

import "fmt"

type link struct {
	data string
	nextLink *link
}

func main() {



	genesis := &link{"data_struct", nil}

	secondLink := &link{"jalins", genesis}



	fmt.Println(secondLink.nextLink)

}
