package goflwr

import (
	pb "goflwr/proto"

	serde "goflwr/src/go/flwr/common"
	"log"
)

func Handle(client ClientWrapper, server_msg pb.ServerMessage) (*pb.ClientMessage, int, bool) {

	switch server_msg.ProtoReflect().WhichOneof(server_msg.ProtoReflect().Descriptor().Oneofs().ByName("msg")).JSONName() {
	case "getParametersIns":
		client_message := _getParametersIns(client, server_msg.GetGetParametersIns())
		return &pb.ClientMessage{Msg: &pb.ClientMessage_GetParametersRes_{GetParametersRes: client_message}}, 0, true
	default:
		log.Print("LIADITA")
		return &pb.ClientMessage{}, 0, true
	}

}

func _getParametersIns(cw ClientWrapper, get_parameters_msg *pb.ServerMessage_GetParametersIns) *pb.ClientMessage_GetParametersRes {

	get_parameters_ins := serde.GetParametersInsFromProto(get_parameters_msg)

	get_parameters_res := cw.GetParameters(get_parameters_ins)

	get_parameters_res_proto := serde.GetParametersResToProto(get_parameters_res)

	return get_parameters_res_proto
}
