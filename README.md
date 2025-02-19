# Network Packet Sniffer and Traffic Generator

A network monitoring tool implemented in Go that consists of two main components:
1. A packet sniffer that captures network traffic
2. A traffic generator for testing purposes

## Features

- Packet capture from specified network interfaces
- Support for multiple protocol layers (Ethernet, IPv4, TCP, UDP)
- Detailed packet logging including timestamps and layer information
- Traffic generator supporting multiple protocols:
  - HTTP
  - SIP
  - DNS
  - SNMP
  - FTP
  - TCP/UDP

## Prerequisites

Before running this application, ensure you have:

- Go 1.24 or later
- Docker and Docker Compose (for containerized deployment)
- Required Go packages:
  - github.com/google/gopacket
  - libpcap-dev (for packet capture capabilities)

## Installation & Setup

1. Clone the repository
2. Install dependencies:

```bash
go mod download
```

## Running the Application

### Using Docker Compose

The easiest way to run both the sniffer and traffic generator:

```bash
docker-compose up
```

This will start:
- A packet sniffer container with network capture capabilities
- A traffic generator container that sends test traffic

### Manual Running

To run the sniffer:
```bash
go run cmd/sniffer/main.go -i <interface_name> -o <output_file>
```

To run the traffic generator:
```bash
go run cmd/traffic-generator/main.go -host <target_host> -interval <duration>
```

## Configuration

### Sniffer Options
- `-i`: Network interface to capture (default: "eth0")
- `-o`: Output file for packet logs (default: "packets.log")
- `-promisc`: Enable promiscuous mode (default: true)

### Traffic Generator Options
- `-host`: Target host (default: "localhost")
- `-interval`: Interval between packets (default: 3s)

## Security Note

Please be aware that capturing network packets requires elevated privileges and should be done with caution. Ensure you have appropriate permissions and understand the security implications before running this application in any environment.

## Resources

For more information about the technologies used:

- [Packet Capture with GoPacket](https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket)
- [Golang Logging Best Practices](https://www.honeybadger.io/blog/golang-logging/)
- [Context Package in Go](https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go)
- [Understanding Go Context](https://medium.com/rungo/understanding-the-context-package-b2e407a9cdae)
