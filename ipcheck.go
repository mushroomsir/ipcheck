package ipcheck

import (
	"net"
	"strings"
)

// IsRange ...
func IsRange(addr string, cidrs ...string) bool {
	originIP := net.ParseIP(addr)
	if originIP == nil {
		return false
	}
	for _, ip := range cidrs {
		i := strings.IndexByte(ip, '/')
		if i < 0 {
			pIP := net.ParseIP(ip)
			if pIP == nil {
				return false
			}
			if pIP.Equal(originIP) {
				return true
			}
		} else {
			_, nets, err := net.ParseCIDR(ip)
			if err != nil {
				return false
			}
			if nets.Contains(originIP) {
				return true
			}
		}
	}
	return false
}
