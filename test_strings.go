package main

import (
	"encoding/binary"
	"fmt"
	"strings"
)

func main() {
	name := "bruce allen kelley, jr"

	fmt.Println("Strings:")
	fmt.Println("  name = ", name)
	fmt.Println("  name[0] = ", name[0])
	fmt.Println("  name[1:3] = ", name[1:3])
	fmt.Println("  name[:3] = ", name[:3])

	pos := strings.IndexAny(name, ",")
	if pos == -1 {
		fmt.Println("Comma was not found")
	} else {
		fmt.Println("Comma found at position ", pos)
	}

	// testing indexing to -17 for a '-'
	long_name := "kube-dns-57dd96bb49-bmcmw"
	short_name := long_name
	long_name_b := []byte(long_name) // string -> []byte
	len_name := binary.Size(long_name_b)
	var ascii_dash byte
	ascii_dash = 45

	pos = strings.Index(long_name, "-ip-")
	fmt.Println("-ip- pos=", pos)

	pos = strings.Index(long_name, "dns")
	fmt.Println("dns pos=", pos)

	dash_pos := len_name - 17
	fmt.Println("Dash_pos=", dash_pos)
	if long_name_b[dash_pos] == ascii_dash {
		short_name = long_name[0:dash_pos]
	}

	fmt.Println("Short name = ", short_name)
	fmt.Println("name=", long_name, " len name as byte[] array bs[] using: binary.Size(bs) is ", len_name)
	fmt.Println("name=", long_name, " len name as byte[] array bs[] using: len(bs) =", len(long_name_b))
	fmt.Println("bs[-15]=", long_name_b[len_name-15])
	fmt.Println("bs[-16]=", long_name_b[len_name-16])
	fmt.Println("bs[-17]=", long_name_b[len_name-17])
	if long_name_b[len_name-17] == 45 {
		fmt.Println("DASH at location -17")
	} else {
		fmt.Println("NO DASH at location -17")
	}

	// using binary.Size against byte array
	fmt.Println("Length using binary.Size...")
	thousandBytes := make([]byte, 1000)
	tenBytes := make([]byte, 10)
	fmt.Println("ten bytes=", binary.Size(tenBytes))
	fmt.Println("thousand bytes=", binary.Size(thousandBytes))

	/*
		Here's a simple program that prints a string constant with a single
		character three different ways, once as a plain string,
		once as an ASCII-only quoted string, and once as individual bytes
		in hexadecimal. To avoid any confusion, we create a "raw string", enclosed
		by back quotes, so it can contain only literal text. (Regular strings,
		enclosed by double quotes, can contain escape sequences).
	*/
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

}
