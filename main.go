package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	hasMX = len(mxRecords) > 0

	getRecords(domain, "v=spf1", &hasSPF, &spfRecord)
	// txtRecords, err := net.LookupTXT(domain)

	// if err != nil {
	// 	log.Printf("Error: %v\n", err)
	// }

	// for _, record := range txtRecords {
	// 	if strings.HasPrefix(record, "v=spf1") {
	// 		hasSPF = true
	// 		spfRecord = record
	// 		break
	// 	}
	// }
	getRecords("_dmarc." + domain, "v=DMARC1", &hasDMARC, &dmarcRecord)
	// dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
   	// if err != nil {
	// 	log.Printf("Error: %v\n", err)
   	// }

   	// for _, record := range dmarcRecords {
	// 	if strings.HasPrefix(record, "v=DMARC1") {
	// 		hasDMARC = true
	// 		dmarcRecord = record
	// 		break
	// 	}
	// }

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}


func getRecords(domain string, prefix string, hasEntity *bool, entityRecord *string) {
	records, err := net.LookupTXT(domain)
   	if err != nil {
		log.Printf("Error: %v\n", err)
   	}

   	for _, record := range records {
		if strings.HasPrefix(record, prefix) {
			*hasEntity = true
			*entityRecord = record
			break
		}
	}
}