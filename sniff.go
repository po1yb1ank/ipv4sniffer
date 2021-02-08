package main

import "C"
import (
	"github.com/google/gopacket"
	"golang.org/x/net/ipv4"
	"log"
	"os"

	"github.com/google/gopacket/pcap"

	"fmt"
)
type PacketInfo struct {
	Ipheader *ipv4.Header
	Payload []byte
	Cm	*ipv4.ControlMessage
}

func (p *PacketInfo) String()string{
	return fmt.Sprintf("%#v", p)
}

func Sniff() string{
	f, err := os.Create("log.txt")
	log.SetOutput(f)

	devs, err := pcap.FindAllDevs()
	if err != nil{
		log.Fatal(err)
	}
	log.Println(len(devs))
	/*
	if handle, err := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever); err != nil {
		log.Fatal(err)
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil {  // optional
		log.Fatal(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			output := handlePacket(packet)  // Do something with a packet here.
			return output
		}
	}
	return ""
	 */
	return "s"
}

func handlePacket(packet gopacket.Packet) string {

	return packet.String()
}


/*
func Sniff(dest string) PacketInfo{
	var datagramm []byte
	f, err := os.Create("log.txt")

	c, err := net.ListenPacket("ip:1",dest) // OSPF for IPv4
	if err != nil {
		log.Fatal("in listenPacket",err)
	}
	defer c.Close()
	r, err := ipv4.NewRawConn(c)
	if err != nil {
		log.Fatal("in newRawConn", err)
	}

	ipheader, payload, cm, err := r.ReadFrom(datagramm)
	if err != nil{
		log.Fatal("in ReadFrom",err)
	}

	return PacketInfo{
		Ipheader: ipheader,
		Payload:  payload,
		Cm:       cm,
	}
}
*/
