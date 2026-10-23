#include <cstdio>
#include <cstdlib>
#include "token_bucket.hpp"

#define CHECK(c) do { if (!(c)) { std::fprintf(stderr, "FAIL line %d: %s\n", __LINE__, #c); std::exit(1); } } while (0)

int main() {
    double t = 0;
    TokenBucket b(3, 1, [&] { return t; });
    CHECK(b.allow() && b.allow() && b.allow());
    CHECK(!b.allow());
    t = 2;
    CHECK(b.allow() && b.allow() && !b.allow());
    t = 100;
    CHECK(!b.allow(5));
    CHECK(b.allow(3));
    bool threw = false;
    try { TokenBucket bad(0, 1, [] { return 0.0; }); } catch (const std::invalid_argument&) { threw = true; }
    CHECK(threw);
    std::puts("all tests passed");
}
