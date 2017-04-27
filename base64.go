package main

import b64 "encoding/base64"
import "fmt"

func main() {
	data := "chenlei"
	sEncode := b64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(sEncode)

	sDecode, _ := b64.StdEncoding.DecodeString(sEncode)
	// if err != nil {
	// 	fmt.Println(string(sDecode))
	// } else {
	// 	fmt.Println("Decoding error")
	// }
	fmt.Println(string(sDecode))
}
