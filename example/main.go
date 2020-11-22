package main

import (
	"fmt"

	"github.com/mathantunes/go-vies"
)

func main() {
	v := vies.NewValidator(nil)
	resp, err := v.Validate("FI25160553")
	if err != nil {
		fmt.Print(fmt.Errorf("%s", err.Error()))
	}
	if !resp.Valid {
		fmt.Print(fmt.Errorf("It seems like the VAT provided is not valid 😕"))
	} else {
		fmt.Printf("Yay! it is a valid VAT, look %v", resp)
	}
	return
}
