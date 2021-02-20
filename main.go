package main

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v6"
)

type environment struct {
	Home string `env:"HOME"`
}

func main() {
	e := environment{}
	_ = env.Parse(&e)

	dir := flag.String("dir", e.Home, "desired directory")
	flag.Parse()

	fmt.Printf("Directory = %s\n", *dir)
}
