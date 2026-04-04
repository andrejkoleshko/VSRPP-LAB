package flags_test

import (
    "os"
    "testing"

    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/pkg/flags"
)

func TestParse_DefaultPath(t *testing.T) {
    os.Args = []string{"cmd"}

    f := flags.Parse()
    if f.Path == "" {
        t.Fatalf("expected non-empty path")
    }
}
