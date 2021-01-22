package codec

import (
  "encoding/json"
  "fmt"
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

//func (c ByteCodec) Decode(data []byte, i interface{}) error {
//  // reflect.Indirect取得 interface{} 实际的类型和值
//  reflect.Indirect(reflect.ValueOf(i)).SetBytes(data)
//  return nil
//}

//func (c ByteCodec) Decode(data []byte, i interface{}) error {
//  reflect.Indirect(reflect.ValueOf(i)).SetBytes(data)
//  return nil
//}

func (c ByteCodec) Decode(data []byte, i interface{}) error {
  return json.Unmarshal(data, i)
}



