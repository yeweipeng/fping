package main

import (
	"time"
	"fmt"
	"flag"
	"os"
	"net"
  "bufio"
	"log"
	"ping"
)

var (
	ip_file string
	ping_interval uint
)

func ping_routine(ip string, ping_interval uint) {
	for {
		ping.Ping(ip)
		time.Sleep(time.Duration(ping_interval) * time.Second)
	}
}

func pingMonitor(ip_list []string, ping_interval uint) {
	for _, ip := range ip_list {
		go ping_routine(ip, ping_interval)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func isIPv4 (ip string) bool {
	ipv4 := net.ParseIP(ip)
	return ipv4.To4() != nil
}

func readIP4File (file_path string) []string {
	ip_set := make(map[string]bool)
	fd, err := os.Open(file_path)
	checkError(err)
  defer fd.Close()

  scanner := bufio.NewScanner(fd)
  for scanner.Scan() {
		ip := scanner.Text()
		if isIPv4(ip) {
			ip_set[ip] = true
		}
  }

	err = scanner.Err()
	checkError(err)

	var ip_list []string
	for ip, _ := range ip_set {
		ip_list = append(ip_list, ip)
	}

	return ip_list
}

func initFlag () {
	flag.StringVar(&ip_file, "f", "", "包含以\\n为间隔符的ip文件")
	flag.UintVar(&ping_interval, "s", 60, "ping的间隔时长秒数")
}

func main () {
	initFlag()
	flag.Parse()
	if ip_file == "" {
		flag.Usage()
		return
	}

	ip_list := readIP4File(ip_file)
	fmt.Println(ip_list)

	pingMonitor(ip_list, ping_interval)
	time.Sleep(time.Second * 1000)
}