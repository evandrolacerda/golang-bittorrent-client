package main

import (
	"fmt"
	"log"
	"os"
	"torrentclient/torrentfile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	fmt.Println(inPath)
	fmt.Println(outPath)

	torrentFile, err := torrentfile.Open(inPath)

	if err != nil {
		fmt.Println(err)
	}

	err = torrentFile.DownloadToFile(outPath)

	if err != nil {
		log.Fatal(err)
	}

}
