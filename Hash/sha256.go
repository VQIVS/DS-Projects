package Hash

import (
	"crypto/sha256"
	"fmt"
)

func RunSHA256() {
	fmt.Println("Running SHA-256 example")
	data := "Hello, World!"
	hash := sha256.Sum256([]byte(data))
	fmt.Printf("SHA-256 hash of '%s': %x\n", data, hash)
}
