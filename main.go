package main

import (
	"fmt"

  "github.com/docker/go-plugins-helpers/volume"
)

func main() {
  d := NewDriver("root")
  h := volume.NewHandler(d)
  h.ServeUnix("/run/docker/plugins/dvd4sf.sock", 0)
	fmt.Println("works")
}
