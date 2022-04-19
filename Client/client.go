package main

import (
	"google.golang.org/grpc"
	"log"
)

func Run() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure)
	if err != nil {
		log.Fatal("Could not dial Server: ", err)
	}

	defer conn.Close()

	//client :=
	//test

}
