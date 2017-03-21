package dvd4sf

import (
  "github.com/docker/go-plugins-helpers/volume"
)

type dvd4sfDriver struct {
  root    string
}

func NewDriver(root string) dvd4sfDriver {
  return dvd4sfDriver {root: root}
}

func (d *dvd4sfDriver) Create(req volume.Request) volume.Response {
  volume.Response{Err: "not implemented"}
}

func (d *dvd4sfDriver) List(req volume.Request) volume.Response {
  volume.Response{Err: "not implemented"}
}

func (d *dvd4sfDriver) Get(req volume.Request) volume.Response {
  volume.Response{Err: "not implemented"}
}

func (d *dvd4sfDriver) Remove(req volume.Request) volume.Response {
  volume.Response{Err: "not implemented"}
}

func (d *dvd4sfDriver) Path(req volume.Request) volume.Response {
  volume.Response{Err: "not implemented"}
}

func (d *dvd4sfDriver) Mount(req volume.MountRequest) volume.Response {
  volume.Response{Err: "not implemented"}
}

func (d *dvd4sfDriver) Unmount(req volume.UnmountRequest) volume.Response {
  volume.Response{Err: "not implemented"}
}

func (d *dvd4sfDriver) Capabilities(req volume.Request) volume.Response {
  volume.Response{Err: "not implemented"}
}
