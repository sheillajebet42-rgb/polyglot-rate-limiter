#pragma once
#include <algorithm>
#include <functional>
#include <mutex>
#include <stdexcept>

class TokenBucket {
public:
    using Clock = std::function<double()>;  // seconds

    TokenBucket(double capacity, double refill_per_sec, Clock clock)
        : capacity_(capacity), rate_(refill_per_sec), clock_(std::move(clock)), tokens_(capacity) {
        if (capacity <= 0 || refill_per_sec < 0) throw std::invalid_argument("invalid parameters");
        last_ = clock_();
    }

    bool allow(double cost = 1) {
        std::lock_guard<std::mutex> lock(mu_);
        const double now = clock_();
        tokens_ = std::min(capacity_, tokens_ + (now - last_) * rate_);
        last_ = now;
        if (cost <= tokens_) { tokens_ -= cost; return true; }
        return false;
    }

private:
    double capacity_, rate_;
    Clock clock_;
    double tokens_, last_ = 0;
    std::mutex mu_;
};
