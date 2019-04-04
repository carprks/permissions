package probe_test

import (
	"github.com/stretchr/testify/assert"
	"main/src/probe"
	"testing"
)

func TestProbe(t *testing.T) {
	tests := []struct{
		expect probe.Healthy
		err error
	}{
		{
			expect: probe.Healthy{
				Status: "pass",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		response, err := probe.Probe()
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.expect, response)
	}
}
