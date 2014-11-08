package common

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func scripts_add_vm_to_trusted_hosts_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x2a, 0x48,
		0x2c, 0x4a, 0xcc, 0xd5, 0x88, 0x2e, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x8f,
		0x55, 0xc9, 0x2c, 0xd0, 0xe4, 0xe5, 0x0a, 0x4e, 0x2d, 0xd1, 0xf5, 0x2c,
		0x49, 0xcd, 0x55, 0xd0, 0x0d, 0x48, 0x2c, 0xc9, 0x50, 0x08, 0x0f, 0xf6,
		0x4d, 0xcc, 0xb3, 0x8a, 0x89, 0xc9, 0xc9, 0x4f, 0x4e, 0xcc, 0xc9, 0xc8,
		0x2f, 0x2e, 0x89, 0x89, 0x71, 0xce, 0xc9, 0x4c, 0xcd, 0x03, 0xd2, 0x21,
		0x45, 0xa5, 0xc5, 0x25, 0xa9, 0x29, 0x1e, 0x40, 0xc1, 0x62, 0x05, 0xdd,
		0xb0, 0xc4, 0x9c, 0xd2, 0x54, 0x05, 0xa0, 0x19, 0x0a, 0xba, 0x6e, 0xf9,
		0x45, 0xc9, 0xa9, 0x0a, 0xba, 0xce, 0xf9, 0x79, 0xc9, 0x89, 0x25, 0xa9,
		0x79, 0x40, 0xcc, 0xcb, 0x05, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xd7,
		0xe7, 0xa3, 0x6b, 0x00, 0x00, 0x00,
	},
		"scripts/add_vm_to_trusted_hosts.ps1",
	)
}

func scripts_copy_exported_vm_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7a, 0xbf,
		0x7b, 0x7f, 0x41, 0x62, 0x51, 0x62, 0xae, 0x46, 0x74, 0x71, 0x49, 0x51,
		0x66, 0x5e, 0x7a, 0xac, 0x4a, 0x71, 0x51, 0x72, 0x40, 0x62, 0x49, 0x86,
		0x8e, 0x02, 0x5c, 0x24, 0xa5, 0xb8, 0x04, 0x4d, 0xa4, 0x2c, 0x23, 0xc5,
		0x25, 0xb3, 0xc8, 0x2f, 0x31, 0x37, 0x15, 0x59, 0x30, 0x17, 0x28, 0xa6,
		0xc9, 0xcb, 0xe5, 0x9c, 0x5f, 0x50, 0xa9, 0xeb, 0x59, 0x92, 0x9a, 0xab,
		0xa0, 0x0b, 0xd2, 0xa6, 0x00, 0x33, 0x31, 0x06, 0x49, 0x9b, 0x82, 0xae,
		0x4b, 0x6a, 0x71, 0x49, 0x66, 0x5e, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x02,
		0xcc, 0x06, 0x05, 0xdd, 0xa2, 0xd4, 0xe4, 0xd2, 0xa2, 0xe2, 0x54, 0x6b,
		0xbc, 0xa6, 0x80, 0xec, 0xc1, 0x6e, 0x00, 0x61, 0x7d, 0x31, 0x5a, 0x7a,
		0x15, 0xb9, 0x39, 0xd8, 0x75, 0x43, 0x95, 0x00, 0x0d, 0x01, 0x04, 0x00,
		0x00, 0xff, 0xff, 0x8d, 0x43, 0xf9, 0x79, 0x16, 0x01, 0x00, 0x00,
	},
		"scripts/copy_exported_vm.ps1",
	)
}

func scripts_enable_vm_integration_service_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x2a, 0x48,
		0x2c, 0x4a, 0xcc, 0xd5, 0x88, 0x2e, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x8f,
		0x55, 0x29, 0xcb, 0xf5, 0x4b, 0xcc, 0x4d, 0xd5, 0x51, 0x80, 0x0b, 0x64,
		0xe6, 0x95, 0xa4, 0xa6, 0x17, 0x25, 0x96, 0x64, 0xe6, 0xe7, 0x05, 0xa7,
		0x16, 0x95, 0x65, 0x26, 0xa7, 0x82, 0x14, 0x68, 0xf2, 0x72, 0xb9, 0xe6,
		0x25, 0x26, 0xe5, 0xa4, 0xea, 0x86, 0xf9, 0x7a, 0x62, 0xa8, 0x50, 0x00,
		0x8a, 0x82, 0x54, 0x29, 0x40, 0x8d, 0x53, 0xd0, 0x85, 0xf0, 0xb0, 0x9b,
		0xc5, 0xcb, 0x05, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xf8, 0xd1, 0x24,
		0x84, 0x00, 0x00, 0x00,
	},
		"scripts/enable_vm_integration_service.ps1",
	)
}

