package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanBuiltIn(t *testing.T) {
	ports, err := GoScan("google.com", []int{80})
	assert.Nil(t, err)
	assert.Len(t, ports, 1)
	assert.Equal(t, ports[0].Port, 80)
	assert.True(t, ports[0].Open)
}

func TestScanNmap(t *testing.T) {
	ports, err := NmapScan("google.com")
	assert.Nil(t, err)
	assert.Len(t, ports, 2)
	assert.Equal(t, ports[0].Port, 80)
	assert.True(t, ports[0].Open)
	assert.Equal(t, ports[1].Port, 443)
	assert.True(t, ports[1].Open)
}
