#!/usr/bin/env bash
# Runs the test suite of every implementation whose toolchain is installed.
set -u
status=0

run() {  # run <name> <tool> <command...>
  local name=$1 tool=$2; shift 2
  if ! command -v "$tool" >/dev/null 2>&1; then echo "== $name: skipped ($tool not installed)"; return; fi
  echo "== $name"
  bash -c "$*" || status=1
}

run python     python3 'cd python && python3 -m unittest -q'
run go         go      'cd go && go test -race ./...'
run rust       cargo   'cd rust && cargo test -q'
run typescript node    'cd typescript && node --experimental-strip-types --no-warnings --test tokenBucket.test.ts'
run java       javac   'cd java/src && javac -d /tmp/rl_java *.java && java -cp /tmp/rl_java TokenBucketTest'
run cpp        g++     'cd cpp && g++ -std=c++17 -Wall -o /tmp/rl_cpp test.cpp -pthread && /tmp/rl_cpp'
exit $status
