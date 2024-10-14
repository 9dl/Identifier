package identifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type TPMInfo struct {
	TpmPresent                bool   `json:"TpmPresent"`
	TpmReady                  bool   `json:"TpmReady"`
	TpmEnabled                bool   `json:"TpmEnabled"`
	TpmActivated              bool   `json:"TpmActivated"`
	TpmOwned                  bool   `json:"TpmOwned"`
	RestartPending            bool   `json:"RestartPending"`
	ManufacturerId            int    `json:"ManufacturerId"`
	ManufacturerIdTxt         string `json:"ManufacturerIdTxt"`
	ManufacturerVersion       string `json:"ManufacturerVersion"`
	ManufacturerVersionFull20 string `json:"ManufacturerVersionFull20"`
}

func cleanString(s string) string {
	return strings.Trim(s, "\u0000 \t\n\r")
}

func GetTPMStatus() (TPMInfo, error) {
	cmd := exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", "-Command", "Get-Tpm | ConvertTo-Json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return TPMInfo{}, fmt.Errorf("error executing PowerShell command: %v", err)
	}

	var tpmInfo TPMInfo

	err = json.Unmarshal(out.Bytes(), &tpmInfo)
	if err != nil {
		return TPMInfo{}, fmt.Errorf("error unmarshaling TPM information: %v", err)
	}

	tpmInfo.ManufacturerVersion = cleanString(tpmInfo.ManufacturerVersion)
	tpmInfo.ManufacturerVersionFull20 = cleanString(tpmInfo.ManufacturerVersionFull20)

	return tpmInfo, nil
}
