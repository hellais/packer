package qemu

import (
//	"errors"
//	"fmt"
//	"log"
//	"os"
//	"os/exec"
//	"path/filepath"
//	"strings"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/builder/qemu"
//	commonssh "github.com/mitchellh/packer/common/ssh"
	"github.com/mitchellh/packer/packer"
)

const BuilderId = "raspberry.qemu"

type config struct {
	common.PackerConfig `mapstructure:",squash"`

	Accelerator     string     `mapstructure:"accelerator"`
	BootCommand     []string   `mapstructure:"boot_command"`
	DiskInterface   string     `mapstructure:"disk_interface"`
	DiskSize        uint       `mapstructure:"disk_size"`
	DiskCache       string     `mapstructure:"disk_cache"`
	FloppyFiles     []string   `mapstructure:"floppy_files"`
	Format          string     `mapstructure:"format"`
	Headless        bool       `mapstructure:"headless"`
	DiskImage       bool       `mapstructure:"disk_image"`
	HTTPDir         string     `mapstructure:"http_directory"`
	HTTPPortMin     uint       `mapstructure:"http_port_min"`
	HTTPPortMax     uint       `mapstructure:"http_port_max"`
	ISOChecksum     string     `mapstructure:"iso_checksum"`
	ISOChecksumType string     `mapstructure:"iso_checksum_type"`
	ISOUrls         []string   `mapstructure:"iso_urls"`
	MachineType     string     `mapstructure:"machine_type"`
	NetDevice       string     `mapstructure:"net_device"`
	OutputDir       string     `mapstructure:"output_directory"`
	QemuArgs        [][]string `mapstructure:"qemuargs"`
	QemuBinary      string     `mapstructure:"qemu_binary"`
	ShutdownCommand string     `mapstructure:"shutdown_command"`
	SSHHostPortMin  uint       `mapstructure:"ssh_host_port_min"`
	SSHHostPortMax  uint       `mapstructure:"ssh_host_port_max"`
	SSHPassword     string     `mapstructure:"ssh_password"`
	SSHPort         uint       `mapstructure:"ssh_port"`
	SSHUser         string     `mapstructure:"ssh_username"`
	SSHKeyPath      string     `mapstructure:"ssh_key_path"`
	VNCPortMin      uint       `mapstructure:"vnc_port_min"`
	VNCPortMax      uint       `mapstructure:"vnc_port_max"`
	VMName          string     `mapstructure:"vm_name"`

	// TODO(mitchellh): deprecate
	RunOnce bool `mapstructure:"run_once"`

	RawBootWait        string `mapstructure:"boot_wait"`
	RawSingleISOUrl    string `mapstructure:"iso_url"`
	RawShutdownTimeout string `mapstructure:"shutdown_timeout"`
	RawSSHWaitTimeout  string `mapstructure:"ssh_wait_timeout"`

	bootWait        time.Duration ``
	shutdownTimeout time.Duration ``
	sshWaitTimeout  time.Duration ``
	tpl             *packer.ConfigTemplate
}


type Builder struct {
	config config
	runner multistep.Runner
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
  qemuBuilder := new(qemu.Builder)
  qemuBuilder.config = b.config

  warnings, err := qemuBuilder.Preapre(raws);
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
  qemuBuilder.config = b.config

  artifact, err := qemuBuilder.Run(ui, hook, cache)
  if err != nil {
    return nil, err
  }

	return artifact, nil
}
