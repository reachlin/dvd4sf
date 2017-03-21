package main

import (
	"fmt"

  "github.com/docker/go-plugins-helpers/volume"
)

func main() {
  d := NewDriver("root")
  h := volume.NewHandler(d)
  h.ServeUnix("root", "test_volume")
	fmt.Println("works")
}
