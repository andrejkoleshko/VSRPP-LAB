package cache_test

import (
    "testing"
    "time"

    "github.com/alicebob/miniredis/v2"

    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/adapters/cache"
)

func newTestRedisCache(t *testing.T) (*cache.RedisCache, *miniredis.Miniredis) {
    t.Helper()

    s, err := miniredis.Run()
    if err != nil {
        t.Fatalf("failed to start miniredis: %v", err)
    }

    c := cache.NewRedisCache(s.Addr())
    return c, s
}

func TestRedisCache_SetGet(t *testing.T) {
    c, s := newTestRedisCache(t)
    defer s.Close()

    if err := c.Set("key", 5.5); err != nil {
        t.Fatalf("Set error: %v", err)
    }

    v, ok := c.Get("key")
    if !ok {
        t.Fatalf("expected key to exist")
    }
    if v != 5.5 {
        t.Fatalf("expected 5.5, got %v", v)
    }
}

func TestRedisCache_Get_Miss(t *testing.T) {
    c, s := newTestRedisCache(t)
    defer s.Close()

    v, ok := c.Get("missing")
    if ok {
        t.Fatalf("expected key to be missing, got %v", v)
    }
}

func TestRedisCache_Set_TTL(t *testing.T) {
    c, s := newTestRedisCache(t)
    defer s.Close()

    if err := c.Set("key", 1.1); err != nil {
        t.Fatalf("Set error: %v", err)
    }

    s.FastForward(11 * time.Minute)

    _, ok := c.Get("key")
    if ok {
        t.Fatalf("expected key to expire")
    }
}
