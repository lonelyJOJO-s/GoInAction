单个Goruntines中，happen-before顺序即为程序定义的顺序

初始化：main.init < main.main
Goroutine 创建: go < Goroutine 开始执行
Goroutine 销毁: Goroutine 退出 = ∀ e
channel: 如果 ch 是一个 buffered channel，则 ch<-val < val <- ch
channel: 如果 ch 是一个 buffered channel 则 close(ch) < val <- ch & val == isZero(val)
channel: 如果 ch 是一个 unbuffered channel 则，ch<-val > val <- ch
channel: 如果 ch 是一个容量 len(ch) == C 的 buffered channel，则 从 channel 中收到第 k 个值 < k+C 个值得发送完成
mutex: 如果对于 sync.Mutex/sync.RWMutex 的锁 l 有 n < m, 则第 n 次调用 l.Unlock() < 第 m 次调用 l.Lock() 的返回
mutex: 任何发生在 sync.RWMutex 上的调用 l.RLock, 存在一个 n 使得 l.RLock > 第 n 次调用 l.Unlock，且与之匹配的 l.RUnlock < 第 n+1 次调用 l.Lock
once: f() 在 once.Do(f) 中的调用 < once.Do(f) 的返回