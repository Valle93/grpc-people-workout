package main

import (
	"context"
	"fmt"
	"log"

	pb "example.com/grpc-workout/people"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50050"
)

func main() {

	fmt.Println("hello main world!")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	fmt.Println("we continued")

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPersonCRUDClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	// var people = make(map[string]int32)

	var first string
	var name string
	var age int32
	var id int32

	for {

		fmt.Println("\nLet's use the grpc api, you can add, get, delete, update, list and echo")

		fmt.Scanln(&first)

		switch first {

		case "get":

			fmt.Println("id")
			fmt.Scanln(&id)

			r, err := c.GetPerson(ctx, &pb.ID{Id: id})
			if err != nil {
				log.Fatalf("Could not Update user: %v", err)
			}
			fmt.Printf("\nperson:\nName: %s\nAge: %d\nID: %d\n", r.GetName(), r.Age, r.Id)

		case "update":

			fmt.Println("id")
			fmt.Scanln(&id)
			fmt.Println("name:")
			fmt.Scanln(&name)
			fmt.Println("age:")
			fmt.Scanln(&age)

			r, err := c.UpdatePerson(ctx, &pb.Person{Id: id, Name: name, Age: age})
			if err != nil {
				log.Fatalf("Could not Update user: %v", err)
			}
			fmt.Printf("\n%v\n", r.Response)

		case "delete":

			fmt.Println("id")
			fmt.Scanln(&id)

			r, err := c.DeletePerson(ctx, &pb.ID{Id: id})
			if err != nil {
				log.Fatalf("Could not Delete user: %v", err)
			}
			fmt.Printf("\n%v\n", r.Response)

		case "add":

			fmt.Println("name:")
			fmt.Scanln(&name)
			fmt.Println("age:")
			fmt.Scanln(&age)

			r, err := c.CreateNewPerson(ctx, &pb.NewPerson{Name: name, Age: age})
			if err != nil {
				log.Fatalf("could not create user: %v", err)
			}
			fmt.Printf("added person:\nName: %s\nAge: %d\nID: %d", r.GetName(), r.Age, r.Id)

		case "list":

			r, err := c.ListAllPeople(ctx, &pb.Empty{})

			if err != nil {
				log.Fatalf("could not retrieve people: %v", err)
			}

			fmt.Println(r.String())

		case "echo":

			r, err := c.Greeting(ctx, &pb.Empty{})

			if err != nil {
				log.Fatalf("could not get a greeting: %v", err)
			}

			fmt.Println(r.Response)

		}
	}
}
