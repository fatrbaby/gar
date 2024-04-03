package shortcut

import (
	"github.com/pkg/errors"
	"net"
)

func LocalIp() (string, error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet
		isIpNet bool
		err     error
	)

	if addrs, err = net.InterfaceAddrs(); err != nil {
		return "", err
	}

	for _, addr = range addrs {
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet {
			if !ipNet.IP.IsLoopback() {
				if ipNet.IP.IsPrivate() {
					if ipNet.IP.To4() != nil {
						return ipNet.IP.String(), nil
					}
				}
			}
		}
	}

	return "", errors.New("can not found local IP")
}