func scripts_export_vm_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x2a, 0x48,
		0x2c, 0x4a, 0xcc, 0xd5, 0x88, 0x2e, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x8f,
		0x55, 0x29, 0xcb, 0xf5, 0x4b, 0xcc, 0x4d, 0xd5, 0x51, 0x80, 0x0b, 0x14,
		0x24, 0x96, 0x64, 0x68, 0xf2, 0x72, 0xb9, 0x56, 0x14, 0xe4, 0x17, 0x95,
		0xe8, 0x86, 0xf9, 0x2a, 0xe8, 0x82, 0x14, 0x28, 0x40, 0x15, 0x2a, 0xe8,
		0x06, 0x00, 0xe5, 0x15, 0xc0, 0xaa, 0x78, 0xb9, 0x00, 0x01, 0x00, 0x00,
		0xff, 0xff, 0x21, 0x43, 0x62, 0xb1, 0x4c, 0x00, 0x00, 0x00,
	},
		"scripts/export_vm.ps1",
	)
}

func scripts_get_host_free_memory_in_mb_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7a, 0xbf,
		0x7b, 0xbf, 0x86, 0x7b, 0x6a, 0x89, 0x6e, 0x78, 0x6e, 0xa6, 0x7f, 0x52,
		0x56, 0x6a, 0x72, 0x89, 0x42, 0x78, 0x66, 0x9e, 0xb1, 0x51, 0xbc, 0x7f,
		0x41, 0x6a, 0x51, 0x62, 0x49, 0x66, 0x5e, 0x7a, 0x70, 0x65, 0x71, 0x49,
		0x6a, 0xae, 0xa6, 0x9e, 0x5b, 0x51, 0x6a, 0x6a, 0x40, 0x46, 0x65, 0x71,
		0x66, 0x72, 0x62, 0x8e, 0x6f, 0x6a, 0x6e, 0x7e, 0x51, 0xa5, 0x82, 0xbe,
		0x82, 0xa1, 0x81, 0x91, 0x09, 0x20, 0x00, 0x00, 0xff, 0xff, 0xba, 0x9d,
		0xac, 0xa0, 0x42, 0x00, 0x00, 0x00,
	},
		"scripts/get_host_free_memory_in_mb.ps1",
	)
}

func scripts_get_network_adapter_address_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x5c, 0x8e,
		0x31, 0x4b, 0xc5, 0x30, 0x14, 0x46, 0xe7, 0x06, 0xf2, 0x1f, 0xee, 0x90,
		0xa1, 0x1d, 0x2a, 0xee, 0xe2, 0x50, 0x44, 0xc4, 0xc1, 0x22, 0x08, 0x2e,
		0xa5, 0x43, 0x68, 0x6f, 0x35, 0x98, 0x26, 0xf1, 0xe6, 0x56, 0x29, 0xa5,
		0xff, 0xfd, 0x25, 0x8f, 0x76, 0x79, 0x90, 0xe9, 0xe4, 0x7c, 0x87, 0x1b,
		0x34, 0xe9, 0xb9, 0xec, 0x22, 0x93, 0x71, 0x5f, 0xbd, 0xfa, 0x9b, 0x5b,
		0x3d, 0x63, 0x25, 0x05, 0xd3, 0x0a, 0x1b, 0x48, 0x51, 0x28, 0x3d, 0xea,
		0xc0, 0x48, 0xf0, 0x08, 0x2f, 0xc8, 0xf5, 0xe7, 0x5b, 0x8b, 0xfc, 0xef,
		0xe9, 0xa7, 0x39, 0x70, 0x26, 0x69, 0x02, 0xc7, 0x14, 0xea, 0x67, 0x22,
		0x4f, 0xcd, 0xc0, 0xc6, 0x3b, 0xf8, 0x30, 0x16, 0x1d, 0xdb, 0xf5, 0xc9,
		0x3b, 0x36, 0x6e, 0xc1, 0x87, 0x1c, 0x34, 0x21, 0xb5, 0xce, 0xec, 0xdd,
		0xeb, 0x7b, 0x33, 0x8e, 0x84, 0x31, 0x62, 0xec, 0xee, 0xfb, 0x2c, 0x98,
		0xa9, 0xcc, 0x4e, 0x8d, 0xbf, 0xa0, 0xdc, 0x62, 0x6d, 0x05, 0x5b, 0xa2,
		0x05, 0x21, 0x2f, 0xe4, 0x40, 0x4d, 0xda, 0x46, 0x4c, 0x60, 0x97, 0x22,
		0xbd, 0x41, 0xf3, 0xf0, 0x7d, 0x15, 0x6e, 0xfe, 0xf7, 0x74, 0xfc, 0x89,
		0x4c, 0x90, 0xe2, 0x12, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xc3, 0x7f, 0x46,
		0xe9, 0x00, 0x00, 0x00,
	},
		"scripts/get_network_adapter_address.ps1",
	)
}

