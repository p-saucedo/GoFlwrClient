package goflwr

import (
	"context"
	"log"
	"time"

	pb "goflwr/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect() pb.FlowerService_JoinClient {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	//defer conn.Close()

	c := pb.NewFlowerServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	//defer cancel()

	r, err := c.Join(ctx)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r

}
