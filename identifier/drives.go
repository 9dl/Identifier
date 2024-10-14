package identifier

import (
	"github.com/StackExchange/wmi"
	"log"
)

type DiskInfo struct {
	DeviceID    string
	PNPDeviceID string
	Description string
	MediaType   string
}

func GetDiskInfo() (int, int, []DiskInfo, []DiskInfo) {
	var disks []DiskInfo
	query := "SELECT DeviceID, PNPDeviceID, Description, MediaType FROM Win32_DiskDrive"

	err := wmi.Query(query, &disks)
	if err != nil {
		log.Fatalf("Error fetching disk information: %v", err)
	}

	var removableDisks []DiskInfo
	var nonRemovableDisks []DiskInfo

	for _, disk := range disks {
		if disk.MediaType == "Removable Media" {
			removableDisks = append(removableDisks, disk)
		} else {
			nonRemovableDisks = append(nonRemovableDisks, disk)
		}
	}

	return len(removableDisks), len(nonRemovableDisks), removableDisks, nonRemovableDisks
}