func scripts_is_current_user_administrator_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7a, 0xbf,
		0x7b, 0xbf, 0x4a, 0x66, 0x4a, 0x6a, 0x5e, 0x49, 0x66, 0x49, 0xa5, 0x6d,
		0x74, 0x70, 0x65, 0x71, 0x49, 0x6a, 0xae, 0x5e, 0x70, 0x6a, 0x72, 0x69,
		0x11, 0x50, 0x40, 0x2f, 0xa0, 0x28, 0x33, 0x2f, 0x39, 0xb3, 0x20, 0x31,
		0x47, 0x2f, 0x3c, 0x33, 0x2f, 0x25, 0xbf, 0xbc, 0xd8, 0x13, 0xaa, 0x34,
		0xd6, 0xca, 0xca, 0x3d, 0xb5, 0xc4, 0xb9, 0xb4, 0xa8, 0x08, 0xc8, 0xd7,
		0xd0, 0xb4, 0xe6, 0xe5, 0x52, 0x29, 0x80, 0xa9, 0xb5, 0xcd, 0x4b, 0x2d,
		0xd7, 0xcd, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0x51, 0x20, 0x68, 0x20, 0x5c,
		0x40, 0x03, 0xee, 0x0c, 0xb0, 0x69, 0x89, 0x29, 0xb9, 0x99, 0x79, 0x99,
		0xc5, 0x25, 0x45, 0x89, 0x25, 0xf9, 0x45, 0x41, 0xf9, 0x39, 0xa9, 0x44,
		0x38, 0xce, 0xa9, 0x34, 0x33, 0xa7, 0xc4, 0x33, 0x0f, 0xa4, 0x1a, 0xe8,
		0x3e, 0x47, 0x64, 0x13, 0x80, 0x46, 0x16, 0xa5, 0x96, 0x94, 0x16, 0xe5,
		0x29, 0x20, 0xdc, 0xa9, 0xe7, 0x59, 0x0c, 0x51, 0xac, 0x81, 0x69, 0x9b,
		0xa6, 0x35, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x08, 0xb9, 0x93, 0x19,
		0x01, 0x00, 0x00,
	},
		"scripts/is_current_user_administrator.ps1",
	)
}

func scripts_mount_dvd_drive_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7a, 0xbf,
		0x7b, 0x7f, 0x41, 0x62, 0x51, 0x62, 0xae, 0x46, 0x74, 0x71, 0x49, 0x51,
		0x66, 0x5e, 0x7a, 0xac, 0x4a, 0x59, 0xae, 0x5f, 0x62, 0x6e, 0xaa, 0x8e,
		0x02, 0x5c, 0xa0, 0x20, 0xb1, 0x24, 0x43, 0x93, 0x97, 0x2b, 0x38, 0xb5,
		0x44, 0x37, 0xcc, 0xd7, 0xa5, 0x2c, 0xc5, 0xa5, 0x28, 0xb3, 0x2c, 0x55,
		0x01, 0xc8, 0x06, 0xa9, 0x53, 0x80, 0xaa, 0x57, 0xd0, 0x0d, 0x00, 0x2a,
		0x53, 0x00, 0x2b, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x8c, 0x33,
		0xca, 0x54, 0x00, 0x00, 0x00,
	},
		"scripts/mount_dvd_drive.ps1",
	)
}

