package easy

import (
	"bufio"
	"errors"
	_ "fmt"
	"net"
	"os"
	_ "strconv"
	"strings"
)

type IPGener struct {
	NetSegment string
	IPList     DataSlice
}

func NewIPGener(net_seg string) *IPGener {
	ipgn := new(IPGener)
	ipgn.NetSegment = net_seg
	ipgn.ParseSegment()
	return ipgn
}

func (ipg *IPGener) ParseSegment() {
	seg := ipg.NetSegment
	if strings.Contains(seg, "/") {
		ipg.ParseCIDR()
	} else if strings.Contains(seg, "-") {
		err := ipg.ParseRange()
		CheckErr(err)
	}
}

func (ipg *IPGener) ParseCIDR() {
	ip, ipnet, err := net.ParseCIDR(ipg.NetSegment)
	CheckErr(err)
	ip = ip.To4()
	var ips DataSlice
	for ip.Mask(ipnet.Mask); ipnet.Contains(ip); ipg.inc(ip) {
		ips = append(ips, ip.String())
	}
	ipg.IPList = append(ipg.IPList, ips...)
}

func (ipg *IPGener) inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func (ipg *IPGener) ParseRange() error {
	ip_range := strings.Split(ipg.NetSegment, "-")
	if len(ip_range) != 2 {
		errors.New("net segment format is wrong." + ipg.NetSegment)
	}
	for _, ip := range ip_range {
		if net.ParseIP(ip) == nil {
			errors.New("net segment format is wrong." + ipg.NetSegment)
		}
	}
	ipg.computeRangeIp(net.ParseIP(ip_range[0]).To4(), net.ParseIP(ip_range[1]).To4())
	return nil
}

func (ipg *IPGener) getIpLong(ip net.IP) int64 {
	return int64(ip[0])*256*256*256 + int64(ip[1])*256*256 + int64(ip[2])*256 + int64(ip[3])
}

func (ipg *IPGener) ipRangeContains(ip_start net.IP, ip_end net.IP, ip net.IP) bool {
	if ipg.getIpLong(ip) >= ipg.getIpLong(ip_start) && ipg.getIpLong(ip) <= ipg.getIpLong(ip_end) {
		return true
	} else {
		return false
	}
}

func (ipg *IPGener) computeRangeIp(ip_start net.IP, ip_end net.IP) {
	var ips DataSlice
	for ip := ip_start; ipg.ipRangeContains(ip_start, ip_end, ip); ipg.inc(ip) {
		ips = append(ips, ip.String())
	}
	ipg.IPList = append(ipg.IPList, ips...)
}

func (ipg *IPGener) WriteFile() {
	output_f, _ := os.OpenFile("ipList.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer output_f.Close()
	writer := bufio.NewWriter(output_f)
	for _, ip := range ipg.IPList {
		writer.WriteString(ip + "\n")
	}
	writer.Flush()
}
