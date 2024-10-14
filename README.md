# Identifier

## Description

Identifier is a simple command-line tool written in Go that provides essential system information, including:

- Wi-Fi Information
- Motherboard Information
- Drive Information
- BIOS Information
- TPM Status (needs UAC else error)
- Hyper-V Status

## Features

- View details of the currently connected Wi-Fi adapter.
- Get motherboard specifications. (detects Manufacturer)
- List removable and non-removable drives.
- Display the BIOS serial number.
- Check the status of Trusted Platform Module (TPM).
- Determine if Hyper-V is active.

## Usage

1. Clone the repository:

   ```bash
   cd identifier
   ```

2. Ensure you have Go installed. If not, download it from [golang.org](https://golang.org).

3. Build the application:

   ```bash
   go build
   ```

4. Run the application:

   ```bash
   ./identifier
   ```
