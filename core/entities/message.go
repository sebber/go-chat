package entities

import(
  "sync"
)

type Message struct {
  Author  string `json:"author"`
  Text  string `json:"text"`
}

type Messages struct {
  sync.RWMutex
  Store []*Message
}

func (m *Messages) GetAll() []*Message {
  m.RLock()

  messages := make([]*Message, len(m.Store))

  j := 0
  for i := len(m.Store)-1; i >= 0; i-- {
    messages[j] = m.Store[i]
    j++
  }

  m.RUnlock()

  return messages
}

func (m *Messages) Put(message Message) Message {
  m.Lock()

  m.Store = append(m.Store, &message)

  m.Unlock()

  return message
}
