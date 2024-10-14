package identifier

import (
	"github.com/StackExchange/wmi"
	"strings"
)

type NetworkAdapter struct {
	Name                string
	Manufacturer        string
	NetConnectionID     string
	ProductName         string
	AdapterType         string
	NetConnectionStatus uint16
	MACAddress          string
}

func GetNetworkAdapters() ([]NetworkAdapter, error) {
	var adapters []NetworkAdapter
	query := "SELECT Name, Manufacturer, NetConnectionID, ProductName, AdapterType, NetConnectionStatus, MACAddress FROM Win32_NetworkAdapter"
	err := wmi.Query(query, &adapters)
	return adapters, err
}

func GetActiveWifiAdapter(adapters []NetworkAdapter) *NetworkAdapter {
	for _, adapter := range adapters {
		if adapter.NetConnectionStatus == 2 && isWifiAdapter(adapter.Name) {
			return &adapter
		}
	}
	return nil
}

func isWifiAdapter(name string) bool {
	lowerName := strings.ToLower(name)
	return strings.Contains(lowerName, "wi-fi") || strings.Contains(lowerName, "wireless")
}

func GetConnectionStatus(status uint16) string {
	statusMap := map[uint16]string{
		0: "Disconnected",
		1: "Connecting",
		2: "Connected",
		3: "Disconnecting",
		4: "Hardware Failure",
		5: "Invalid Address",
		6: "Media Disconnected",
	}
	return statusMap[status]
}
