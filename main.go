package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#ck!"
	replace := strings.NewReplacer("#", "o")
	fixed := replace.Replace(broken)
	fmt.Println(fixed)
}