func scripts_mount_floppy_drive_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7a, 0xbf,
		0x7b, 0x7f, 0x41, 0x62, 0x51, 0x62, 0xae, 0x46, 0x74, 0x71, 0x49, 0x51,
		0x66, 0x5e, 0x7a, 0xac, 0x4a, 0x59, 0xae, 0x5f, 0x62, 0x6e, 0xaa, 0x8e,
		0x02, 0x5c, 0xa0, 0x20, 0xb1, 0x24, 0x43, 0x93, 0x97, 0x2b, 0x38, 0xb5,
		0x44, 0x37, 0xcc, 0xd7, 0x2d, 0x27, 0xbf, 0xa0, 0xa0, 0xd2, 0x25, 0xb3,
		0x38, 0xdb, 0xa5, 0x28, 0xb3, 0x2c, 0x55, 0x01, 0x28, 0x04, 0x52, 0xae,
		0x00, 0xd5, 0xa6, 0xa0, 0x1b, 0x00, 0x54, 0xad, 0x00, 0xd6, 0x03, 0x08,
		0x00, 0x00, 0xff, 0xff, 0x4f, 0x0e, 0x15, 0x8d, 0x5b, 0x00, 0x00, 0x00,
	},
		"scripts/mount_floppy_drive.ps1",
	)
}

func scripts_new_vm_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x6c, 0x8f,
		0xbf, 0x0a, 0xc2, 0x30, 0x10, 0x87, 0xf7, 0x42, 0xdf, 0xe1, 0x28, 0x1d,
		0x14, 0x8d, 0x6f, 0xe0, 0xa2, 0x0e, 0x22, 0xb4, 0x08, 0x85, 0x2e, 0xe2,
		0x10, 0x34, 0x98, 0x40, 0xd3, 0x96, 0xf4, 0x6c, 0xad, 0x4f, 0x6f, 0xfe,
		0x98, 0x60, 0xd5, 0x25, 0xe4, 0xee, 0x92, 0xef, 0x77, 0x5f, 0x4b, 0x15,
		0x95, 0xb3, 0x53, 0x87, 0x4a, 0xd4, 0xb7, 0x73, 0xda, 0xcb, 0x9c, 0x4a,
		0xb6, 0x84, 0xd0, 0x68, 0x29, 0x72, 0x5d, 0x56, 0x8d, 0x29, 0x24, 0x93,
		0x8d, 0x1a, 0x0b, 0xa4, 0x0a, 0xef, 0xed, 0x66, 0x44, 0xd6, 0x85, 0x51,
		0xcd, 0x86, 0x72, 0xbf, 0x2b, 0xc4, 0x93, 0xf9, 0xbe, 0x27, 0x74, 0x83,
		0xc0, 0x0b, 0x37, 0xd8, 0x79, 0x1c, 0xc5, 0x51, 0xda, 0xf3, 0xeb, 0x03,
		0xd6, 0xf0, 0x8e, 0x82, 0x05, 0x24, 0x2b, 0xd3, 0x4a, 0xdc, 0xe8, 0xa8,
		0xf3, 0xf4, 0xf4, 0xd0, 0x88, 0x9a, 0xd8, 0xbb, 0x3b, 0xed, 0x1e, 0x40,
		0xb6, 0x5c, 0x54, 0xee, 0x89, 0xc5, 0x18, 0x5e, 0xce, 0x06, 0x52, 0x66,
		0x40, 0x2c, 0xcc, 0x43, 0x27, 0x9f, 0xb2, 0x9f, 0xad, 0xe1, 0x8f, 0x89,
		0x26, 0x58, 0x85, 0x40, 0x77, 0xe9, 0xf9, 0xd4, 0x0b, 0xbe, 0x45, 0x81,
		0x14, 0xc1, 0x0f, 0x3e, 0x5c, 0xe3, 0xe8, 0x15, 0x00, 0x00, 0xff, 0xff,
		0x6f, 0x67, 0x47, 0x55, 0x5a, 0x01, 0x00, 0x00,
	},
		"scripts/new_vm.ps1",
	)
}

