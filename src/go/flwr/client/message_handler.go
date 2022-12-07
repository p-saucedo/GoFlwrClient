package goflwr

import (
	pb "goflwr/proto"

	serde "goflwr/src/go/flwr/common"
	"log"
)

func Handle(client ClientWrapper, server_msg *pb.ServerMessage) (*pb.ClientMessage, int, bool) {

	switch server_msg.ProtoReflect().WhichOneof(server_msg.ProtoReflect().Descriptor().Oneofs().ByName("msg")).JSONName() {

	case "reconnectIns":
		client_message, sleepDuration := _ReconnectIns(client, server_msg.GetReconnectIns())

		return &pb.ClientMessage{Msg: &pb.ClientMessage_DisconnectRes_{DisconnectRes: client_message}}, sleepDuration, false

	case "getParametersIns":
		client_message := _getParametersIns(client, server_msg.GetGetParametersIns())

		return &pb.ClientMessage{Msg: &pb.ClientMessage_GetParametersRes_{GetParametersRes: client_message}}, 0, true

	case "getPropertiesIns":
		client_message := _getPropertiesIns(client, server_msg.GetGetPropertiesIns())

		return &pb.ClientMessage{Msg: &pb.ClientMessage_GetPropertiesRes_{GetPropertiesRes: client_message}}, 0, true

	case "fitIns":
		client_message := _fitIns(client, server_msg.GetFitIns())

		return &pb.ClientMessage{Msg: &pb.ClientMessage_FitRes_{FitRes: client_message}}, 0, true

	case "evaluateIns":
		client_message := _evaluateIns(client, server_msg.GetEvaluateIns())

		return &pb.ClientMessage{Msg: &pb.ClientMessage_EvaluateRes_{EvaluateRes: client_message}}, 0, true

	default:
		log.Print("LIADITA")
		return &pb.ClientMessage{}, 0, true
	}

}

func _ReconnectIns(cw ClientWrapper, _ReconnectMsg *pb.ServerMessage_ReconnectIns) (*pb.ClientMessage_DisconnectRes, int) {

	reason := pb.Reason_ACK
	sleepDuration := 0

	if _ReconnectMsg.Seconds != 0 {
		reason = pb.Reason_RECONNECT
		sleepDuration = int(_ReconnectMsg.Seconds)
	}

	return &pb.ClientMessage_DisconnectRes{Reason: reason}, sleepDuration
}

func _getParametersIns(cw ClientWrapper, _getParametersMsg *pb.ServerMessage_GetParametersIns) *pb.ClientMessage_GetParametersRes {

	__getParametersIns := serde.GetParametersInsFromProto(_getParametersMsg)

	__getParametersRes := cw.GetParameters(__getParametersIns)

	return serde.GetParametersResToProto(__getParametersRes)
}

func _getPropertiesIns(cw ClientWrapper, _getPropertiessMsg *pb.ServerMessage_GetPropertiesIns) *pb.ClientMessage_GetPropertiesRes {

	__getPropertiesIns := serde.GetPropertiesInsFromProto(_getPropertiessMsg)

	__getPropertiesRes := cw.GetProperties(__getPropertiesIns)

	return serde.GetPropertiesResToProto(__getPropertiesRes)
}

func _fitIns(cw ClientWrapper, _fitInsMsg *pb.ServerMessage_FitIns) *pb.ClientMessage_FitRes {

	__fitIns := serde.FitInsFromProto(_fitInsMsg)

	__fitRes := cw.Fit(__fitIns)

	return serde.FitResToProto(__fitRes)
}

func _evaluateIns(cw ClientWrapper, _evaluateInsMsg *pb.ServerMessage_EvaluateIns) *pb.ClientMessage_EvaluateRes {

	__evaluateIns := serde.EvaluateInsFromProto(_evaluateInsMsg)

	__evaluateRes := cw.Evaluate(__evaluateIns)

	return serde.EvaluateResToProto(__evaluateRes)
}
