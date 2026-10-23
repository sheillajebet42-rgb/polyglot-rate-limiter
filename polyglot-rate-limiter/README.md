# Polyglot Rate Limiter
![CI](https://github.com/YOUR_USER/polyglot-rate-limiter/actions/workflows/ci.yml/badge.svg)

The same **token bucket** in Python, Go, Rust, TypeScript, Java and C++, each passing the scenario in [SPEC.md](SPEC.md). Run everything with `./run_all.sh`.

| Language | Time source | Concurrency approach | Notable idiom |
|---|---|---|---|
| Python | injected callable | `threading.Lock` | `min()`-clamped refill |
| Go | injected `func() time.Time` | `sync.Mutex`; race-tested with 500 goroutines | zero-dep, tiny API |
| Rust | caller passes `now` | `&mut self`, so the borrow checker forbids sharing without a lock | no hidden clock, no `unsafe` |
| TypeScript | injected function | none needed (single-threaded) | runs natively with `node --experimental-strip-types` |
| Java | `DoubleSupplier` | `synchronized` | tests run with plain `java`, no build tool |
| C++ | `std::function<double()>` | `std::mutex` | RAII lock guard |

## Why a token bucket
It permits bursts up to `capacity` while enforcing an average of `refill_per_sec`, and it needs only O(1) state per key (two numbers), unlike a sliding-window log.

## Key design point
Every implementation takes its clock as a parameter, so no test ever sleeps and refill math is verified exactly.
