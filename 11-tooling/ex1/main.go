package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed udhr/en.txt
var english string

//go:embed udhr/it.txt
var italian string

// registry mapping language -> text
var rightsByLang = map[string]string{
	"en": english,
	"it": italian,
}

func main() {
	// check pamars
	if len(os.Args) < 2 {
		log.Fatal("Missing lang param (it | en)")
	}
	lang := strings.ToLower(os.Args[1])

	text, ok := rightsByLang[lang]
	if !ok {
		log.Fatalf("unsupported language: %s", lang)
	}

	fmt.Println(text)
}
