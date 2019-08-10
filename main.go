package main

import (
	"fmt"
	"log"
	"net/url"

	gitUrls "github.com/whilp/git-urls"
)

func main() {
	original := "https://github.com/whilp/git-urls"
	fmt.Println("Original:", original)
	parsed, err := gitUrls.Parse(original)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Parsed:", parsed.String())
	parsed.User = nil
	fmt.Println("Corrected:", parsed.String())
	parsed.User = url.User("")
	fmt.Println("Wrong again:", parsed.String())
}
