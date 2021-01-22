package protocol

const (
  Header = "HalokidRpc"
)

type Message struct {
  Header      string
  Service     string
  Method      string
  Payload     []byte
}

func NewMessage() *Message {
  return &Message{
    Header: Header,
  }
}

func (m Message) CheckHeader() bool {
  return m.Header == Header
}







