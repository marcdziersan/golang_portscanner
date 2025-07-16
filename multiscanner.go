package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var commonPorts = []int{22, 80, 443, 3306, 8080, 8443} // beliebig erweiterbar
var timeout = 300 * time.Millisecond
var wg sync.WaitGroup

func main() {
	// 1. Eigene IPs ermitteln
	ownIPs := getLocalIPv4Addresses()

	// 2. Alle erreichbaren IPs im LAN annehmen (192.168.178.x Beispiel!)
	subnet := detectSubnet(ownIPs)
	if subnet == "" {
		fmt.Println("Kein lokales Subnetz gefunden.")
		return
	}
	fmt.Println("Gefundenes Subnetz:", subnet+".0/24")

	// 3. Scannen
	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("%s.%d", subnet, i)
		wg.Add(1)
		go scanHost(ip)
	}

	wg.Wait()
	fmt.Println("Scan abgeschlossen.")
}

// Scannt einen Host auf bekannte Ports
func scanHost(ip string) {
	defer wg.Done()
	open := []int{}
	for _, port := range commonPorts {
		address := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err == nil {
			open = append(open, port)
			conn.Close()
		}
	}
	if len(open) > 0 {
		fmt.Printf("[+] %s: Offen -> %v\n", ip, open)
	}
}

// Holt lokale IPs
func getLocalIPv4Addresses() []string {
	var result []string
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP.String())
			}
		}
	}
	return result
}

// Versucht das Subnetz aus den eigenen IPs abzuleiten (z. B. 192.168.178)
func detectSubnet(ips []string) string {
	for _, ip := range ips {
		if strings.HasPrefix(ip, "192.168.") || strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "172.") {
			parts := strings.Split(ip, ".")
			if len(parts) >= 3 {
				return parts[0] + "." + parts[1] + "." + parts[2]
			}
		}
	}
	return ""
}
