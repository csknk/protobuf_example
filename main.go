package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/csknk/protobuf_example"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := pb.Person{
		Id:    42,
		Name:  "Robert Alice",
		Email: "bob@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "111-4321", Type: pb.Person_HOME},
		},
	}
	book := &pb.AddressBook{}

	book.People = append(book.People, &p)

	fname := "address_book"

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the address book back in...
	readBook(fname)

}

func readBook(fileName string) {
	rawData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(rawData, book); err != nil {
		log.Fatalln("Failed to read address book:", err)
	}
	for _, p := range book.People {
		w := os.Stdout
		writePerson(w, p)
	}
}

func writePerson(w io.Writer, p *pb.Person) {
	fmt.Fprintf(w, "ID:%d\n", p.Id)
	fmt.Fprintf(w, "Name:%s\n", p.Name)
	fmt.Fprintf(w, "Email:%s\n", p.Email)

}
