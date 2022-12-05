package goflwr

import (
	pb "goflwr/proto"
	typing "goflwr/src/go/flwr/typing"
	"log"
)

func GetParametersInsFromProto(msg *pb.ServerMessage_GetParametersIns) *typing.GetParametersIns {
	config := ConfigFromProto(msg.GetConfig())

	return &typing.GetParametersIns{Config: config}
}

func GetParametersResToProto(msg *typing.GetParametersRes) *pb.ClientMessage_GetParametersRes {
	status := &pb.Status{Code: pb.Code(0), Message: "dalepa"}
	parameters := ParametersToProto(msg.Parameters)
	return &pb.ClientMessage_GetParametersRes{Status: status, Parameters: parameters}
}

func FitInsFromProto(msg *pb.ServerMessage_FitIns) *typing.FitIns {
	return &typing.FitIns{Parameters: ParametersFromProto(msg.Parameters), Config: ConfigFromProto(msg.Config)}
}

func ParametersFromProto(proto *pb.Parameters) typing.Parameters {
	return typing.Parameters{Tensors: proto.Tensors, TensorType: proto.TensorType}
}

func ParametersToProto(msg typing.Parameters) *pb.Parameters {
	return &pb.Parameters{Tensors: msg.Tensors, TensorType: msg.TensorType}
}

func PropertiesFromProto(proto map[string]*pb.Scalar) typing.Properties {
	properties := typing.Properties{}

	for key, value := range proto {
		properties[key] = ScalarFromProto(value)
	}

	return properties

}

func ConfigFromProto(proto map[string]*pb.Scalar) typing.Config {
	config := typing.Config{}

	for key, value := range proto {
		config[key] = ScalarFromProto(value)
	}

	return config

}

func ConfigToMap(config typing.Config) map[string]interface{} {

	mappedConfig := make(map[string]interface{}, len(config))

	for k, v := range config {
		mappedConfig[k] = v
	}

	return mappedConfig
}

func MapToMetrics(metrics map[string]interface{}) typing.Metrics {

	return &typing.Metrics(metrics)
}

/*func ScalarToProto(scalar interface{}) pb.Scalar {
	switch reflect.TypeOf(scalar) {
	case bool:
		return pb.Scalar{Scalar_Bool: scalar}
	}

}*/

func ScalarFromProto(scalar *pb.Scalar) int {
	log.Print(scalar.ProtoReflect().WhichOneof(scalar.ProtoReflect().Descriptor().Oneofs().ByName("scalar")).JSONName())
	return 1
}
