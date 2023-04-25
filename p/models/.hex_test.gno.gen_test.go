// Code generated by github.com/gnolang/gno. DO NOT EDIT.

//go:build gno,test
// +build gno,test

package models

import (
	"fmt"
	"testing"
)

func TestHex(t *testing.T) {
	hexEncoder := EncodingHex{}
	data := []byte("Hello, World!")

	// Encode to hex
	encoded := hexEncoder.EncodeToString(data)
	fmt.Println("Encoded:", encoded)

	// Decode from hex
	decoded, err := hexEncoder.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))
}