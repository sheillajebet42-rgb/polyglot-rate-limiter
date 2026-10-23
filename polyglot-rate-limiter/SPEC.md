# Token bucket spec (every implementation must satisfy this)

- `TokenBucket(capacity, refill_per_sec, clock)`: starts **full**.
- `allow(cost = 1) -> bool`: first refills `elapsed * refill_per_sec` tokens (never above `capacity`), then consumes `cost` tokens if available and returns `true`; otherwise consumes nothing and returns `false`.
- Time comes from an **injectable clock** (seconds, monotonic), so tests never sleep.
- Not thread-safe unless the language section says otherwise.

## Shared scenario (capacity 3, refill 1/s, clock starts at t=0)
| t | call | result |
|---|---|---|
| 0 | allow x3 | true, true, true |
| 0 | allow | false |
| 2 | allow x2 | true, true |
| 2 | allow | false |
| 100 | allow(cost=5) | false (cost > capacity can never succeed) |
| 100 | allow(cost=3) | true (bucket refilled to capacity, not beyond) |
