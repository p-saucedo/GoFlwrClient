package goflwr

import (
	pb "goflwr/proto"
	typing "goflwr/src/go/flwr/typing"
)

func GetParametersInsFromProto(msg *pb.ServerMessage_GetParametersIns) *typing.GetParametersIns {
	config := ConfigFromProto(msg.GetConfig())

	return &typing.GetParametersIns{Config: config}
}

func GetParametersResToProto(msg *typing.GetParametersRes) *pb.ClientMessage_GetParametersRes {
	status := &pb.Status{Code: pb.Code_OK, Message: "Success"}
	parameters := ParametersToProto(msg.Parameters)
	return &pb.ClientMessage_GetParametersRes{Status: status, Parameters: parameters}
}

func FitInsFromProto(msg *pb.ServerMessage_FitIns) *typing.FitIns {
	return &typing.FitIns{Parameters: ParametersFromProto(msg.Parameters), Config: ConfigFromProto(msg.Config)}
}

func FitResToProto(msg *typing.FitRes) *pb.ClientMessage_FitRes {
	status := &pb.Status{Code: pb.Code_OK, Message: "Success"}
	parameters := ParametersToProto(msg.Parameters)
	metrics := MetricsToProto(msg.Metrics)

	return &pb.ClientMessage_FitRes{Status: status, Parameters: parameters, NumExamples: int64(msg.NumExamples), Metrics: metrics}
}

func EvaluateInsFromProto(msg *pb.ServerMessage_EvaluateIns) *typing.EvaluateIns {
	return &typing.EvaluateIns{Parameters: ParametersFromProto(msg.Parameters), Config: ConfigFromProto(msg.Config)}
}

func EvaluateResToProto(msg *typing.EvaluateRes) *pb.ClientMessage_EvaluateRes {
	status := &pb.Status{Code: pb.Code_OK, Message: "Success"}
	metrics := MetricsToProto(msg.Metrics)

	return &pb.ClientMessage_EvaluateRes{Status: status, Loss: msg.Loss, NumExamples: int64(msg.NumExamples), Metrics: metrics}
}

func ParametersFromProto(proto *pb.Parameters) typing.Parameters {
	return typing.Parameters{Tensors: proto.Tensors, TensorType: proto.TensorType}
}

func ParametersToProto(msg typing.Parameters) *pb.Parameters {
	return &pb.Parameters{Tensors: msg.Tensors, TensorType: msg.TensorType}
}

func MetricsToProto(msg typing.Metrics) map[string]*pb.Scalar {
	protoMetrics := make(map[string]*pb.Scalar, len(msg))

	for k, v := range msg {
		protoMetrics[k] = ScalarToProto(v)
	}

	return protoMetrics
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

	typeMetrics := make(map[string]typing.Scalar, len(metrics))

	for k, v := range metrics {
		typeMetrics[k] = v
	}

	return typing.Metrics(typeMetrics)
}

func ScalarFromProto(scalar *pb.Scalar) typing.Scalar {
	switch scalar.ProtoReflect().WhichOneof(scalar.ProtoReflect().Descriptor().Oneofs().ByName("scalar")).JSONName() {
	case "sint64", "double", "bool", "bytes", "string":
		return typing.Scalar(scalar)
	default:
		return nil
	}
}

func ScalarToProto(sc typing.Scalar) *pb.Scalar {
	switch sc.(type) {
	case int:
		return &pb.Scalar{Scalar: &pb.Scalar_Sint64{Sint64: sc.(int64)}}
	case string:
		return &pb.Scalar{Scalar: &pb.Scalar_String_{String_: sc.(string)}}
	case float64, float32:
		return &pb.Scalar{Scalar: &pb.Scalar_Double{Double: sc.(float64)}}
	case bool:
		return &pb.Scalar{Scalar: &pb.Scalar_Bool{Bool: sc.(bool)}}
	case byte:
		return &pb.Scalar{Scalar: &pb.Scalar_Bytes{Bytes: sc.([]byte)}}
	default:
		return &pb.Scalar{Scalar: &pb.Scalar_Bytes{Bytes: sc.([]byte)}}
	}

}
