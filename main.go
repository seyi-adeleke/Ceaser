package main

import (
 	flag "github.com/ogier/pflag"
 	"fmt"
 	"os"
 	)

var (
 	plainText  string
 	)

func main() {
	flag.Parse()

 	if flag.NFlag() == 0 {
 		fmt.Printf("Usage: %s [options]\n", os.Args[0])
 		fmt.Println("Options:")
 		flag.PrintDefaults()
 		os.Exit(1)
 	}
 	fmt.Printf("String to encrypt: %s\n", plainText)
 }

 func init() {
 	flag.StringVarP(&plainText, "encrypt", "e", "", "encrypts a string")
 }