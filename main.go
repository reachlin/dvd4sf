package main

import (
	"fmt"

  "./dvd4sf"

  "github.com/docker/go-plugins-helpers/volume"
)

func main() {
  d := dvd4sf.NewDriver("root")
  h := volume.NewHandler(d)
  h.ServeUnix("root", "test_volume")
	fmt.Println("works")
}
