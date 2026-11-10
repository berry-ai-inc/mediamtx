package rtsp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// berry's, unbracketed IPv6 hosts in URLs sent by cameras
func TestBracketIPv6Host(t *testing.T) {
	for _, ca := range []struct {
		name string
		in   string
		out  string
	}{
		{
			"bare host",
			"rtsp://fe80::20f:7cff:fe1b:eee4/media/video1/video",
			"rtsp://[fe80::20f:7cff:fe1b:eee4]/media/video1/video",
		},
		{
			"bare host, rtsps",
			"rtsps://fe80::20f:7cff:fe1b:eee4/media/video1/",
			"rtsps://[fe80::20f:7cff:fe1b:eee4]/media/video1/",
		},
		{
			"bare host with port parses as host:port already",
			"rtsp://fe80::20f:7cff:fe1b:eee4:554/media/video1",
			"rtsp://fe80::20f:7cff:fe1b:eee4:554/media/video1",
		},
		{
			"bare host ending in a decimal group parses already",
			"rtsp://2001:db8::1/media/video1",
			"rtsp://2001:db8::1/media/video1",
		},
		{
			"bracketed host",
			"rtsp://[fe80::1]:554/media/video1",
			"rtsp://[fe80::1]:554/media/video1",
		},
		{
			"IPv4 host",
			"rtsp://192.168.1.10:554/media/video1",
			"rtsp://192.168.1.10:554/media/video1",
		},
		{
			"hostname",
			"rtsp://cafe:554/media/video1",
			"rtsp://cafe:554/media/video1",
		},
		{
			"relative path",
			"/media/video1/",
			"/media/video1/",
		},
	} {
		t.Run(ca.name, func(t *testing.T) {
			require.Equal(t, ca.out, bracketIPv6Host(ca.in))
		})
	}
}
