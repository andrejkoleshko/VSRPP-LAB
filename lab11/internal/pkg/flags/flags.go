package flags

import "flag"

type Flags struct {
    Path string
}

func Parse() *Flags {
    config := flag.String("config", "", "path to config")
    flag.Parse()

    return &Flags{
        Path: *config,
    }
}
