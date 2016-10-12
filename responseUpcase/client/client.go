package main

import (
	pb "../responseUpcase"
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	addrFlag = flag.String("addr", ":5000", "Address host:post")
)

func main() {
	conn, err := grpc.Dial(*addrFlag, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	defer conn.Close()

	c := pb.NewConvertToUpcaseClient(conn)

	var word string

	fmt.Scan(&word)

	resp, err := c.ConvertToUpcase(context.Background(), &pb.Request{Word: word})

	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}
	log.Printf("%s is converted to : %s", word, resp.Result)
}
