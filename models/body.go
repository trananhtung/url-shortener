package models

import (
	"fmt"
	"hash/fnv"

	"github.com/asaskevich/govalidator"
	"github.com/mr-tron/base58"
)

type CreateShortenURL struct {
	URL string `json:"url" binding:"required" valid:"url"`
}

func (url *CreateShortenURL) Validate() error {
	_, err := govalidator.ValidateStruct(url)
	if err != nil {
		return err
	}

	return nil

}

func (url *CreateShortenURL) HashURL() string {
	hash := fnv.New32a()
	hash.Write([]byte(url.URL))
	hashValue := hash.Sum32()

	// Convert the hash value to bytes
	hashBytes := []byte(fmt.Sprintf("%d", hashValue))

	// Encode the hash bytes using Base58
	encodedHash := base58.Encode(hashBytes)

	return encodedHash
}
