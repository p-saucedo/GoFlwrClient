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

func GetPropertiesInsFromProto(msg *pb.ServerMessage_GetPropertiesIns) *typing.GetPropertiesIns {
	config := ConfigFromProto(msg.GetConfig())

	return &typing.GetPropertiesIns{Config: config}
}

func GetPropertiesResToProto(msg *typing.GetPropertiesRes) *pb.ClientMessage_GetPropertiesRes {
	status := &pb.Status{Code: pb.Code_OK, Message: "Success"}
	properties := PropertiesToProto(msg.Properties)
	return &pb.ClientMessage_GetPropertiesRes{Status: status, Properties: properties}
}

func FitInsFromProto(msg *pb.ServerMessage_FitIns) *typing.FitIns {
	return &typing.FitIns{Parameters: ParametersFromProto(msg.Parameters), Config: ConfigFromProto(msg.Config)}
}

func FitResToProto(msg *typing.FitRes) *pb.ClientMessage_FitRes {
	status := &pb.Status{Code: pb.Code_OK, Message: "Success"}
	parameters := ParametersToProto(msg.Parameters)
	metrics := MetricsToProto(msg.Metrics)

	return &pb.ClientMessage_FitRes{Status: status, Parameters: parameters, NumExamples: msg.NumExamples, Metrics: metrics}
}

func EvaluateInsFromProto(msg *pb.ServerMessage_EvaluateIns) *typing.EvaluateIns {
	return &typing.EvaluateIns{Parameters: ParametersFromProto(msg.Parameters), Config: ConfigFromProto(msg.Config)}
}

func EvaluateResToProto(msg *typing.EvaluateRes) *pb.ClientMessage_EvaluateRes {
	status := &pb.Status{Code: pb.Code_OK, Message: "Success"}
	metrics := MetricsToProto(msg.Metrics)

	return &pb.ClientMessage_EvaluateRes{Status: status, Loss: msg.Loss, NumExamples: msg.NumExamples, Metrics: metrics}
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

func PropertiesToProto(msg typing.Properties) map[string]*pb.Scalar {
	protoProperties := make(map[string]*pb.Scalar, len(msg))

	for k, v := range msg {
		protoProperties[k] = ScalarToProto(v)
	}

	return protoProperties
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

func MapToProperties(metrics map[string]interface{}) typing.Properties {

	typeProperties := make(map[string]typing.Scalar, len(metrics))

	for k, v := range metrics {
		typeProperties[k] = v
	}

	return typing.Properties(typeProperties)
}

func ScalarFromProto(scalar *pb.Scalar) typing.Scalar {
	switch scalar.ProtoReflect().WhichOneof(scalar.ProtoReflect().Descriptor().Oneofs().ByName("scalar")).JSONName() {
	case "sint64":
		return scalar.GetSint64()
	case "bool":
		return scalar.GetBool()
	case "double":
		return scalar.GetDouble()
	case "bytes":
		return scalar.GetBytes()
	case "string":
		return scalar.GetString_()
	default:
		return nil
	}
}

func ScalarToProto(sc typing.Scalar) *pb.Scalar {

	switch scalar := sc.(type) {
	case int:
		return &pb.Scalar{Scalar: &pb.Scalar_Sint64{Sint64: int64(scalar)}}
	case string:
		return &pb.Scalar{Scalar: &pb.Scalar_String_{String_: scalar}}
	case float64:
		return &pb.Scalar{Scalar: &pb.Scalar_Double{Double: float64(scalar)}}
	case bool:
		return &pb.Scalar{Scalar: &pb.Scalar_Bool{Bool: scalar}}
	case byte:
		return &pb.Scalar{Scalar: &pb.Scalar_Bytes{Bytes: sc.([]byte)}}
	default:
		return &pb.Scalar{Scalar: &pb.Scalar_Bytes{Bytes: sc.([]byte)}}
	}

}