func scripts_start_vm_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7a, 0xbf,
		0x7b, 0x7f, 0x41, 0x62, 0x51, 0x62, 0xae, 0x46, 0x74, 0x71, 0x49, 0x51,
		0x66, 0x5e, 0x7a, 0xac, 0x4a, 0x59, 0xae, 0x5f, 0x62, 0x6e, 0xaa, 0x26,
		0x2f, 0x57, 0x70, 0x49, 0x62, 0x51, 0x89, 0x6e, 0x98, 0xaf, 0x82, 0x2e,
		0x48, 0x40, 0x01, 0x2a, 0xc1, 0xcb, 0x05, 0x08, 0x00, 0x00, 0xff, 0xff,
		0xc2, 0x42, 0x6b, 0xb2, 0x33, 0x00, 0x00, 0x00,
	},
		"scripts/start_vm.ps1",
	)
}

func scripts_stop_vm_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7a, 0xbf,
		0x7b, 0x7f, 0x41, 0x62, 0x51, 0x62, 0xae, 0x46, 0x74, 0x71, 0x49, 0x51,
		0x66, 0x5e, 0x7a, 0xac, 0x4a, 0x59, 0xae, 0x5f, 0x62, 0x6e, 0xaa, 0x26,
		0x2f, 0x57, 0x66, 0x9a, 0x82, 0x86, 0x86, 0x7b, 0x6a, 0x89, 0x6e, 0x98,
		0xaf, 0x82, 0x2e, 0x48, 0x4c, 0x01, 0x26, 0xa7, 0x17, 0x5c, 0x92, 0x58,
		0x92, 0xaa, 0xa0, 0x9b, 0x5a, 0xa8, 0x10, 0xed, 0x9b, 0x99, 0x5c, 0x94,
		0x5f, 0x9c, 0x9f, 0x56, 0xa2, 0xe7, 0x51, 0x59, 0x90, 0x5a, 0x14, 0xa6,
		0x17, 0x90, 0x5f, 0x9e, 0x5a, 0x14, 0x9c, 0x91, 0x9a, 0x93, 0xa3, 0x17,
		0xe6, 0x0b, 0x56, 0x18, 0x6b, 0x65, 0x15, 0x54, 0x9a, 0x97, 0x07, 0x34,
		0x5d, 0x53, 0xa1, 0x9a, 0x97, 0x4b, 0x01, 0x08, 0x82, 0x4b, 0xf2, 0x0b,
		0x30, 0xcc, 0xe5, 0xe5, 0xaa, 0xe5, 0xe5, 0x02, 0x04, 0x00, 0x00, 0xff,
		0xff, 0xab, 0x99, 0xe4, 0xea, 0x91, 0x00, 0x00, 0x00,
	},
		"scripts/stop_vm.ps1",
	)
}

func scripts_unmount_dvd_drive_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7a, 0xbf,
		0x7b, 0x7f, 0x41, 0x62, 0x51, 0x62, 0xae, 0x46, 0x74, 0x71, 0x49, 0x51,
		0x66, 0x5e, 0x7a, 0xac, 0x4a, 0x59, 0xae, 0x5f, 0x62, 0x6e, 0xaa, 0x26,
		0x2f, 0x57, 0x70, 0x6a, 0x89, 0x6e, 0x98, 0xaf, 0x4b, 0x59, 0x8a, 0x4b,
		0x51, 0x66, 0x59, 0xaa, 0x02, 0x90, 0x0d, 0x92, 0x50, 0x80, 0x2a, 0x50,
		0xd0, 0x0d, 0x48, 0x2c, 0xc9, 0x50, 0x50, 0xc9, 0x2b, 0xcd, 0xc9, 0xe1,
		0xe5, 0x02, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x41, 0x27, 0x5d, 0x47,
		0x00, 0x00, 0x00,
	},
		"scripts/unmount_dvd_drive.ps1",
	)
}

