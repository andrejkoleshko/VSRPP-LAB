package cache

import (
    "encoding/json"
    "os"
)

type FileCache struct {
    path string
    data map[string]float32
}

func NewFileCache(path string) *FileCache {
    fc := &FileCache{
        path: path,
        data: map[string]float32{},
    }

    file, err := os.ReadFile(path)
    if err == nil {
        json.Unmarshal(file, &fc.data)
    }

    return fc
}

func (c *FileCache) Get(key string) (float32, bool) {
    v, ok := c.data[key]
    return v, ok
}

func (c *FileCache) Set(key string, value float32) error {
    c.data[key] = value
    raw, err := json.MarshalIndent(c.data, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(c.path, raw, 0644)
}
