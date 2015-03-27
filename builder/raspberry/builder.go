package qemu

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/builder/qemu"
	commonssh "github.com/mitchellh/packer/common/ssh"
	"github.com/mitchellh/packer/packer"
)
type Builder struct {
	config qemu.config
	runner multistep.Runner
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
  qemuBuilder := new(qemu.Builder)
  quemuBuilder.config = b.config

  warnings, err := quemuBuilder.Preapre(raws);
	if err != nil {
		return warnings, err
	}

	md, err := common.DecodeConfig(&b.config, raws...)
	if err != nil {
		return nil, err
	}

	b.config.tpl, err = packer.NewConfigTemplate()
	if err != nil {
		return nil, err
	}
	b.config.tpl.UserVars = b.config.PackerUserVars

	// Accumulate any errors
	errs := common.CheckUnusedConfig(md)

  //warning = append(warnings, "TEST");
	return warnings, nil
}

func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
  qemuBuilder := new(qemu.Builder)
  quemuBuilder.config = b.config

  artifact, err := qemuBuilder.Run(ui, hook, cache)
  if err != nil {
    return nil, err
  }

	return artifact, nil
}
