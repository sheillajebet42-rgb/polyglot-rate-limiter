export type Clock = () => number; // seconds

export class TokenBucket {
  private readonly capacity: number;
  private readonly refillPerSec: number;
  private readonly clock: Clock;
  private tokens: number;
  private last: number;

  constructor(capacity: number, refillPerSec: number, clock: Clock = () => performance.now() / 1000) {
    if (capacity <= 0 || refillPerSec < 0) throw new RangeError("invalid parameters");
    this.capacity = capacity;
    this.refillPerSec = refillPerSec;
    this.clock = clock;
    this.tokens = capacity;
    this.last = clock();
  }

  // JavaScript is single-threaded, so no locking is needed.
  allow(cost = 1): boolean {
    const now = this.clock();
    this.tokens = Math.min(this.capacity, this.tokens + (now - this.last) * this.refillPerSec);
    this.last = now;
    if (cost <= this.tokens) {
      this.tokens -= cost;
      return true;
    }
    return false;
  }
}
