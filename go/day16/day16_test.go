package day16

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParsingVersion(t *testing.T) {
	for _, tt := range []struct {
		packet   string
		expected uint64
	}{
		{
			"8A004A801A8002F478",
			16,
		},
		{
			"620080001611562C8802118E34",
			12,
		},
		{
			"C0015000016115A2E0802F182340",
			23,
		},
		{
			"A0016C880162017C3686B18A3D4780",
			31,
		},
	} {
		t.Run(tt.packet, func(t *testing.T) {
			packet, err := Parse(strings.NewReader(tt.packet))
			require.NoError(t, err)

			assert.Equal(t, tt.expected, packet.Version())
		})
	}

}
