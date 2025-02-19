package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
    targetHost := flag.String("host", "localhost", "Target host")
    interval := flag.Duration("interval", 3*time.Second, "Interval between packets")
    flag.Parse()

    //go generateTCPTraffic(*targetHost, *interval)
    //go generateUDPTraffic(*targetHost, *interval)
	go generateHTTPTraffic(*targetHost, *interval)
	go generateSIPTraffic(*targetHost, *interval)
	go generateDNSTraffic(*targetHost, *interval)
	go generateSNMPTraffic(*targetHost, *interval)
	go generateFTPTraffic(*targetHost, *interval)
	//go generateSSHTraffic(*targetHost, *interval)

    // Keep the program running
    select {}
}

func generateTCPTraffic(host string, interval time.Duration) {
    ports := []int{80, 443, 8080, 22}

    for {
        for _, port := range ports {
            conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
            if err != nil {
                log.Printf("TCP connection failed to %s:%d: %v\n", host, port, err)
                continue
            }
            conn.Write([]byte("TEST TCP PACKET"))
            conn.Close()
        }
        time.Sleep(interval)
    }
}

func generateUDPTraffic(host string, interval time.Duration) {
    ports := []int{53, 123, 161}

    for {
        for _, port := range ports {
            conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", host, port))
            if err != nil {
                log.Printf("UDP connection failed to %s:%d: %v\n", host, port, err)
                continue
            }
            conn.Write([]byte("TEST UDP PACKET"))
            conn.Close()
        }
        time.Sleep(interval)
    }
}

// Generate HTTP traffic
func generateHTTPTraffic(host string, interval time.Duration) {
    for {
        conn, err := net.Dial("tcp", fmt.Sprintf("%s:80", host))
        if err != nil {
            log.Printf("HTTP connection failed to %s: %v\n", host, err)
            continue
        }
        request := "GET / HTTP/1.1\r\nHost: " + host + "\r\n\r\n"
        conn.Write([]byte(request))
        conn.Close()
        time.Sleep(interval)
    }
}

func generateSIPTraffic(host string, interval time.Duration) {
    for {
        conn, err := net.Dial("udp", fmt.Sprintf("%s:5060", host))
        if err != nil {
            log.Printf("SIP connection failed to %s: %v\n", host, err)
            continue
        }
        request := "REGISTER sip:" + host + " SIP/2.0\r\n\r\n"
        conn.Write([]byte(request))
        conn.Close()
        time.Sleep(interval)
    }
}

func generateDNSTraffic(host string, interval time.Duration) {
    for {
        conn, err := net.Dial("udp", fmt.Sprintf("%s:53", host))
        if err != nil {
            log.Printf("DNS connection failed to %s: %v\n", host, err)
            continue
        }
        // Basit bir DNS sorgusu paketi
        dnsQuery := []byte{0x00, 0x01, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
        conn.Write(dnsQuery)
        conn.Close()
        time.Sleep(interval)
    }
}

func generateSNMPTraffic(host string, interval time.Duration) {
    for {
        conn, err := net.Dial("udp", fmt.Sprintf("%s:161", host))
        if err != nil {
            log.Printf("SNMP connection failed to %s: %v\n", host, err)
            continue
        }
        // Basit bir SNMP GET isteği
        snmpRequest := []byte{0x30, 0x26, 0x02, 0x01, 0x01, 0x04, 0x06, 0x70, 0x75, 0x62, 0x6C, 0x69, 0x63}
        conn.Write(snmpRequest)
        conn.Close()
        time.Sleep(interval)
    }
}

func generateFTPTraffic(host string, interval time.Duration) {
    for {
        conn, err := net.Dial("tcp", fmt.Sprintf("%s:21", host))
        if err != nil {
            log.Printf("FTP connection failed to %s: %v\n", host, err)
            continue
        }
        // Basit bir FTP komutu
        ftpCommand := "USER anonymous\r\n"
        conn.Write([]byte(ftpCommand))
        conn.Close()
        time.Sleep(interval)
    }
}

func generateSSHTraffic(host string, interval time.Duration) {
    for {
        conn, err := net.Dial("tcp", fmt.Sprintf("%s:22", host))
        if err != nil {
            log.Printf("SSH connection failed to %s: %v\n", host, err)
            continue
        }
        // SSH bağlantısı için basit bir deneme
        conn.Write([]byte("SSH-2.0-TEST\r\n"))
        conn.Close()
        time.Sleep(interval)
    }
}