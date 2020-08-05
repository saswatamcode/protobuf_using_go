package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {

	myBook := &Book{
		Name: "Animal Farm",
		Isbn: 104,
		Author: &Author{
			Name:             "George Orwell",
			YearOfPublishing: 1945,
		},
	}

	// marshalling protocol buffer data
	data, err := proto.Marshal(myBook)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)

	myNewBook := &Book{}

	// unmarshalling encoded protocol buffer
	err = proto.Unmarshal(data, myNewBook)
	if err != nil {
		log.Fatal("Unmarshaling error: ", err)
	}

	// accessing fields with getter methods
	fmt.Println(myNewBook.GetName())
	fmt.Println(myNewBook.GetIsbn())
	fmt.Println(myNewBook.Author.GetName())
	fmt.Println(myNewBook.Author.GetYearOfPublishing())

}
