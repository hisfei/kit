package n

import (
	"net"
)

func GetHostIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	ips := "#"
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = ips + ipnet.IP.String() + "/"
			}
		}
	}
	return ips, nil
}
