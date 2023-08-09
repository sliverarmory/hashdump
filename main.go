package main

import (
	"fmt"

	"github.com/sliverarmory/secretsdump/pkg/hashdump"
)

func main() {
	result, err := hashdump.Hashdump()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
