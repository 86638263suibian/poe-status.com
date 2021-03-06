package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanBuiltIn(t *testing.T) {
	ports, err := GoScan("google.com", []int{80})
	assert.Nil(t, err)
	assert.Len(t, ports, 1)
	assert.Equal(t, ports[0].Port, int32(80))
	assert.True(t, ports[0].Open)
}

func TestScanNmap(t *testing.T) {
	ports, err := NmapScan("google.com")
	assert.Nil(t, err)
	assert.True(t, len(ports) > 1)
	for _, port := range ports {
		if port.Port == 80 || port.Port == 443 {
			assert.True(t, ports[0].Open)
		}
	}
}

func TestNmapAvailable(t *testing.T) {
	available := NmapAvailable()
	assert.True(t, available)
}

func TestPingAvailable(t *testing.T) {
	available := PingAvailable()
	assert.True(t, available)
}
