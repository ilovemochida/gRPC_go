package main

import (
	"flag"
	"golang.org/x/net/context"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	pb "../responseUpcase"
)

var (
	addrFlag = flag.String("addr", ":5000", "Address host:post")
)

type server struct{}

func (s *server) ConvertToUpcase(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Result: strings.ToUpper(in.Word)}, nil
}

func main() {
	lis, err := net.Listen("tcp", *addrFlag)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterConvertToUpcaseServer(s, &server{})
	s.Serve(lis)
}
