package goflwr

import (
	"context"
	"log"
	"time"

	pb "goflwr/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(serverAddres string) pb.FlowerService_JoinClient {
	conn, err := grpc.Dial(serverAddres, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

func StartClient(serverAddress string, client interface{}) {
	c, h := client.(IClient)

	if !h {
		panic("Not a Client type")
	}

	clientWrapper := &ClientWrapper{Client: c}

	var clientMessage *pb.ClientMessage = nil
	var sleepDuration int = 0
	var keepGoing bool = true

	conn := Connect(serverAddress)

	for {

		for {
			serverMessage, err := conn.Recv()

			if err != nil {
				panic(err)
			}

			clientMessage, sleepDuration, keepGoing = Handle(*clientWrapper, *serverMessage)

			err = conn.Send(clientMessage)

			log.Print(clientMessage)
			if err != nil {
				panic(err)
			}

			if !keepGoing {
				break
			}

		}

		if sleepDuration == 0 {
			log.Println("INFO: Disconnect and shut down")
			break
		}

		log.Printf("INFO: Disocnnect, then re-establish connection after %d second(s)\n", sleepDuration)

		time.Sleep(time.Duration(sleepDuration))

	}
}
