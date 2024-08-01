package main

import (
	"fmt"
	"os"

	"github.com/chennqqi/xar"
)

func main() {
	binary, _ := os.Executable()
	if len(os.Args) != 2 {
		fmt.Println("usage", binary, "<xar filename>")
		return
	}
	filename := os.Args[1]
	r, err := xar.OpenReader(filename)
	if err != nil {
		fmt.Println("read", filename, err)
		return
	}

	if r.HasSignature() {
		fmt.Printf("XAR archive has a signature. ValidSignature=%v\n", r.ValidSignature())
		fmt.Printf("Certificates = %v\n", r.Certificates)
		fmt.Printf("\n")
	} else {
		fmt.Printf("XAR archive does not have a signature.\n")
	}

	// dump all files in the xar archive
	for _, xf := range r.File {
		fmt.Printf("name:            %v\n", xf.Name)
		fmt.Printf("type:            %v\n", xf.Type)
		fmt.Printf("info:            %v\n", xf.Info)
		fmt.Printf("valid checksum:  %v\n", xf.VerifyChecksum())
		fmt.Printf("\n")
	}
}
