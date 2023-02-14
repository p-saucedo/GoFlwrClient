package goflwr

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	pb "goflwr/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(serverAddres string) (pb.FlowerService_JoinClient, context.CancelFunc, func() error) {
	conn, err := grpc.Dial(serverAddres, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	//defer conn.Close()

	c := pb.NewFlowerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()

	r, err := c.Join(ctx)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r, cancel, conn.Close

}

func StartClient(serverAddress string, client interface{}) error {
	c, h := client.(IClient)

	if !h {
		return errors.New("Not a Client object")
	}

	clientWrapper := &ClientWrapper{Client: c}

	var clientMessage *pb.ClientMessage = nil
	var sleepDuration int = 0
	var keepGoing bool = true

	for {

		conn, ctxCancel, connClose := Connect(serverAddress)

		for {
			serverMessage, err := conn.Recv()

			if err == io.EOF {
				return errors.New("End of file")
			}

			if err != nil {
				log.Println(err)
				log.Println("Failed to receive a message")
			}

			log.Printf("\n\nMessage received from server: %s\n", serverMessage)

			clientMessage, sleepDuration, keepGoing = Handle(*clientWrapper, serverMessage)

			log.Printf("\n\nClient message response: %s\n", clientMessage)

			err = conn.Send(clientMessage)

			if err != nil {
				ctxCancel()
				connClose()
				return errors.New("Bad idea")
			}

			if !keepGoing {
				break
			}

		}

		if sleepDuration == 0 {
			log.Println("INFO: Disconnect and shut down")
			break
		}

		log.Printf("INFO: Disconnect, then re-establish connection after %d second(s)\n", sleepDuration)

		time.Sleep(time.Duration(sleepDuration))

	}

	return nil
}
