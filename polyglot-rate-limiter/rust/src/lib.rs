/// Token bucket. Time is passed in explicitly (`now`, in seconds), which makes it
/// trivially testable and avoids hidden global state.
pub struct TokenBucket {
    capacity: f64,
    refill_per_sec: f64,
    tokens: f64,
    last: f64,
}

impl TokenBucket {
    pub fn new(capacity: f64, refill_per_sec: f64, now: f64) -> Self {
        assert!(capacity > 0.0 && refill_per_sec >= 0.0, "invalid parameters");
        Self { capacity, refill_per_sec, tokens: capacity, last: now }
    }

    pub fn allow(&mut self, cost: f64, now: f64) -> bool {
        let elapsed = (now - self.last).max(0.0);
        self.tokens = (self.tokens + elapsed * self.refill_per_sec).min(self.capacity);
        self.last = now;
        if cost <= self.tokens {
            self.tokens -= cost;
            true
        } else {
            false
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn shared_scenario() {
        let mut b = TokenBucket::new(3.0, 1.0, 0.0);
        assert!(b.allow(1.0, 0.0) && b.allow(1.0, 0.0) && b.allow(1.0, 0.0));
        assert!(!b.allow(1.0, 0.0));
        assert!(b.allow(1.0, 2.0) && b.allow(1.0, 2.0));
        assert!(!b.allow(1.0, 2.0));
        assert!(!b.allow(5.0, 100.0));
        assert!(b.allow(3.0, 100.0));
    }

    #[test]
    fn clock_going_backwards_is_ignored() {
        let mut b = TokenBucket::new(1.0, 1.0, 10.0);
        assert!(b.allow(1.0, 10.0));
        assert!(!b.allow(1.0, 5.0));
    }
}
