package codec

import (
  "fmt"
  "reflect"
  "testing"
)

const (
  i = "Halokid data"
)

func TestByteCodec_Encode(t *testing.T) {
  bc := ByteCodec{}
  ib, err := bc.Encode([]byte(i))
  t.Log("ib: ", ib, "err: ", err)
}



func TestByteCodec_Decode(t *testing.T) {
  //data := []byte("hello")
  bc := &ByteCodec{}
  //var ix interface{}
  //ix = "a"
  //err := bc.Decode(data, "a")
  //t.Log("data:", data, "err:", err)

  //reflect.ValueOf("a")

  //data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
  //data := []byte{}
  //var data []byte
  type Reply struct {
    Greet      string
  }
  type XX struct {
    reply    interface{}
    //name     string
  }
  //xx := &XX{}
  //reflect.Indirect(reflect.ValueOf("a")).SetBytes(data)

  type Msg struct {
    Payload      []byte
  }
  msg := &Msg{}
  //msg.Payload = []byte{99, 88}
  msg.Payload = []byte{123, 34, 71, 114, 101, 101, 116, 34, 58, 34, 228, 189, 160, 229, 165, 189, 32, 74, 105, 109, 109, 121, 34, 125}

  var xx *XX
  xx = new(XX)
  //xx.reply = []byte{1}
  //reply := &Reply{}
  var reply *Reply
  reply = new(Reply)
  reply.Greet = ""
  xx.reply = reply

  //buf := &bytes.Buffer{}
  //err := binary.Read(buf, binary.BigEndian, &reply)
  //xx.reply = buf.Bytes()

  //buf := new(bytes.Buffer)
  //if err := binary.Write(buf, binary.LittleEndian, reply); err != nil {
  //  t.Log("error")
  //}
  //t.Log("buf.Bytes ------------", buf.Bytes())
  //xx.reply = buf.Bytes()

  //var iter interface{}
  //iter = "yyyyyyyyyy"
  //xx.reply = &iter
  //xx.name = "zzzzzzzzz"
  //data := []byte{99}
  //err := bc.Decode(data, xx.reply)
  //y := reflect.ValueOf(xx.reply)
  //t.Log("y ---------", y)
  t.Log("len data ---------", len(msg.Payload))
  t.Logf("msg.payload --------- %+v, %+v", reflect.TypeOf(msg.Payload),  msg.Payload)
  t.Logf("xx.reply --------- %+v, %+v", reflect.TypeOf(xx.reply), xx.reply)
  err := bc.Decode(msg.Payload, xx.reply)

  //err := reflect.ValueOf(msg.Payload)
  t.Log(err)
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