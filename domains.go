package main

import (
	"encoding/csv" // csv package for parsing CSV files.
	"fmt"
	"io"      // basic interfaces for I/O operations
	"sort"    // for sorting slices.
	"strings" // for string manipulation functions
)

// KeyValue is a struct to hold the key-value pairs.
type KeyValue struct {
	Key   string
	Value int
}

// It reads from an io.Reader and returns a sorted slice of email domain counts.
func CountDomains(reader io.Reader) ([]KeyValue, error) {
	csvReader := csv.NewReader(reader) // Create a new CSV reader from the io.Reader

	// Skip the header
	if _, err := csvReader.Read(); err != nil { // Attempt to read (and discard) the header row.
		return nil, err
	}

	// Create a map to count the unique domains
	domainMap := make(map[string]int)

	// Iterate through the records
	for {
		// Read each record from csv.
		record, err := csvReader.Read()
		if err == io.EOF {
			break // If the end of the file is reached, break out of the loop.
		}
		if err != nil {
			if _, ok := err.(*csv.ParseError); ok { // Check if the error is a parsing one.
				return nil, fmt.Errorf("file contains malformed CSV data: %v", err)
			}
			return nil, fmt.Errorf("an error occured while reading the CSV file: %v", err)
		}

		// Check if the record is empty (this would be the case for an empty line)
		if len(record) == 0 {
			continue // Skip this record and continue to the next one.
		}

		// Extract the email from the record.
		email := record[2]
		// Check if the email contains the "@" character.
		if atIdx := strings.Index(email, "@"); atIdx != -1 {
			// Split the email string by "@" and get the domain part.
			domain := email[atIdx+1:]
			// Increment the count for this domain.
			domainMap[domain]++
		}
	}

	// Convert the map to a slice of key-value pairs
	var ss []KeyValue
	for k, v := range domainMap {
		ss = append(ss, KeyValue{k, v})
	}

	// Sort the slice based on the frequency of the domains, then by the domain name
	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})

	return ss, nil
}
