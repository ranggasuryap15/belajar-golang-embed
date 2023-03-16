package belajar_golang_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

// cara membuat embed seperti ini
//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}