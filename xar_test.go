// Copyright (c) 2011 Mikkel Krautz <mikkel@krautz.dk>
// The use of this source code is goverened by a BSD-style
// license that can be found in the LICENSE-file.

package goxar

import (
	"fmt"
	"testing"
)

func TestOpenFile(t *testing.T) {
	r, err := OpenReader("payload.xar")
	if err != nil {
		t.Errorf("%v", err)
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
