package limiter

import (
	"sync"
	"time"
)

// NewSlidingWindowLimiter 滑动窗口限频器 eg: NewSlidingWindowLimiter(1*time.Minute, 100) 窗口大小为1分钟，窗口内最大请求次数为100。
func NewSlidingWindowLimiter(windowSize time.Duration, maxOps int) ILimiter {
	return &slidingWindowLimiter{
		windowSize: windowSize,
		maxOps:     maxOps,
	}
}

// slidingWindowLimiter 滑动窗口限频器实现
type slidingWindowLimiter struct {
	windowSize time.Duration
	maxOps     int
	opsTimes   []time.Time
	mu         sync.Mutex
}

// Allow 检查是否允许请求
func (l *slidingWindowLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	threshold := now.Add(-l.windowSize)

	// Remove outdated opsTimes
	i := 0
	for ; i < len(l.opsTimes); i++ {
		if l.opsTimes[i].After(threshold) {
			break
		}
	}
	l.opsTimes = l.opsTimes[i:]

	if len(l.opsTimes) >= l.maxOps {
		return false
	}

	l.opsTimes = append(l.opsTimes, now)
	return true
}
