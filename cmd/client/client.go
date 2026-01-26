package main

import (
	"context"
	"log"

	exam "github.com/priority-infra/go_exam/internal/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// creating a connection to the server
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can't dial server: %v", err)
	}
	// creating grpc client for exam service
	c := exam.NewExamClient(conn)

	// getting an employee by ID
	d, err := c.GetEmployeeByID(context.Background(), &exam.ID{Value: 7})
	if err != nil {
		log.Fatalf("can't get employee %d: %v", 1, err)
	}
	log.Printf("the value for employee %d is %+v", 1, d)

	// getting a list of employees
	l, err := c.GetEmployeeList(context.Background(), &exam.Empty{})
	if err != nil {
		log.Fatalf("can't get employees: %v", err)
	}
	log.Printf("employees list: %+v", l)
}
