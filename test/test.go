package main

import (
	"fmt"

	opus "github.com/thebakedpotato/go-opus"
)

func main() {
	err := opus.CreateEncoder()
	if err == opus.Ok {
		fmt.Println("everything is okie dokie")
	}

	err = opus.CreateDecoder()
	if err == opus.Ok {
		fmt.Println("Decoder success")
	}
}
