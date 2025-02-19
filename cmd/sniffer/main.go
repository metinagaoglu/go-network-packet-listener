package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/metinagaoglu/go-network-packet-listener/internal/network"
)

func main() {
	// Command line flags
	interfaceName := flag.String("i", "eth0", "Interface to capture packets from")
	outputFile := flag.String("o", "packets.log", "Output file for packet logs")
	promiscuous := flag.Bool("promisc", true, "Enable promiscuous mode")
	flag.Parse()

	// Open output file
	f, err := os.OpenFile(*outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Open device
	handle, err := pcap.OpenLive(*interfaceName, 1600, *promiscuous, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set up packet source
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Handle SIGINT and SIGTERM
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Printf("Starting packet capture on interface %s\n", *interfaceName)

	// Start packet processing in a goroutine
	go func() {
		for packet := range packetSource.Packets() {
			network.LogPacket(f, packet)
		}
	}()

	// Wait for interrupt signal
	<-signalChan
	fmt.Println("\nShutting down...")
}