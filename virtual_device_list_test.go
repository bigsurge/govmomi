/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package govmomi

import (
	"math/rand"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
)

var devices = VirtualDeviceList([]types.BaseVirtualDevice{
	&types.VirtualIDEController{
		VirtualController: types.VirtualController{
			VirtualDevice: types.VirtualDevice{
				DynamicData: types.DynamicData{},
				Key:         200,
				DeviceInfo: &types.Description{
					DynamicData: types.DynamicData{},
					Label:       "IDE 0",
					Summary:     "IDE 0",
				},
				Backing:       nil,
				Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
				SlotInfo:      nil,
				ControllerKey: 0,
				UnitNumber:    0,
			},
			BusNumber: 0,
			Device:    []int{3001, 3000},
		},
	},
	&types.VirtualIDEController{
		VirtualController: types.VirtualController{
			VirtualDevice: types.VirtualDevice{
				DynamicData: types.DynamicData{},
				Key:         201,
				DeviceInfo: &types.Description{
					DynamicData: types.DynamicData{},
					Label:       "IDE 1",
					Summary:     "IDE 1",
				},
				Backing:       nil,
				Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
				SlotInfo:      nil,
				ControllerKey: 0,
				UnitNumber:    0,
			},
			BusNumber: 1,
			Device:    []int{3002},
		},
	},
	&types.VirtualPS2Controller{
		VirtualController: types.VirtualController{
			VirtualDevice: types.VirtualDevice{
				DynamicData: types.DynamicData{},
				Key:         300,
				DeviceInfo: &types.Description{
					DynamicData: types.DynamicData{},
					Label:       "PS2 controller 0",
					Summary:     "PS2 controller 0",
				},
				Backing:       nil,
				Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
				SlotInfo:      nil,
				ControllerKey: 0,
				UnitNumber:    0,
			},
			BusNumber: 0,
			Device:    []int{600, 700},
		},
	},
	&types.VirtualPCIController{
		VirtualController: types.VirtualController{
			VirtualDevice: types.VirtualDevice{
				DynamicData: types.DynamicData{},
				Key:         100,
				DeviceInfo: &types.Description{
					DynamicData: types.DynamicData{},
					Label:       "PCI controller 0",
					Summary:     "PCI controller 0",
				},
				Backing:       nil,
				Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
				SlotInfo:      nil,
				ControllerKey: 0,
				UnitNumber:    0,
			},
			BusNumber: 0,
			Device:    []int{500, 12000, 1000, 4000},
		},
	},
	&types.VirtualSIOController{
		VirtualController: types.VirtualController{
			VirtualDevice: types.VirtualDevice{
				DynamicData: types.DynamicData{},
				Key:         400,
				DeviceInfo: &types.Description{
					DynamicData: types.DynamicData{},
					Label:       "SIO controller 0",
					Summary:     "SIO controller 0",
				},
				Backing:       nil,
				Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
				SlotInfo:      nil,
				ControllerKey: 0,
				UnitNumber:    0,
			},
			BusNumber: 0,
			Device:    []int{9000},
		},
	},
	&types.VirtualKeyboard{
		VirtualDevice: types.VirtualDevice{
			DynamicData: types.DynamicData{},
			Key:         600,
			DeviceInfo: &types.Description{
				DynamicData: types.DynamicData{},
				Label:       "Keyboard ",
				Summary:     "Keyboard",
			},
			Backing:       nil,
			Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
			SlotInfo:      nil,
			ControllerKey: 300,
			UnitNumber:    0,
		},
	},
	&types.VirtualPointingDevice{
		VirtualDevice: types.VirtualDevice{
			DynamicData: types.DynamicData{},
			Key:         700,
			DeviceInfo: &types.Description{
				DynamicData: types.DynamicData{},
				Label:       "Pointing device",
				Summary:     "Pointing device; Device",
			},
			Backing: &types.VirtualPointingDeviceDeviceBackingInfo{
				VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{},
				HostPointingDevice:             "autodetect",
			},
			Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
			SlotInfo:      nil,
			ControllerKey: 300,
			UnitNumber:    1,
		},
	},
	&types.VirtualMachineVideoCard{
		VirtualDevice: types.VirtualDevice{
			DynamicData: types.DynamicData{},
			Key:         500,
			DeviceInfo: &types.Description{
				DynamicData: types.DynamicData{},
				Label:       "Video card ",
				Summary:     "Video card",
			},
			Backing:       nil,
			Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
			SlotInfo:      nil,
			ControllerKey: 100,
			UnitNumber:    0,
		},
		VideoRamSizeInKB: 4096,
		NumDisplays:      1,
		UseAutoDetect:    false,
		Enable3DSupport:  false,
		Use3dRenderer:    "automatic",
	},
	&types.VirtualMachineVMCIDevice{
		VirtualDevice: types.VirtualDevice{
			DynamicData: types.DynamicData{},
			Key:         12000,
			DeviceInfo: &types.Description{
				DynamicData: types.DynamicData{},
				Label:       "VMCI device",
				Summary:     "Device on the virtual machine PCI bus that provides support for the virtual machine communication interface",
			},
			Backing:     nil,
			Connectable: (*types.VirtualDeviceConnectInfo)(nil),
			SlotInfo: &types.VirtualDevicePciBusSlotInfo{
				VirtualDeviceBusSlotInfo: types.VirtualDeviceBusSlotInfo{},
				PciSlotNumber:            33,
			},
			ControllerKey: 100,
			UnitNumber:    17,
		},
		Id: 1754519335,
		AllowUnrestrictedCommunication: false,
	},
	&types.VirtualLsiLogicController{
		VirtualSCSIController: types.VirtualSCSIController{
			VirtualController: types.VirtualController{
				VirtualDevice: types.VirtualDevice{
					DynamicData: types.DynamicData{},
					Key:         1000,
					DeviceInfo: &types.Description{
						DynamicData: types.DynamicData{},
						Label:       "SCSI controller 0",
						Summary:     "LSI Logic",
					},
					Backing:       nil,
					Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
					SlotInfo:      nil,
					ControllerKey: 100,
					UnitNumber:    3,
				},
				BusNumber: 0,
				Device:    nil,
			},
			HotAddRemove:       true,
			SharedBus:          "noSharing",
			ScsiCtlrUnitNumber: 7,
		},
	},
	&types.VirtualCdrom{
		VirtualDevice: types.VirtualDevice{
			DynamicData: types.DynamicData{},
			Key:         3001,
			DeviceInfo: &types.Description{
				DynamicData: types.DynamicData{},
				Label:       "CD/DVD drive 1",
				Summary:     "ATAPI cdrom-200-1",
			},
			Backing: &types.VirtualCdromAtapiBackingInfo{
				VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{
					VirtualDeviceBackingInfo: types.VirtualDeviceBackingInfo{},
					DeviceName:               "cdrom-200-1",
					UseAutoDetect:            false,
				},
			},
			Connectable: &types.VirtualDeviceConnectInfo{
				DynamicData:       types.DynamicData{},
				StartConnected:    true,
				AllowGuestControl: true,
				Connected:         false,
				Status:            "untried",
			},
			SlotInfo:      nil,
			ControllerKey: 200,
			UnitNumber:    1,
		},
	},
	&types.VirtualDisk{
		VirtualDevice: types.VirtualDevice{
			DynamicData: types.DynamicData{},
			Key:         3000,
			DeviceInfo: &types.Description{
				DynamicData: types.DynamicData{},
				Label:       "Hard disk 1",
				Summary:     "30,720 KB",
			},
			Backing: &types.VirtualDiskFlatVer2BackingInfo{
				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
					VirtualDeviceBackingInfo: types.VirtualDeviceBackingInfo{},
					FileName:                 "[datastore1] bar/bar.vmdk",
					Datastore:                &types.ManagedObjectReference{Type: "Datastore", Value: "53fe43cc-75dc5110-3643-000c2918dc41"},
					BackingObjectId:          "3-3000-0",
				},
				DiskMode:        "persistent",
				Split:           false,
				WriteThrough:    false,
				ThinProvisioned: false,
				EagerlyScrub:    true,
				Uuid:            "6000C296-d0af-1209-1975-10c98eae10e4",
				ContentId:       "d46395062e2d1b1790985bdec573b211",
				ChangeId:        "",
				Parent: &types.VirtualDiskFlatVer2BackingInfo{
					VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
						VirtualDeviceBackingInfo: types.VirtualDeviceBackingInfo{},
						FileName:                 "[datastore1] ttylinux.vmdk",
						Datastore:                &types.ManagedObjectReference{Type: "Datastore", Value: "53fe43cc-75dc5110-3643-000c2918dc41"},
						BackingObjectId:          "3-3000-1",
					},
					DiskMode:        "persistent",
					Split:           false,
					WriteThrough:    false,
					ThinProvisioned: false,
					EagerlyScrub:    true,
					Uuid:            "6000C296-d0af-1209-1975-10c98eae10e4",
					ContentId:       "1c2dad9e1662219e962a620c6d238a7c",
					ChangeId:        "",
					Parent:          (*types.VirtualDiskFlatVer2BackingInfo)(nil),
					DeltaDiskFormat: "",
					DigestEnabled:   false,
					DeltaGrainSize:  0,
				},
				DeltaDiskFormat: "redoLogFormat",
				DigestEnabled:   false,
				DeltaGrainSize:  0,
			},
			Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
			SlotInfo:      nil,
			ControllerKey: 200,
			UnitNumber:    0,
		},
		CapacityInKB:    30720,
		CapacityInBytes: 31457280,
		Shares: &types.SharesInfo{
			DynamicData: types.DynamicData{},
			Shares:      1000,
			Level:       "normal",
		},
		StorageIOAllocation: &types.StorageIOAllocationInfo{
			DynamicData: types.DynamicData{},
			Limit:       -1,
			Shares: &types.SharesInfo{
				DynamicData: types.DynamicData{},
				Shares:      1000,
				Level:       "normal",
			},
			Reservation: 0,
		},
		DiskObjectId:          "3-3000",
		VFlashCacheConfigInfo: (*types.VirtualDiskVFlashCacheConfigInfo)(nil),
	},
	&types.VirtualDisk{
		VirtualDevice: types.VirtualDevice{
			DynamicData: types.DynamicData{},
			Key:         3002,
			DeviceInfo: &types.Description{
				DynamicData: types.DynamicData{},
				Label:       "Hard disk 2",
				Summary:     "10,000,000 KB",
			},
			Backing: &types.VirtualDiskFlatVer2BackingInfo{
				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
					VirtualDeviceBackingInfo: types.VirtualDeviceBackingInfo{},
					FileName:                 "[datastore1] bar/disk-201-0.vmdk",
					Datastore:                &types.ManagedObjectReference{Type: "Datastore", Value: "53fe43cc-75dc5110-3643-000c2918dc41"},
					BackingObjectId:          "3-3002-0",
				},
				DiskMode:        "persistent",
				Split:           false,
				WriteThrough:    false,
				ThinProvisioned: true,
				EagerlyScrub:    false,
				Uuid:            "6000C293-fde5-4457-5118-dd267ea992a7",
				ContentId:       "90399989b9d520eed6793ab0fffffffe",
				ChangeId:        "",
				Parent:          (*types.VirtualDiskFlatVer2BackingInfo)(nil),
				DeltaDiskFormat: "",
				DigestEnabled:   false,
				DeltaGrainSize:  0,
			},
			Connectable:   (*types.VirtualDeviceConnectInfo)(nil),
			SlotInfo:      nil,
			ControllerKey: 201,
			UnitNumber:    0,
		},
		CapacityInKB:    10000000,
		CapacityInBytes: 10240000000,
		Shares: &types.SharesInfo{
			DynamicData: types.DynamicData{},
			Shares:      1000,
			Level:       "normal",
		},
		StorageIOAllocation: &types.StorageIOAllocationInfo{
			DynamicData: types.DynamicData{},
			Limit:       -1,
			Shares: &types.SharesInfo{
				DynamicData: types.DynamicData{},
				Shares:      1000,
				Level:       "normal",
			},
			Reservation: 0,
		},
		DiskObjectId:          "3-3002",
		VFlashCacheConfigInfo: (*types.VirtualDiskVFlashCacheConfigInfo)(nil),
	},
	&types.VirtualE1000{
		VirtualEthernetCard: types.VirtualEthernetCard{
			VirtualDevice: types.VirtualDevice{
				DynamicData: types.DynamicData{},
				Key:         4000,
				DeviceInfo: &types.Description{
					DynamicData: types.DynamicData{},
					Label:       "Network adapter 1",
					Summary:     "VM Network",
				},
				Backing: &types.VirtualEthernetCardNetworkBackingInfo{
					VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{
						VirtualDeviceBackingInfo: types.VirtualDeviceBackingInfo{},
						DeviceName:               "VM Network",
						UseAutoDetect:            false,
					},
					Network:           &types.ManagedObjectReference{Type: "Network", Value: "HaNetwork-VM Network"},
					InPassthroughMode: false,
				},
				Connectable: &types.VirtualDeviceConnectInfo{
					DynamicData:       types.DynamicData{},
					StartConnected:    true,
					AllowGuestControl: true,
					Connected:         false,
					Status:            "untried",
				},
				SlotInfo: &types.VirtualDevicePciBusSlotInfo{
					VirtualDeviceBusSlotInfo: types.VirtualDeviceBusSlotInfo{},
					PciSlotNumber:            32,
				},
				ControllerKey: 100,
				UnitNumber:    7,
			},
			AddressType:      "generated",
			MacAddress:       "00:0c:29:93:d7:27",
			WakeOnLanEnabled: true,
		},
	},
	&types.VirtualSerialPort{
		VirtualDevice: types.VirtualDevice{
			DynamicData: types.DynamicData{},
			Key:         9000,
			DeviceInfo: &types.Description{
				DynamicData: types.DynamicData{},
				Label:       "Serial port 1",
				Summary:     "Remote localhost:0",
			},
			Backing: &types.VirtualSerialPortURIBackingInfo{
				VirtualDeviceURIBackingInfo: types.VirtualDeviceURIBackingInfo{
					VirtualDeviceBackingInfo: types.VirtualDeviceBackingInfo{},
					ServiceURI:               "localhost:0",
					Direction:                "client",
					ProxyURI:                 "",
				},
			},
			Connectable: &types.VirtualDeviceConnectInfo{
				DynamicData:       types.DynamicData{},
				StartConnected:    true,
				AllowGuestControl: true,
				Connected:         false,
				Status:            "untried",
			},
			SlotInfo:      nil,
			ControllerKey: 400,
			UnitNumber:    0,
		},
		YieldOnPoll: true,
	},
})

