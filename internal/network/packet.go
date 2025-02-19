package network

import (
	"fmt"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func LogPacket(f *os.File, packet gopacket.Packet) {
    // Time stamp
    timestamp := time.Now().Format(time.RFC3339)

    // Basic packet info
    packetInfo := fmt.Sprintf("\n=== Packet Captured at %s ===\n", timestamp)

    // Ethernet layer
    if ethLayer := packet.Layer(layers.LayerTypeEthernet); ethLayer != nil {
        eth, _ := ethLayer.(*layers.Ethernet)
        packetInfo += fmt.Sprintf("Ethernet: %s -> %s\n", eth.SrcMAC, eth.DstMAC)
    }

    // IPv4 layer
    if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
        ip, _ := ipLayer.(*layers.IPv4)
        packetInfo += fmt.Sprintf("IPv4: %s -> %s\n", ip.SrcIP, ip.DstIP)
    }

    // TCP layer
    if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
        tcp, _ := tcpLayer.(*layers.TCP)
        packetInfo += fmt.Sprintf("TCP: %d -> %d\n", tcp.SrcPort, tcp.DstPort)
        packetInfo += fmt.Sprintf("Flags: SYN:%t ACK:%t FIN:%t RST:%t\n",
            tcp.SYN, tcp.ACK, tcp.FIN, tcp.RST)
    }

    // UDP layer
    if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
        udp, _ := udpLayer.(*layers.UDP)
        packetInfo += fmt.Sprintf("UDP: %d -> %d\n", udp.SrcPort, udp.DstPort)
    }

	// Application layer if exists

    // When iterating through packet.Layers() above,
    // if it lists Payload layer then that is the same as
    // this applicationLayer. applicationLayer contains the payload
    applicationLayer := packet.ApplicationLayer()
    if applicationLayer != nil {
		packetInfo += fmt.Sprintf("L7: %s\n", applicationLayer.Payload())
    }

	// Iterate over all layers, printing out each layer type
	packetInfo += fmt.Sprintf("/%s \n","All packet layers:\n")
	for _, layer := range packet.Layers() {
		packetInfo += fmt.Sprintf("%s\n", layer.LayerType())
	}

    if err := packet.ErrorLayer(); err != nil {
        fmt.Println("Error decoding some part of the packet:", err)
    }

    // Write to file
    f.WriteString(packetInfo)
} 