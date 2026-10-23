import java.util.function.DoubleSupplier;

public class TokenBucket {
    private final double capacity;
    private final double refillPerSec;
    private final DoubleSupplier clock; // seconds
    private double tokens;
    private double last;

    public TokenBucket(double capacity, double refillPerSec, DoubleSupplier clock) {
        if (capacity <= 0 || refillPerSec < 0) throw new IllegalArgumentException("invalid parameters");
        this.capacity = capacity;
        this.refillPerSec = refillPerSec;
        this.clock = clock;
        this.tokens = capacity;
        this.last = clock.getAsDouble();
    }

    public TokenBucket(double capacity, double refillPerSec) {
        this(capacity, refillPerSec, () -> System.nanoTime() / 1e9);
    }

    public synchronized boolean allow(double cost) {
        double now = clock.getAsDouble();
        tokens = Math.min(capacity, tokens + (now - last) * refillPerSec);
        last = now;
        if (cost <= tokens) {
            tokens -= cost;
            return true;
        }
        return false;
    }

    public boolean allow() { return allow(1); }
}
