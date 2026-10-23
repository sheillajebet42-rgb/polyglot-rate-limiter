package ratelimiter

import (
	"sync"
	"testing"
	"time"
)

func TestSharedScenario(t *testing.T) {
	now := time.Unix(0, 0)
	b := NewTokenBucket(3, 1, func() time.Time { return now })
	for i := 0; i < 3; i++ {
		if !b.Allow(1) {
			t.Fatalf("call %d should pass", i)
		}
	}
	if b.Allow(1) {
		t.Fatal("4th should fail")
	}
	now = time.Unix(2, 0)
	if !b.Allow(1) || !b.Allow(1) || b.Allow(1) {
		t.Fatal("after 2s expected exactly 2 tokens")
	}
	now = time.Unix(100, 0)
	if b.Allow(5) || !b.Allow(3) {
		t.Fatal("cost > capacity must fail; refill is capped at capacity")
	}
}

func TestConcurrentNeverExceedsCapacity(t *testing.T) {
	now := time.Unix(0, 0)
	b := NewTokenBucket(100, 0, func() time.Time { return now })
	var wg sync.WaitGroup
	var mu sync.Mutex
	ok := 0
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if b.Allow(1) {
				mu.Lock()
				ok++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	if ok != 100 {
		t.Fatalf("allowed %d, want exactly 100", ok)
	}
}
