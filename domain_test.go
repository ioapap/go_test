// Tests

package main

import (
	"os"
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

func TestCountDomainsFileDoesNotExist(t *testing.T) {
	_, err := CountDomains(os.Stdin)                                                                                 // assuming we're not actually reading from Stdin
	assert.Error(t, err, "An error was expected when trying to count domains from a non-existent file, but got nil") // Assert that an error should have been returned.
}

func TestCountDomainsInvalidCSV(t *testing.T) {
	const invalidCSVData = `first_name|last_name|email|gender|ip_address
Mildred|Hernandez|mhernandez0@github.io|Female|38.194.51.128
Bonnie|Ortiz|bortiz1@cyberchimps.com|Female|197.54.209.129` // The | instead of ,

	reader := strings.NewReader(invalidCSVData)
	_, err := CountDomains(reader)
	assert.Error(t, err, "An error was expected when trying to count domains from an invalid CSV, but got nil")
}

func TestCountDomainsWithEmptyLine(t *testing.T) {
	const csvDataWithEmptyLine = `first_name,last_name,email,gender,ip_address

Mildred,Hernandez,mhernandez0@github.io,Female,38.194.51.128
Bonnie,Ortiz,bortiz1@cyberchimps.com,Female,197.54.209.129`

	reader := strings.NewReader(csvDataWithEmptyLine)
	got, err := CountDomains(reader)
	assert.NoError(t, err) // Assert no error should have been returned despite the empty line

	want := []KeyValue{{"cyberchimps.com", 1}, {"github.io", 1}}
	assert.Equal(t, want, got) // Assert that the 'got' slice should equal the 'want' slice.
}
