package main

import (
	"math/rand"
	"time"

	"github.com/katallaxie/puzzler/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute()
}
