package main

import (
	"github.com/hetue/boot"
	"github.com/hetue/gradle/internal"
)

func main() {
	starter := boot.New()
	starter.Build().Boot(internal.New)
}
