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
	}

	data, err := proto.Marshal(myBook)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)

	myNewBook := &Book{}
	err = proto.Unmarshal(data, myNewBook)
	if err != nil {
		log.Fatal("Unmarshaling error: ", err)
	}

	fmt.Println(myNewBook.GetName())
	fmt.Println(myNewBook.GetIsbn())

}
