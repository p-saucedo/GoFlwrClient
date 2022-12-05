package goflwr

import (
	pb "goflwr/proto"

	serde "goflwr/src/go/flwr/common"
	"log"
)

func Handle(client ClientWrapper, server_msg *pb.ServerMessage) (*pb.ClientMessage, int, bool) {

	switch server_msg.ProtoReflect().WhichOneof(server_msg.ProtoReflect().Descriptor().Oneofs().ByName("msg")).JSONName() {

	case "getParametersIns":
		client_message := _getParametersIns(client, server_msg.GetGetParametersIns())

		return &pb.ClientMessage{Msg: &pb.ClientMessage_GetParametersRes_{GetParametersRes: client_message}}, 0, true

	case "fitIns":
		client_message := _fitIns(client, server_msg.GetFitIns())

		return &pb.ClientMessage{Msg: &pb.ClientMessage_FitRes_{FitRes: client_message}}, 0, true
	default:
		log.Print("LIADITA")
		return &pb.ClientMessage{}, 0, true
	}

}

func _getParametersIns(cw ClientWrapper, _getParametersMsg *pb.ServerMessage_GetParametersIns) *pb.ClientMessage_GetParametersRes {

	__getParametersIns := serde.GetParametersInsFromProto(_getParametersMsg)

	__getParametersRes := cw.GetParameters(__getParametersIns)

	return serde.GetParametersResToProto(__getParametersRes)
}

func _fitIns(cw ClientWrapper, _fitInsMsg *pb.ServerMessage_FitIns) *pb.ClientMessage_FitRes {

	__fitIns := serde.FitInsFromProto(_fitInsMsg)

	__fitRes := cw.Fit(__fitIns)

	return serde.FitResToProto(__fitRes)
}
