package main

import (
	"bytes"
	"net"
	"fmt"
	"time"
	"encoding/binary"
)

type ICMP struct {
	Type uint8
	Code uint8
	Checksum uint16
	Mark uint16
	Seq uint16
}

func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}

func main() {
	const ip_pkg_len = 20
	conn, err := net.Dial("ip:icmp", "github.com")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	var icmp ICMP
	icmp.Type = 8
	icmp.Code = 0
	icmp.Checksum = 0 
	icmp.Mark = 123
	icmp.Seq = 456

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.Checksum = CheckSum(buffer.Bytes())
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)

	_, err = conn.Write(buffer.Bytes())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(buffer.Bytes())

	read_buffer := make([]byte, ip_pkg_len + 8)
	conn.SetReadDeadline(time.Now().Add(time.Millisecond * 1000))
	_, err = conn.Read(read_buffer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(read_buffer[ip_pkg_len:])
}