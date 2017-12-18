package ipcheck_test

import (
	"strconv"
	"testing"

	"github.com/mushroomsir/ipcheck"
	"github.com/stretchr/testify/require"
)

func TestCheck(t *testing.T) {
	require := require.New(t)

	require.Equal("1.2.3.4", ipcheck.Check("1.2.3.4").OriginalIP)
	cases := []struct {
		Expected bool
		Actual   bool
	}{
		{true, ipcheck.Check("10.10.10.10").IsValid},
		{false, ipcheck.Check("256.256.256.256").IsValid},
		// isBogon Tests
		{true, ipcheck.Check("10.0.0.1").IsBogon},
		{false, ipcheck.Check("8.8.8.8").IsBogon},
	}
	for _, c := range cases {
		require.Equal(c.Expected, c.Actual)
	}
}
func TestIPcheck(t *testing.T) {
	require := require.New(t)

	cases := []struct {
		Expected bool
		Actual   bool
	}{
		{false, ipcheck.IsRange("::1", "::2/128")},
		{false, ipcheck.IsRange("::1", []string{"::2", "::3/128"}...)},
		{true, ipcheck.IsRange("::1", "::1")},
		{true, ipcheck.IsRange("::1", []string{"::1"}...)},
		{true, ipcheck.IsRange("2001:cdba::3257:9652", "2001:cdba::3257:9652/128")},
		{ipcheck.IsRange("::1", "::1"), ipcheck.IsRange("::1", []string{"::1", "::1", "::1"}...)},
		{true, ipcheck.IsRange("2001:cdba:0000:0000:0000:0000:3257:9652", "2001:cdba:0:0:0:0:3257:9652")},
		{true, ipcheck.IsRange("2001:cdba:0000:0000:0000:0000:3257:9652", "2001:cdba::3257:9652")},
		{true, ipcheck.IsRange("2001:cdba:0:0:0:0:3257:9652", "2001:cdba:0000:0000:0000:0000:3257:9652/128")},
		{false, ipcheck.IsRange("102.1.5.0", "102.1.5.1")},
		{true, ipcheck.IsRange("102.1.5.0", "102.1.5.0")},
		{true, ipcheck.IsRange("102.1.5.92", "102.1.5.0/24")},
		{true, ipcheck.IsRange("102.1.5.92", []string{"102.1.5.0/24", "192.168.1.0/24"}...)},
		{true, ipcheck.IsRange("0:0:0:0:0:FFFF:222.1.41.90", "222.1.41.90")},
		{true, ipcheck.IsRange("0:0:0:0:0:FFFF:222.1.41.90", "222.1.41.0/24")},
		// should fail when comparing IPv6 with IPv4
		{false, ipcheck.IsRange("::5", "102.1.1.2")},
		{false, ipcheck.IsRange("::1", "0.0.0.1")},
		{false, ipcheck.IsRange("195.58.1.62", "::1/128")},
	}
	for _, c := range cases {
		require.Equal(c.Expected, c.Actual)
	}
	for i := 0; i <= 255; i++ {
		require.Equal(true, ipcheck.IsRange("102.1.5."+strconv.Itoa(i), "102.1.5.0/24"))
	}
}
