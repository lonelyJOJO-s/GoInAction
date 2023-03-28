package main

import (
	"sync"
	"time"
)

// 定义订阅者 主题
type subscriber chan interface{}
type topicFunc func(v interface{}) bool

// 发布者模型
type Publisher struct {
	m           sync.Mutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

// 创建一个发布者
func NewPublisher(PublisherTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     PublisherTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// 添加一个新的订阅者，订阅全部主题
func (p *Publisher) Subscriber() chan interface{} {
	return p.SubscriberTopic(nil)
}

// 添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscriberTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup

	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}

	wg.Wait()
}

// 发送主题，可以容忍一定程度的的超时
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// 如果消息被过滤， 不发送信息
	if topic != nil && !topic(v) {
		return
	}

	select {
	// 那个先准备好执行哪一个
	case sub <- v:
	case <-time.After(p.timeout):
	}

}

// 关闭发布者对象
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// func main() {
// 	p := NewPublisher(time.Millisecond*100, 10)
// 	defer p.Close()

// 	all := p.Subscriber()

// 	golang := p.SubscriberTopic(func(v interface{}) bool {
// 		if s, ok := v.(string); ok {
// 			return strings.Contains(s, "golang")
// 		}
// 		return false
// 	})

// 	p.Publish("hello, world")
// 	p.Publish("learn about golang")

// 	go func() {
// 		for msg := range all {
// 			fmt.Println("all:", msg)
// 		}
// 	}()

// 	go func() {
// 		for msg := range golang {
// 			fmt.Println("golang:", msg)
// 		}
// 	}()

// 	time.Sleep(3 * time.Second)

// }
