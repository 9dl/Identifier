package main

import (
	"Identifier/identifier"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func clearScreen() {
	var cmd *exec.Cmd
	if os.PathSeparator == '\\' { // Windows
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		clearScreen()
		fmt.Println("\nSystem Information Menu:")
		fmt.Println("1. Wi-Fi Information")
		fmt.Println("2. Motherboard Information")
		fmt.Println("3. Drive Information")
		fmt.Println("4. BIOS Information")
		fmt.Println("5. TPM Status")
		fmt.Println("6. Hyper-V Status")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			log.Printf("Invalid input: %v", err)
			continue
		}

		switch choice {
		case 1:
			wifi()
			wait()
		case 2:
			motherboard()
			wait()
		case 3:
			drives()
			wait()
		case 4:
			bios()
			wait()
		case 5:
			tpm()
			wait()
		case 6:
			hyperv()
			wait()
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}

		fmt.Scanln()
	}
}

// Function to wait for user input
func wait() {
	fmt.Println()
	fmt.Print("Press Enter to continue...")
	fmt.Scanln() // Wait for the user to press Enter
}

func wifi() {
	var showAll bool
	fmt.Print("Show all MAC addresses? (true/false): ")
	if _, err := fmt.Scan(&showAll); err != nil {
		log.Printf("Invalid input: %v", err)
		return
	}

	adapters, err := identifier.GetNetworkAdapters()
	if err != nil {
		log.Printf("Failed to get network adapters: %v", err)
		return
	}

	activeWifiAdapter := identifier.GetActiveWifiAdapter(adapters)
	if activeWifiAdapter != nil {
		fmt.Println("Currently connected to Wi-Fi:")
		fmt.Printf("Adapter Name: %s\nManufacturer: %s\nProduct Name: %s\nMAC Address: %s\nType: %s\nStatus: %s\n",
			activeWifiAdapter.Name, activeWifiAdapter.Manufacturer, activeWifiAdapter.ProductName, activeWifiAdapter.MACAddress, "Wi-Fi", identifier.GetConnectionStatus(activeWifiAdapter.NetConnectionStatus))

		if showAll {
			fmt.Println("All MAC Addresses:")
			for _, a := range adapters {
				if a.MACAddress != "" {
					fmt.Printf("MAC Address: %s - %s\n", a.MACAddress, a.Name)
				}
			}
		}
	} else {
		fmt.Println("No active Wi-Fi connection found.")
	}
}

func motherboard() {
	info, err := identifier.GetMotherboardInfo()
	if err != nil {
		log.Printf("Failed to get motherboard info: %v", err)
		return
	}
	for _, board := range info {
		originalManufacturer := board.Manufacturer
		detectedManufacturer := identifier.DetectManufacturer(originalManufacturer)
		fmt.Printf("Original Manufacturer: %s\nDetected Manufacturer: %s\nProduct: %s\nSerial Number: %s\nVersion: %s\n",
			originalManufacturer, detectedManufacturer, board.Product, board.SerialNumber, board.Version)
	}
}

func drives() {
	removableCount, nonRemovableCount, removableDisks, nonRemovableDisks := identifier.GetDiskInfo()

	fmt.Printf("Removable Disks Count: %d\n", removableCount)
	for _, disk := range removableDisks {
		fmt.Printf("Removable Disk - Device ID: %s, PNP Device ID: %s, Description: %s, Media Type: %s\n",
			disk.DeviceID, disk.PNPDeviceID, disk.Description, disk.MediaType)
	}

	fmt.Printf("Non-Removable Disks Count: %d\n", nonRemovableCount)
	for _, disk := range nonRemovableDisks {
		fmt.Printf("Non-Removable Disk - Device ID: %s, PNP Device ID: %s, Description: %s, Media Type: %s\n",
			disk.DeviceID, disk.PNPDeviceID, disk.Description, disk.MediaType)
	}
}

func bios() {
	serialNumber, err := identifier.GetBIOSInfo()
	if err != nil {
		log.Printf("Failed to get BIOS serial number: %v", err)
		return
	}
	fmt.Printf("BIOS Serial Number: %s\n", serialNumber)
}

func tpm() {
	tpmInfo, err := identifier.GetTPMStatus()
	if err != nil {
		log.Printf("Failed to get TPM status: %v", err)
		return
	}

	if tpmInfo.TpmPresent {
		fmt.Println("TPM is present.")
		if tpmInfo.TpmActivated {
			fmt.Println("TPM is activated.")
		} else {
			fmt.Println("TPM is not activated.")
		}

		if tpmInfo.TpmEnabled {
			fmt.Println("TPM is enabled.")
		} else {
			fmt.Println("TPM is not enabled.")
		}

		if tpmInfo.TpmOwned {
			fmt.Println("TPM is owned.")
		} else {
			fmt.Println("TPM is not owned.")
		}

		fmt.Printf("Manufacturer ID: %d (%s)\n", tpmInfo.ManufacturerId, tpmInfo.ManufacturerIdTxt)
		fmt.Printf("Manufacturer Version: %s\n", tpmInfo.ManufacturerVersion)
		fmt.Printf("Manufacturer Version Full: %s\n", tpmInfo.ManufacturerVersionFull20)
		if tpmInfo.RestartPending {
			fmt.Println("A restart is pending for TPM changes.")
		}
	} else {
		fmt.Println("TPM is not present.")
	}
}

func hyperv() {
	hyperVActive, err := identifier.GetHyperVStatus()
	if err != nil {
		log.Printf("Failed to get Hyper-V status: %v", err)
		return
	}
	if hyperVActive {
		fmt.Println("Hyper-V is active.")
	} else {
		fmt.Println("Hyper-V is not active.")
	}
}
