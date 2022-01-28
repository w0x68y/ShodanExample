package main

import (
	"BlackHatGo/part3/shodan/shodan"
	"fmt"
	"log"
	"os"
)

const apiKey = "KeyHere"

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan serchterm")
	}
	//apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("Query Credits: %d\nScan Credits: %d\n\n", info.QueryCredits, info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}