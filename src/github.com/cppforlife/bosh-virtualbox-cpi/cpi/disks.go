package cpi

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	apiv1 "github.com/cppforlife/bosh-cpi-go/apiv1"

	bdisk "github.com/cppforlife/bosh-virtualbox-cpi/disk"
	bvm "github.com/cppforlife/bosh-virtualbox-cpi/vm"
)

type Disks struct {
	creator  bdisk.Creator
	finder   bdisk.Finder
	vmFinder bvm.Finder
}

func NewDisks(creator bdisk.Creator, finder bdisk.Finder, vmFinder bvm.Finder) Disks {
	return Disks{creator, finder, vmFinder}
}

func (a Disks) CreateDisk(size int, _ apiv1.DiskCloudProps, _ *apiv1.VMCID) (apiv1.DiskCID, error) {
	disk, err := a.creator.Create(size)
	if err != nil {
		return apiv1.DiskCID{}, bosherr.WrapErrorf(err, "Creating disk of size '%d'", size)
	}

	return disk.ID(), nil
}

func (a Disks) DeleteDisk(cid apiv1.DiskCID) error {
	disk, err := a.finder.Find(cid)
	if err != nil {
		return bosherr.WrapErrorf(err, "Finding disk '%s'", cid)
	}

	err = disk.Delete()
	if err != nil {
		return bosherr.WrapErrorf(err, "Deleting disk '%s'", cid)
	}

	return nil
}

func (a Disks) AttachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	_, err := a.AttachDiskV2(vmCID, diskCID)
	return err
}

func (a Disks) AttachDiskV2(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) (apiv1.DiskHint, error) {
	vm, err := a.vmFinder.Find(vmCID)
	if err != nil {
		return apiv1.DiskHint{}, bosherr.WrapErrorf(err, "Finding VM '%s'", vmCID)
	}

	disk, err := a.finder.Find(diskCID)
	if err != nil {
		return apiv1.DiskHint{}, bosherr.WrapErrorf(err, "Finding disk '%s'", diskCID)
	}

	hint, err := vm.AttachDisk(disk)
	if err != nil {
		return apiv1.DiskHint{}, bosherr.WrapErrorf(err, "Attaching disk '%s' to VM '%s'", diskCID, vmCID)
	}

	return hint, nil
}

func (a Disks) DetachDisk(vmCID apiv1.VMCID, diskCID apiv1.DiskCID) error {
	vm, err := a.vmFinder.Find(vmCID)
	if err != nil {
		return bosherr.WrapErrorf(err, "Finding VM '%s'", vmCID)
	}

	disk, err := a.finder.Find(diskCID)
	if err != nil {
		return bosherr.WrapErrorf(err, "Finding disk '%s'", diskCID)
	}

	err = vm.DetachDisk(disk)
	if err != nil {
		return bosherr.WrapErrorf(err, "Detaching disk '%s' to VM '%s'", diskCID, vmCID)
	}

	return nil
}

func (a Disks) HasDisk(cid apiv1.DiskCID) (bool, error) {
	disk, err := a.finder.Find(cid)
	if err != nil {
		return false, bosherr.WrapErrorf(err, "Finding disk '%s'", cid)
	}

	return disk.Exists()
}

func (a Disks) SetDiskMetadata(cid apiv1.DiskCID, meta apiv1.DiskMeta) error {
	return nil
}

func (a Disks) ResizeDisk(cid apiv1.DiskCID, size int) error {
	return nil
}
