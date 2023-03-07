package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

/*
使用有缓冲的通道实现管理任意数量的goroutine之间的资源共享
*/

// pool管理多个goroutine共享的内存资源
// 被管理资源应实现io.Closer 接口
type pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("pool has been closed")

func New(fn func() (io.Closer, error), size uint) (*pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}

	return &pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resources")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resources")
		return p.factory()
	}
}

func (p *pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}
	p.closed = true

	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}

func (p *pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Panicln("Release:", "Closing")
		r.Close()
	}
}
