/** Dependency-free test (run with `java TokenBucketTest`) so the repo needs no build tool. */
public class TokenBucketTest {
    static double t = 0;

    static void check(boolean cond, String msg) {
        if (!cond) throw new AssertionError(msg);
    }

    public static void main(String[] args) {
        TokenBucket b = new TokenBucket(3, 1, () -> t);
        check(b.allow() && b.allow() && b.allow(), "burst of 3");
        check(!b.allow(), "4th denied");
        t = 2;
        check(b.allow() && b.allow() && !b.allow(), "2 tokens after 2s");
        t = 100;
        check(!b.allow(5), "cost > capacity");
        check(b.allow(3), "refill capped at capacity");

        boolean threw = false;
        try { new TokenBucket(0, 1); } catch (IllegalArgumentException e) { threw = true; }
        check(threw, "invalid args");
        System.out.println("all tests passed");
    }
}
