package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestSlidingWindowLimiter_Allow(t *testing.T) {
	l := NewSlidingWindowLimiter(10*time.Second, 100)

	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		if !l.Allow() {
			time.Sleep(1 * time.Second)
		}
	}
}
