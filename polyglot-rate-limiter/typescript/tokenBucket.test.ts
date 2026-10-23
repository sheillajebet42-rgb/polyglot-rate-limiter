import { test } from "node:test";
import assert from "node:assert/strict";
import { TokenBucket } from "./tokenBucket.ts";

test("shared scenario", () => {
  let t = 0;
  const b = new TokenBucket(3, 1, () => t);
  assert.deepEqual([b.allow(), b.allow(), b.allow()], [true, true, true]);
  assert.equal(b.allow(), false);
  t = 2;
  assert.deepEqual([b.allow(), b.allow(), b.allow()], [true, true, false]);
  t = 100;
  assert.equal(b.allow(5), false);
  assert.equal(b.allow(3), true);
});

test("rejects invalid parameters", () => {
  assert.throws(() => new TokenBucket(0, 1), RangeError);
});
