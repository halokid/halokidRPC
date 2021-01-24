package codec

import (
  "encoding/json"
  "fmt"
  "reflect"
)

type Codec interface {
  Encode(i interface{}) ([]byte, error)
  Decode(data []byte, i interface{}) error
}

type ByteCodec struct {}

func (c ByteCodec) Encode(i interface{}) ([]byte, error) {
  if data, ok := i.([]byte); ok {
    return data, nil
  }
  if data, ok := i.(*[]byte); ok {
    return *data, nil
  }
  return nil, fmt.Errorf("%T 不是[]byte类型", i)
}

func (c ByteCodec) Decode(data []byte, i interface{}) error {
 // reflect.Indirect取得 interface{} 实际的类型和值
 reflect.Indirect(reflect.ValueOf(i)).SetBytes(data)
 return nil
}

// -----------------------------------------

// JSONCodec uses json marshaler and unmarshaler.
type JSONCodec struct{}

// Encode encodes an object into slice of bytes.
func (c JSONCodec) Encode(i interface{}) ([]byte, error) {
  return json.Marshal(i)
}

// Decode decodes an object from slice of bytes.
func (c JSONCodec) Decode(data []byte, i interface{}) error {
  return json.Unmarshal(data, i)
}



