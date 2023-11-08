// Tests

package main

import (
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCountDomains(t *testing.T) {
	const csvData = `first_name,last_name,email,gender,ip_address
Mildred,Hernandez,mhernandez0@github.io,Female,38.194.51.128
Bonnie,Ortiz,bortiz1@cyberchimps.com,Female,197.54.209.129`

	reader := strings.NewReader(csvData)
	got, err := CountDomains(reader)
	assert.NoError(t, err) // Assert that no error should have been returned.

	want := []KeyValue{{"cyberchimps.com", 1}, {"github.io", 1}}
	assert.Equal(t, want, got) // Assert that the 'got' slice should equal the 'want' slice.
}

