package main

import (
 	flag "github.com/ogier/pflag"
 	"fmt"
	"os"
)

var (
 	plainText string
	shift int
 	)

func cipher(text string, shift int, direction int) string {
	internalShift, offset := rune(shift), rune(26)

	runes := []rune(text)
	for index, character := range runes {
		switch direction {
		case -1:
			if character >= 'a'+internalShift && character <= 'z' ||
				character >= 'A'+internalShift && character <= 'Z' {
				character = character - internalShift
			} else if character >= 'a' && character < 'a'+internalShift ||
				character >= 'A' && character < 'A'+internalShift {
				character = character - internalShift + offset
			}
		case +1:
			if character >= 'a' && character <= 'z'-internalShift ||
				character >= 'A' && character <= 'Z'-internalShift {
				character = character + internalShift
			} else if character > 'z'-internalShift && character <= 'z' ||
				character > 'Z'-internalShift && character <= 'Z' {
				character = character + internalShift - offset
			}
		}
		runes[index] = character
	}
	return string(runes)
}




func encode(text string, shift int) string {
	return cipher(text, shift, -1)
}

func decode(text string, shift int) string {
	return cipher(text, shift, 1)
}


func main() {
	flag.Parse()

 	if flag.NFlag() == 0 {
 		fmt.Printf("Usage: %s [options]\n", os.Args[0])
 		fmt.Println("Options:")
 		flag.PrintDefaults()
 		os.Exit(1)
 	}

 	if flag.NFlag() > 0 {
 		if os.Args[1] == "-e" {
			encoded := encode(plainText, shift)
			fmt.Printf("Encrypted string: %s\n", encoded)
		}
		if os.Args[1] == "-d" {
			decoded := decode(plainText, shift)
			fmt.Printf("Decrypted string: %s\n", decoded)

		}
	}

 }

 func init() {
 	flag.StringVarP(&plainText, "encrypt", "e", "", "encrypts a string")
	 flag.StringVarP(&plainText, "decrypt", "d", "", "decrypts a string")
	 flag.IntVar(&shift, "shift", 3, "sets the cipher shift")
 }