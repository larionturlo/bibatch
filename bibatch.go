package bibatch

import (
	"errors"
	"io"
	"sync"
)

type ItemReader struct {
	reader *io.PipeReader
}

type Batch struct {
	items []ItemReader
	size  uint
	mux   sync.Mutex
}

func NewBatch(size uint) *Batch {
	return &Batch{
		items: []ItemReader{},
		size:  size,
		mux:   sync.Mutex{},
	}
}

func (b *Batch) NewWriter() (*io.PipeWriter, error) {
	b.mux.Lock()
	defer b.mux.Unlock()

	if len(b.items) == int(b.size) {
		return nil, errors.New("Batch is already full")
	}

	r, w := io.Pipe()
	b.items = append(b.items, ItemReader{reader: r})
	return w, nil
}

func (b *Batch) Read(data []byte) (int, error) {
	b.mux.Lock()
	defer b.mux.Unlock()

	var n int

	for idx := range b.items {
		k, er := b.items[idx].reader.Read(data[n:])
		if er != nil && er != io.EOF {
			return k, er
		}
		n += k
	}

	return n, nil
}
