package service

import (
  "reflect"
  "sync"
)

type typePools struct {
  mu      sync.RWMutex
  pools   map[reflect.Type]*sync.Pool
  New     func(t reflect.Type) interface{}
}

var argsReplyPools = &typePools{
  pools:   make(map[reflect.Type]*sync.Pool),

  New:     func(t reflect.Type) interface{} {
    var argv reflect.Value

    if t.Kind() == reflect.Ptr {
      argv = reflect.New(t.Elem())
    } else {
      argv = reflect.New(t)
    }

    return argv.Interface()     // 返回类型的 interface{} 描述
  },
}

func (p *typePools) Init(t reflect.Type) {
  tp := &sync.Pool{}
  tp.New = func() interface{} {
    return p.New(t)
  }
  p.mu.Lock()
  defer p.mu.Unlock()
  p.pools[t] = tp
}

func (p *typePools) Put(t reflect.Type, x interface{}) {
  p.mu.RLock()
  pool := p.pools[t]
  p.mu.RUnlock()
  pool.Put(x)
}

func (p *typePools) Get(t reflect.Type) interface{} {
  p.mu.RLock()
  pool := p.pools[t]
  p.mu.RUnlock()
  return pool.Get()
}




