package service

type Echo string

func (e *Echo) Say(name string) []byte {
  return []byte("hello " + name)
}


