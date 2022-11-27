package goflwr

import (
	pb "goflwr/proto"
	"log"
)

func Handle(client Client, server_msg pb.ServerMessage) {

	switch server_msg.ProtoReflect().WhichOneof(server_msg.ProtoReflect().Descriptor().Oneofs().ByName("msg")).JSONName() {
	case "getParametersIns":
		log.Print("EXITO")
	default:
		log.Print("LIADITA")
	}

}

func _getParametersIns(client Client) pb.ClientMessage {
	parameters := client.GetParameters()

}
