package identifier

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

type BIOSInfo struct {
	SerialNumber string
}

func GetBIOSInfo() (string, error) {
	var bios []BIOSInfo
	query := "SELECT SerialNumber FROM Win32_BIOS"

	err := wmi.Query(query, &bios)
	if err != nil {
		return "", fmt.Errorf("error fetching BIOS information: %v", err)
	}

	if len(bios) > 0 {
		return bios[0].SerialNumber, nil
	}
	return "", fmt.Errorf("no BIOS information found")
}
