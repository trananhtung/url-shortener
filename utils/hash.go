package utils

import (
	"fmt"
	"hash/fnv"

	"github.com/mr-tron/base58"
)

func HashURL(url string) string {
	hash := fnv.New32a()
	hash.Write([]byte(url))
	hashValue := hash.Sum32()

	// Convert the hash value to bytes
	hashBytes := []byte(fmt.Sprintf("%d", hashValue))

	// Encode the hash bytes using Base58
	encodedHash := base58.Encode(hashBytes)

	return encodedHash
}
