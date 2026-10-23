import unittest

from token_bucket import TokenBucket


class Clock:
    def __init__(self):
        self.t = 0.0

    def __call__(self):
        return self.t


class TestTokenBucket(unittest.TestCase):
    def test_shared_scenario(self):
        c = Clock()
        b = TokenBucket(3, 1, c)
        self.assertEqual([b.allow() for _ in range(3)], [True] * 3)
        self.assertFalse(b.allow())
        c.t = 2
        self.assertEqual([b.allow(), b.allow(), b.allow()], [True, True, False])
        c.t = 100
        self.assertFalse(b.allow(5))
        self.assertTrue(b.allow(3))

    def test_invalid_args(self):
        with self.assertRaises(ValueError):
            TokenBucket(0, 1)


if __name__ == "__main__":
    unittest.main()
