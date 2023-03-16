package belajar_golang_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// cara membuat embed seperti ini
//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed pp.jpg
var logo []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("PP_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}