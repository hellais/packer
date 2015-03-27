package main

import (
	"github.com/hellais/packer/builder/raspberry"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(raspberry.Builder))
	server.Serve()
}
