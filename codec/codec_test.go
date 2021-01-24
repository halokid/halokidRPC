package codec

import (
  "fmt"
  "reflect"
  "testing"
)

const (
)

func TestByteCodec_Encode(t *testing.T) {
  bc := ByteCodec{}
  i := "Halokid data"
  ib, err := bc.Encode([]byte(i))
  t.Log("ib: ", ib, "err: ", err)
}

func TestByteCodec_Decode(t *testing.T) {
  bc := ByteCodec{}
  i := "Halokid data"
  ibEn, err := bc.Encode([]byte(i))
  t.Log("ibEn: ", ibEn, "err: ", err)

  var iw interface{}    // todo: 声明一个类型，只是源码的定义
  iw = new([]byte)      // todo: 实际的数据储存类型，是写进内存位的，runtime实际的类型
  ibDeErr := bc.Decode(ibEn, iw)
  t.Log("ibDeErr:", ibDeErr, " iw:", iw)
}

func TestJsonCodec(t *testing.T) {
  bc := JSONCodec{}
  i := "Halokid data"
  ibEn, err := bc.Encode(i)
  t.Log("ibEn: ", ibEn, "err: ", err)

  var iw *interface{}    // todo: 声明一个类型，只是源码的定义
  iw = new(interface{})   // todo: 实际的数据储存类型，是写进内存位的，runtime实际的类型

  //iw := new(interface{})   // todo: 或者直接new一块内存
  ibDeErr := bc.Decode(ibEn, iw)
  t.Log("ibDeErr:", ibDeErr, " iw:", *iw)
}

func TestComm(t *testing.T) {
  var num float64 = 1.2345
  pointe := reflect.ValueOf(num)
  pointe.Set(reflect.ValueOf(789))
}

func TestComm2(t *testing.T) {
  var num float64 = 1.2345
  pointer := reflect.ValueOf(&num)
  newValue := pointer.Elem()
  fmt.Println("settability of pointer:", newValue.CanSet())
  // 重新赋值
  newValue.SetFloat(77)
  fmt.Println("new value of pointer:", num)
}