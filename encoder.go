package engine

import (
	"sync"
)

type Encoder struct {
	mut    *sync.Mutex
	Locked bool
}

func NewEncoder() *Encoder {
	e := new(Encoder)
	e.mut = new(sync.Mutex)

	return e
}

func (e *Encoder) Lock() {
	e.mut.Lock()
	e.Locked = true
}

func (e *Encoder) Unlock() {
	e.mut.Unlock()
	e.Locked = false
}

func (e *Encoder) Encode(c Chunk) ([]byte, error) {
	return SerializeChunk(c), nil
}

func (e *Encoder) Decode(b []byte) (*Chunk, error) {
	return DeserializeChunk(b), nil
}
