package main

import (
	"fmt"
	"strings"
)

func main() {

	var os string = "Madam"
        var rs string = ""
	var l = len(os)

	for i := l - 1; i >= 0; i -- {
		rs = rs + string(os[i])
	}
	if strings.ToLower(os) == strings.ToLower(rs) {
		fmt.Println("The string is Palindrome");
	} else {
		fmt.Println("The string is not a Palindrome");
	}
}