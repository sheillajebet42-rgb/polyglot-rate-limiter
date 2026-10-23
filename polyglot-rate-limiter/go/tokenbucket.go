package ratelimiter

import (
	"sync"
	"time"
)

// TokenBucket is safe for concurrent use.
type TokenBucket struct {
	mu       sync.Mutex
	capacity float64
	rate     float64
	tokens   float64
	last     time.Time
	now      func() time.Time
}

func NewTokenBucket(capacity, refillPerSec float64, now func() time.Time) *TokenBucket {
	if now == nil {
		now = time.Now
	}
	return &TokenBucket{capacity: capacity, rate: refillPerSec, tokens: capacity, last: now(), now: now}
}

func (b *TokenBucket) Allow(cost float64) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	t := b.now()
	b.tokens += t.Sub(b.last).Seconds() * b.rate
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}
	b.last = t
	if cost <= b.tokens {
		b.tokens -= cost
		return true
	}
	return false
}
