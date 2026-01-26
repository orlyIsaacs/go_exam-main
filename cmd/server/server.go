package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	exam "github.com/priority-infra/go_exam/internal/protos"
	"google.golang.org/grpc"
)

// Version represents the version of the service
var Version string = "development"

const Service = "exam"

// ListenHost Host and port to listen
var ListenHost string = "0.0.0.0:8000"

// Server this is a main object of the grpc server
var Server *grpc.Server

func main() {
	lis, err := net.Listen("tcp", ListenHost)
	if err != nil {
		log.Fatalf("version: %s, failed to listen: %v", Version, err)
	}
	Server = grpc.NewServer()
	exam.RegisterExamServer(Server, NewExamServer())

	log.Printf("starting to listen of %s, press ctrl+c to stop", ListenHost)
	go closeHandler()
	Server.Serve(lis)
}

func closeHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	Server.GracefulStop()
	log.Printf("exiting gracefully")
}