func TestSelectByType(t *testing.T) {
	tests := []struct {
		dtype  types.BaseVirtualDevice
		expect int
	}{
		{
			(*types.VirtualCdrom)(nil),
			1,
		},
		{
			(*types.VirtualEthernetCard)(nil),
			1,
		},
		{
			(*types.VirtualDisk)(nil),
			2,
		},
		{
			(*types.VirtualController)(nil),
			6,
		},
		{
			(*types.VirtualIDEController)(nil),
			2,
		},
		{
			(*types.VirtualSCSIController)(nil),
			1,
		},
		{
			(*types.VirtualLsiLogicController)(nil),
			1,
		},
		{
			(*types.ParaVirtualSCSIController)(nil),
			0,
		},
	}

	for _, test := range tests {
		d := devices.SelectByType(test.dtype)

		if len(d) != test.expect {
			t.Errorf("%#v has %d", test.dtype, len(devices))
		}
	}
}

func TestFind(t *testing.T) {
	for _, device := range devices {
		name := devices.Name(device)
		d := devices.Find(name)
		if name != devices.Name(d) {
			t.Errorf("expected name: %s, got: %s", name, devices.Name(d))
		}
	}

	d := devices.Find("enoent")
	if d != nil {
		t.Errorf("unexpected: %#v", d)
	}
}

