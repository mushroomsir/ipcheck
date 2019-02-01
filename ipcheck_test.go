package ipcheck_test

import (
	"strconv"
	"testing"

	"github.com/mushroomsir/ipcheck"
	"github.com/stretchr/testify/require"
)

func TestDeepCheck(t *testing.T) {
	require := require.New(t)
	require.Equal("1.2.3.4", ipcheck.Check("1.2.3.4").OriginalIP)
	require.True(ipcheck.DeepCheck("11.9.0.0").IsSafe())

	ipcheck.AddBogonsRang("30.0.0.0/8", "11.0.0.0/8")
	ipcheck.RemoveBogonRang("224.0.0.0/3")
	cases := []struct {
		Expected bool
		Actual   bool
	}{
		{true, ipcheck.DeepCheck("cn.bing.com").IsSafe()},
		{true, ipcheck.DeepCheck("8.8.8.8").IsSafe()},
		{true, ipcheck.DeepCheck("245.0.0.8").IsSafe()},

		{false, ipcheck.DeepCheck("0.0.0.8").IsSafe()},
		{false, ipcheck.DeepCheck("256.256.256.256").IsSafe()},
		{false, ipcheck.DeepCheck("10.0.0.1").IsSafe()},
		{false, ipcheck.DeepCheck("11.0.0.0").IsSafe()},
		{false, ipcheck.DeepCheck("11.9.0.0").IsSafe()},
		{false, ipcheck.DeepCheck("11.9.8.0").IsSafe()},
		{false, ipcheck.DeepCheck("11.9.8.7").IsSafe()},
		{false, ipcheck.DeepCheck("30.0.0.0").IsSafe()},
		{false, ipcheck.DeepCheck("30.9.0.0").IsSafe()},
		{false, ipcheck.DeepCheck("30.9.8.0").IsSafe()},
		{false, ipcheck.DeepCheck("30.9.8.7").IsSafe()},
	}
	for _, c := range cases {
		require.Equal(c.Expected, c.Actual)
	}
	ipcheck.AddBogonsRang("224.0.0.0/3")
	require.False(ipcheck.DeepCheck("245.0.0.8").IsSafe())
}

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

		{false, ipcheck.Check("224.0.0.0").IsSafe()},
		{false, ipcheck.Check("225.8.0.0").IsSafe()},
	}
	for _, c := range cases {
		require.Equal(c.Expected, c.Actual)
	}
}
func TestIsRange(t *testing.T) {
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
