package main

import(
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"

	"fmt"
);

func main() {

	//TODO: interface selection will be added
	handle, err := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever);

	if(err != nil) {
		panic(err)
	}

	//TODO: Make dynamic this filter
	err = handle.SetBPFFilter("tcp");
	if err != nil {
		panic(err)
	}

	fmt.Println("tcp listening on eth0 interface...");

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		//fmt.Printf("%T \n", packet)
		fmt.Println(packet);
	}	  
}

