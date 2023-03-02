package phx

import (
	"encoding/json"
)

// A Serializer describes the required interface for serializers
type Serializer interface {
	Vsn() string
	Encode(*Message) ([]byte, error)
	Decode([]byte) (*Message, error)
}

// JSONSerializerV1 implements the original JSON protocol, which is a JSON object with keys and values

type JSONSerializerV1 struct{}

func NewJSONSerializerV1() *JSONSerializerV1 {
	return &JSONSerializerV1{}
}

func (s *JSONSerializerV1) Vsn() string {
	return "1.0.0"
}

func (s *JSONSerializerV1) Encode(msg *Message) ([]byte, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("encode: %+v -> %s\n", msg, data)
	return data, nil
}

func (s *JSONSerializerV1) Decode(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("decode: %s -> %+v\n", data, msg)
	return &msg, nil
}

//// JSONSerializerV2 implements the V2 protocol, which is basically `[joinRef, ref, topic, event, payload]`.

type JSONSerializerV2 struct{}

func NewJSONSerializerV2() *JSONSerializerV2 {
	return &JSONSerializerV2{}
}

func (s *JSONSerializerV2) Vsn() string {
	return "2.0.0"
}

func (s *JSONSerializerV2) Encode(msg *Message) ([]byte, error) {
	jm := NewJSONMessage(*msg)
	tmp := []any{&jm.JoinRef, &jm.Ref, &jm.Topic, &jm.Event, &jm.Payload}
	data, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("encode: %+v -> %s\n", msg, data)
	return data, nil
}

func (s *JSONSerializerV2) Decode(data []byte) (*Message, error) {
	var jm JSONMessage
	tmp := []any{&jm.JoinRef, &jm.Ref, &jm.Topic, &jm.Event, &jm.Payload}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("decode: %s -> %#v\n", data, tmp)
	msg, err := jm.Message()
	if err != nil {
		return nil, err
	}

	//fmt.Printf("decode: %s -> %+v\n", data, msg)
	return msg, nil
}
