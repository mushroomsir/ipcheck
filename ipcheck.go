package ipcheck

import (
	"context"
	"net"
	"regexp"
	"strings"
)

var validRegEx = regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)

// Bogons list from http://www.team-cymru.org/bogon-reference.html
var bogonsArray = []string{"0.0.0.0/8", "10.0.0.0/8", "100.64.0.0/10", "127.0.0.0/8", "169.254.0.0/16",
	"172.16.0.0/12", "192.0.0.0/24", "192.0.2.0/24", "192.168.0.0/16", "198.18.0.0/15",
	"198.51.100.0/24", "203.0.113.0/24", "224.0.0.0/3", "11.0.0.0/8", "33.0.0.0/8", "30.0.0.0/8"}

// AddBogonsRang ...
func AddBogonsRang(ips ...string) {
	for _, ip := range ips {
		bogonsArray = append(bogonsArray, ip)
	}
}

// RemoveBogonRang ...
func RemoveBogonRang(ip string) {
	for i, v := range bogonsArray {
		if v == ip {
			bogonsArray = append(bogonsArray[:i], bogonsArray[i+1:]...)
			break
		}
	}
}

// IPinfo ...
type IPinfo struct {
	OriginalIP string
	IsValid    bool
	IsBogon    bool
}

// IsSafe the host is valid and not bogon
func (a *IPinfo) IsSafe() bool {
	return a.IsValid && !a.IsBogon
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

// DeepCheckWithContext check bogon ip by DNS with a context
func DeepCheckWithContext(ctx context.Context, host string) *IPinfo {
	info := new(IPinfo)
	info.OriginalIP = host
	realHosts, err := net.DefaultResolver.LookupHost(ctx, host)
	if err != nil {
		return info
	}
	info.IsValid = true
	for _, ip := range realHosts {
		if IsRange(ip, bogonsArray...) {
			info.IsBogon = true
			return info
		}
	}
	return info
}

// DeepCheck check bogon ip by DNS
func DeepCheck(host string) *IPinfo {
	return DeepCheckWithContext(context.Background(), host)
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
