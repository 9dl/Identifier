package identifier

import (
	"github.com/StackExchange/wmi"
	"strings"
)

type BaseBoard struct {
	Product      string
	Manufacturer string
	SerialNumber string
	Version      string
}

var manufacturerMap = map[string]string{
	"asus":       "ASUS",
	"msi":        "MSI",
	"gigabyte":   "Gigabyte",
	"intel":      "Intel",
	"dell":       "Dell",
	"hp":         "HP",
	"acer":       "Acer",
	"lenovo":     "Lenovo",
	"asrock":     "ASRock",
	"supermicro": "Supermicro",
	"biostar":    "Biostar",
	"evga":       "EVGA",
}

func DetectManufacturer(manufacturer string) string {
	manufacturer = strings.TrimSpace(strings.ToLower(manufacturer))
	for key, normalizedName := range manufacturerMap {
		if strings.Contains(manufacturer, key) {
			return normalizedName
		}
	}
	if manufacturer == "" {
		return "Unknown Manufacturer"
	}
	return strings.Title(manufacturer)
}

func GetMotherboardInfo() ([]BaseBoard, error) {
	var result []BaseBoard
	err := wmi.Query("SELECT Product, Manufacturer, SerialNumber, Version FROM Win32_BaseBoard", &result)
	return result, err
}
