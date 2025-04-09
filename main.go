package main

import (
	"log"
)

func main() {

	dbStore, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := dbStore.Init(); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v \n",dbStore)
	server := NewAPIServer(":3000", dbStore)
	server.Run()

}

// like interface give the methods the super power to make it dynamic to any type of structure ... like we will creating a method which can accept various types of structure to itself
// any = interface{}

//HOW THE DECODER WORKS
//frst we call the json package
//we take the NewDecoder function which takes the current io.Writer we are using
//then the Decode function which will decode the parameter we are passing to

//we may talk about middleware soon  here
