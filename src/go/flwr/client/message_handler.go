package goflwr

import (
	pb "goflwr/proto"

	serde "goflwr/src/go/flwr/common"
	"log"
)

func Handle(client Client, server_msg pb.ServerMessage) (pb.ClientMessage, int, bool) {

	switch server_msg.ProtoReflect().WhichOneof(server_msg.ProtoReflect().Descriptor().Oneofs().ByName("msg")).JSONName() {
	case "getParametersIns":
		client_message := _getParametersIns(client, *server_msg.GetGetParametersIns())
		return client_message, 0, true
	default:
		log.Print("LIADITA")
		return pb.ClientMessage{}, 0, true
	}

}

func _getParametersIns(client Client, get_parameters_msg pb.ServerMessage_GetParametersIns) pb.ClientMessage_GetParametersRes {

	get_parameters_ins := serde.GetParametersInsFromProto(get_parameters_msg)

	get_parameters_res := 

	get_parameters_res_proto := serde.GetParametersResToProto(get_get_parameters_res)

	return pb.ClientMessage_GetParametersRes{get_parameters_res: get_parameters_res}
}
