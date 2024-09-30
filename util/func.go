package util

import (
	"net"
	"time"
)

func Contains(l []string, p string) bool {
	for _, v := range l {
		if v == p {
			return true
		}
	}
	return false
}

func UniqueList(s []string) []string {
	keys := make(map[string]struct{})
	res := make([]string, 0)
	for _, val := range s {
		if _, ok := keys[val]; ok {
			continue
		} else {
			keys[val] = struct{}{}
			res = append(res, val)
		}
	}
	return res
}

func GetLocalIP() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addr {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func IsPass(ip string, port string) bool {
	address := net.JoinHostPort(ip, port)
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		return false
	}

	if conn != nil {
		defer conn.Close()
		return true
	} else {
		return false
	}

	return false
}
