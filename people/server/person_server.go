package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "example.com/grpc-workout/people"
	"google.golang.org/grpc"
)

const (
	port = ":50050"
)

func newPeopleServer() *PersonServer {
	return &PersonServer{
		person_list: &pb.ListPeople{},
	}
}

type PersonServer struct {
	pb.UnimplementedPersonCRUDServer
	person_list *pb.ListPeople
}

func (server *PersonServer) Run() error {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterPersonCRUDServer(s, server)

	log.Printf("server listening at %v", lis.Addr())

	return s.Serve(lis)
}

func main() {
	var person_server *PersonServer = newPeopleServer()

	if err := person_server.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func (s *PersonServer) DeletePerson(ctx context.Context, in *pb.ID) (*pb.Response, error) {

	for i, v := range s.person_list.People {

		if v.Id == in.Id {
			s.person_list.People = append(s.person_list.People[:i], s.person_list.People[i+1:]...)

			response := &pb.Response{Response: "person removed"}

			return response, nil
		}
	}

	return &pb.Response{Response: "no such id"}, nil
}

func (s *PersonServer) Greeting(ctx context.Context, in *pb.Empty) (*pb.HelloRes, error) {

	res := pb.HelloRes{Response: "Hello how are you?"}

	return &res, nil

}

func (s *PersonServer) GetPerson(ctx context.Context, in *pb.ID) (*pb.Person, error) {

	for i, v := range s.person_list.People {
		if v.Id == in.Id {
			return s.person_list.People[i], nil
		}
	}
	return &pb.Person{Id: -1, Name: "not found", Age: 0}, nil
}

func (s *PersonServer) UpdatePerson(ctx context.Context, in *pb.Person) (*pb.Response, error) {

	for i, v := range s.person_list.People {
		if v.Id == in.Id {
			s.person_list.People[i] = in
			return &pb.Response{Response: "person updated"}, nil
		}
	}
	return &pb.Response{Response: "no person by such id"}, nil
}

func (s *PersonServer) ListAllPeople(ctx context.Context, in *pb.Empty) (*pb.ListPeople, error) {
	return s.person_list, nil
}

func (s *PersonServer) CreateNewPerson(ctx context.Context, in *pb.NewPerson) (*pb.Person, error) {
	log.Printf("Received: %v", in.GetName())
	var person_id int32 = int32(rand.Intn(10000))
	created_person := &pb.Person{Name: in.GetName(), Age: in.GetAge(), Id: person_id}
	s.person_list.People = append(s.person_list.People, created_person)

	return created_person, nil
}
