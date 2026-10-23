import threading
import time
from typing import Callable


class TokenBucket:
    def __init__(self, capacity: float, refill_per_sec: float, clock: Callable[[], float] = time.monotonic):
        if capacity <= 0 or refill_per_sec < 0:
            raise ValueError("capacity must be > 0 and refill_per_sec >= 0")
        self.capacity = capacity
        self.refill_per_sec = refill_per_sec
        self._clock = clock
        self._tokens = float(capacity)
        self._last = clock()
        self._lock = threading.Lock()  # the Python version is thread-safe

    def allow(self, cost: float = 1) -> bool:
        with self._lock:
            now = self._clock()
            self._tokens = min(self.capacity, self._tokens + (now - self._last) * self.refill_per_sec)
            self._last = now
            if cost <= self._tokens:
                self._tokens -= cost
                return True
            return False
