package identifier

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetHyperVStatus() (bool, error) {
	type ComputerSystem struct {
		HypervisorPresent bool
	}
	var cs []ComputerSystem

	query := "SELECT HypervisorPresent FROM Win32_ComputerSystem"
	err := wmi.Query(query, &cs)
	if err != nil {
		return false, fmt.Errorf("error fetching Hyper-V status: %v", err)
	}

	if len(cs) > 0 {
		return cs[0].HypervisorPresent, nil
	}
	return false, fmt.Errorf("no Hyper-V information found")
}
