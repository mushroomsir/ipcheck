package ipcheck

import (
	"net"
	"regexp"
	"strings"
)

var validRegEx = regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)

// Bogons list from http://www.team-cymru.org/bogon-reference.html
var bogonsArray = []string{"0.0.0.0/8", "10.0.0.0/8", "100.64.0.0/10", "127.0.0.0/8", "169.254.0.0/16",
	"172.16.0.0/12", "192.0.0.0/24", "192.0.2.0/24", "192.168.0.0/16", "198.18.0.0/15",
	"198.51.100.0/24", "203.0.113.0/24", "224.0.0.0/3"}

// IPinfo ...
type IPinfo struct {
	OriginalIP string
	IsValid    bool
	IsBogon    bool
}

// Check ...
func Check(ip string) *IPinfo {
	info := new(IPinfo)
	info.OriginalIP = ip
	if !validRegEx.Match([]byte(ip)) {
		return info
	}
	info.IsValid = true
	if IsRange(ip, bogonsArray...) {
		info.IsBogon = true
	}
	return info
}

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
