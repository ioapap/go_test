// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package main

import (
	"fmt" // fmt package for formatting, printing to the console.
	"log" // logging error messages and other info
	"os"  // for opening files
)

// main is the entry point for the program
func main() {
	// Open the CSV file
	csvfile, err := os.Open("customers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	// Process the CSV file
	domainCounts, err := CountDomains(csvfile)
	if err != nil {
		log.Fatal(err)
	}

	// Print the sorted domain counts
	for _, kv := range domainCounts {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
}
