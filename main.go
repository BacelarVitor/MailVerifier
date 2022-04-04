package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecodr\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: coud not read from input %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecodr string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	hasMX = len(mxRecords) > 0
	
}
