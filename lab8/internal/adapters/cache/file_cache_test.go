package cache_test

import (
    "os"
    "testing"

    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/adapters/cache"
)

func TestFileCache_SetGet(t *testing.T) {
    path := "test_cache.json"
    _ = os.Remove(path)

    c := cache.NewFileCache(path)

    if err := c.Set("key", 12.5); err != nil {
        t.Fatalf("Set error: %v", err)
    }

    v, ok := c.Get("key")
    if !ok {
        t.Fatalf("expected key to exist")
    }
    if v != 12.5 {
        t.Fatalf("expected 12.5, got %v", v)
    }

    _ = os.Remove(path)
}

func TestFileCache_Persist(t *testing.T) {
    path := "test_cache_persist.json"
    _ = os.Remove(path)

    c := cache.NewFileCache(path)

    if err := c.Set("key", 7.3); err != nil {
        t.Fatalf("Set error: %v", err)
    }

    c2 := cache.NewFileCache(path)
    v, ok := c2.Get("key")
    if !ok || v != 7.3 {
        t.Fatalf("expected 7.3, got %v", v)
    }

    _ = os.Remove(path)
}
