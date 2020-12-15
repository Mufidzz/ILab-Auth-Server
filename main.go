package main

import (
	"./auth"
	"./config"
	"fmt"
)

func main() {
	k, _ := auth.ImportKey(config.GetKeyStorage())
	str, _ := auth.GenerateJWT(k, auth.Body{
		IAT: 112,
		ISS: "Hello",
		SLV: "AAA",
		UID: "BBB",
		RO:  "UAI",
	})

	fmt.Printf("%s\n", str)

	jwt1 := str
	b, err := auth.ValidateJWT([]byte(jwt1), k)
	if err != nil {
		panic(err)
	}

	fmt.Println(b)
}
