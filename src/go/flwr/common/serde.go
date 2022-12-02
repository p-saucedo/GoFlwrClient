package goflwr

import (
	pb "goflwr/proto"
	"log"
)

func GetParametersInsFromProto(msg pb.ServerMessage_GetParametersIns) GetParametersIns {
	config := ConfigFromProto(msg.GetConfig())

	return GetParametersIns{config: config}
}

func PropertiesFromProto(proto map[string]*pb.Scalar) Properties {
	properties := Properties{}

	log.Printf("Properties proto: %s", proto)
	for key, value := range proto {
		properties[key] = ScalarFromProto(value)
	}

	return properties

}

func ConfigFromProto(proto map[string]*pb.Scalar) Config {
	config := Config{}
	log.Printf("Config proto: %s", proto)
	for key, value := range proto {
		config[key] = ScalarFromProto(value)
	}
	log.Printf("Config from proto: %s", config)
	return config

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