func TestPickController(t *testing.T) {
	list := devices

	tests := []struct {
		ctype types.BaseVirtualController
		key   int
		unit  int
	}{
		{
			(*types.VirtualIDEController)(nil), 201, 1,
		},
		{
			(*types.VirtualSCSIController)(nil), 1000, 0,
		},
		{
			(*types.VirtualSCSIController)(nil), 1000, 1,
		},
	}

	for _, test := range tests {
		c := list.PickController(test.ctype).GetVirtualController()

		key := c.Key
		if key != test.key {
			t.Errorf("expected controller key: %d, got: %d\n", test.key, key)
		}

		unit := list.newUnitNumber(c)
		if unit != test.unit {
			t.Errorf("expected unit number: %d, got: %d\n", test.unit, unit)
		}

		dev := &types.VirtualDevice{
			Key:           rand.Int(),
			UnitNumber:    unit,
			ControllerKey: key,
		}

		list = append(list, dev)
		c.Device = append(c.Device, dev.Key)
	}

	if list.PickController((*types.VirtualIDEController)(nil)) != nil {
		t.Error("should be nil")
	}

	if list.PickController((*types.VirtualSCSIController)(nil)) == nil {
		t.Errorf("should not be nil")
	}
}

func TestCdrom(t *testing.T) {
	c, err := devices.FindCdrom("")
	if err != nil {
		t.Error(err)
	}

	d := devices.Find(devices.Name(c))

	if c.Key != d.GetVirtualDevice().Key {
		t.Error("device key mismatch")
	}

	for _, name := range []string{"enoent", "ide-200"} {
		c, err = devices.FindCdrom(name)
		if err == nil {
			t.Errorf("FindCdrom(%s) should fail", name)
		}
	}

	c, err = devices.Select(func(device types.BaseVirtualDevice) bool {
		if _, ok := device.(*types.VirtualCdrom); ok {
			return false
		}
		return true
	}).FindCdrom("")

	if err == nil {
		t.Error("FindCdrom('') should fail")
	}
}

func TestName(t *testing.T) {
	tests := []struct {
		device types.BaseVirtualDevice
		expect string
	}{
		{
			&types.VirtualCdrom{},
			"cdrom-0",
		},
		{
			&types.VirtualDisk{},
			"disk-0-0",
		},
		{
			&types.VirtualFloppy{},
			"floppy-0",
		},
		{
			&types.VirtualIDEController{},
			"ide-0",
		},
		{
			&types.VirtualMachineVideoCard{},
			"video-0",
		},
		{
			&types.VirtualPointingDevice{},
			"pointing-0",
		},
		{
			&types.ParaVirtualSCSIController{},
			"pvscsi-0",
		},
		{
			&types.VirtualSerialPort{},
			"serialport-0",
		},
		{
			&types.VirtualE1000{
				VirtualEthernetCard: types.VirtualEthernetCard{
					VirtualDevice: types.VirtualDevice{
						UnitNumber: 7,
					},
				},
			},
			"ethernet-0",
		},
	}

	for _, test := range tests {
		name := devices.Name(test.device)
		if name != test.expect {
			t.Errorf("expected: %s, got: %s", test.expect, name)
		}
	}
}