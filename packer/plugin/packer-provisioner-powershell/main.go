package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/MSOpenTech/packer-hyperv/packer/provisioner/powershell"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(powershell.Provisioner))
	server.Serve()
}