func scripts_unmount_floppy_drive_ps1() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7a, 0xbf,
		0x7b, 0x7f, 0x41, 0x62, 0x51, 0x62, 0xae, 0x46, 0x74, 0x71, 0x49, 0x51,
		0x66, 0x5e, 0x7a, 0xac, 0x4a, 0x59, 0xae, 0x5f, 0x62, 0x6e, 0xaa, 0x26,
		0x2f, 0x57, 0x70, 0x6a, 0x89, 0x6e, 0x98, 0xaf, 0x5b, 0x4e, 0x7e, 0x41,
		0x41, 0xa5, 0x4b, 0x66, 0x71, 0xb6, 0x4b, 0x51, 0x66, 0x59, 0xaa, 0x02,
		0x50, 0x08, 0x24, 0xaf, 0x00, 0x55, 0xa7, 0xa0, 0x1b, 0x90, 0x58, 0x92,
		0xa1, 0xa0, 0x92, 0x57, 0x9a, 0x93, 0x03, 0x08, 0x00, 0x00, 0xff, 0xff,
		0x21, 0x44, 0x11, 0xce, 0x4c, 0x00, 0x00, 0x00,
	},
		"scripts/unmount_floppy_drive.ps1",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"scripts/add_vm_to_trusted_hosts.ps1": scripts_add_vm_to_trusted_hosts_ps1,
	"scripts/copy_exported_vm.ps1": scripts_copy_exported_vm_ps1,
	"scripts/enable_vm_integration_service.ps1": scripts_enable_vm_integration_service_ps1,
	"scripts/export_vm.ps1": scripts_export_vm_ps1,
	"scripts/get_host_free_memory_in_mb.ps1": scripts_get_host_free_memory_in_mb_ps1,
	"scripts/get_network_adapter_address.ps1": scripts_get_network_adapter_address_ps1,
	"scripts/is_current_user_administrator.ps1": scripts_is_current_user_administrator_ps1,
	"scripts/mount_dvd_drive.ps1": scripts_mount_dvd_drive_ps1,
	"scripts/mount_floppy_drive.ps1": scripts_mount_floppy_drive_ps1,
	"scripts/new_vm.ps1": scripts_new_vm_ps1,
	"scripts/start_vm.ps1": scripts_start_vm_ps1,
	"scripts/stop_vm.ps1": scripts_stop_vm_ps1,
	"scripts/unmount_dvd_drive.ps1": scripts_unmount_dvd_drive_ps1,
	"scripts/unmount_floppy_drive.ps1": scripts_unmount_floppy_drive_ps1,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"scripts/add_vm_to_trusted_hosts.ps1": &_bintree_t{scripts_add_vm_to_trusted_hosts_ps1, map[string]*_bintree_t{
	}},
	"scripts/copy_exported_vm.ps1": &_bintree_t{scripts_copy_exported_vm_ps1, map[string]*_bintree_t{
	}},
	"scripts/enable_vm_integration_service.ps1": &_bintree_t{scripts_enable_vm_integration_service_ps1, map[string]*_bintree_t{
	}},
	"scripts/export_vm.ps1": &_bintree_t{scripts_export_vm_ps1, map[string]*_bintree_t{
	}},
	"scripts/get_host_free_memory_in_mb.ps1": &_bintree_t{scripts_get_host_free_memory_in_mb_ps1, map[string]*_bintree_t{
	}},
	"scripts/get_network_adapter_address.ps1": &_bintree_t{scripts_get_network_adapter_address_ps1, map[string]*_bintree_t{
	}},
	"scripts/is_current_user_administrator.ps1": &_bintree_t{scripts_is_current_user_administrator_ps1, map[string]*_bintree_t{
	}},
	"scripts/mount_dvd_drive.ps1": &_bintree_t{scripts_mount_dvd_drive_ps1, map[string]*_bintree_t{
	}},
	"scripts/mount_floppy_drive.ps1": &_bintree_t{scripts_mount_floppy_drive_ps1, map[string]*_bintree_t{
	}},
	"scripts/new_vm.ps1": &_bintree_t{scripts_new_vm_ps1, map[string]*_bintree_t{
	}},
	"scripts/start_vm.ps1": &_bintree_t{scripts_start_vm_ps1, map[string]*_bintree_t{
	}},
	"scripts/stop_vm.ps1": &_bintree_t{scripts_stop_vm_ps1, map[string]*_bintree_t{
	}},
	"scripts/unmount_dvd_drive.ps1": &_bintree_t{scripts_unmount_dvd_drive_ps1, map[string]*_bintree_t{
	}},
	"scripts/unmount_floppy_drive.ps1": &_bintree_t{scripts_unmount_floppy_drive_ps1, map[string]*_bintree_t{
	}},
}}
