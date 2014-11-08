// Copyright (c) Microsoft Open Technologies, Inc.
// All Rights Reserved.
// Licensed under the Apache License, Version 2.0.
// See License.txt in the project root for license information.
package iso

import (
	"fmt"
	"bytes"
	"strconv"
	"github.com/mitchellh/multistep"
	hypervcommon "github.com/MSOpenTech/packer-hyperv/packer/builder/hyperv/common"
	powershell "github.com/MSOpenTech/packer-hyperv/packer/powershell"

	"github.com/mitchellh/packer/packer"
)

// This step creates the actual virtual machine.
//
// Produces:
//   vmName string - The name of the VM
type StepCreateVM struct {
	vmName string
}

func (s *StepCreateVM) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*config)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Creating virtual machine...")

	vmName := config.VMName
//	path :=	config.OutputDir
	path :=	state.Get("packerTempDir").(string)

	powershell, err := powershell.Command()
	ps1, err := hypervcommon.Asset("scripts/new-vm.ps1")
	if err != nil {
		err := fmt.Errorf("Could not load script scripts/new-vm.ps1: %s", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}


	// convert the MB to bytes
	ramBytes := int64(config.RamSizeMB * 1024 * 1024)
	diskSizeBytes := int64(config.DiskSize * 1024 * 1024)

	ram := strconv.FormatInt(ramBytes, 10)
	diskSize := strconv.FormatInt(diskSizeBytes, 10)
	switchName := config.SwitchName

	err = powershell.RunFile(ps1, vmName, path, ram, diskSize, switchName)

	if err != nil {
		err := fmt.Errorf("Error creating virtual machine: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// Set the VM name property on the first command
	if s.vmName == "" {
		s.vmName = vmName
	}

	// Set the final name in the state bag so others can use it
	state.Put("vmName", s.vmName)

	return multistep.ActionContinue
}

func (s *StepCreateVM) Cleanup(state multistep.StateBag) {
	if s.vmName == "" {
		return
	}

	//driver := state.Get("driver").(hypervcommon.Driver)
	ui := state.Get("ui").(packer.Ui)

	powershell, _ := powershell.Command()

	ui.Say("Unregistering and deleting virtual machine...")

	var err error = nil

	var blockBuffer bytes.Buffer
	blockBuffer.WriteString("param([string]$vmName)")
	blockBuffer.WriteString("Remove-VM –Name $vmName -Force }")

	err = powershell.RunFile(blockBuffer.Bytes(), s.vmName)

	if err != nil {
		ui.Error(fmt.Sprintf("Error deleting virtual machine: %s", err))
	}
}
