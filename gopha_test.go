// Copyright 2013 Unknown
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package gopha is an implementation of Perceptual Hash Algorithm in Go programming language.
package gopha

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"testing"
)

func TestPHA(t *testing.T) {
	// Load picture from file.
	infile, err := os.Open("./testdata/pic_original.jpg")
	if err != nil {
		t.Errorf("Load picture: %s.", err)
	}
	defer infile.Close()

	// Decode picture.
	srcImg, _, err := image.Decode(infile)
	if err != nil {
		t.Errorf("Decode picture: %s.", err)
	}

	fg := PHA(srcImg)
	fmt.Println("Fingerprint:", fg)
}
