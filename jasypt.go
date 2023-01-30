package main

import (
	"errors"
	"fmt"
	"github.com/Mystery00/go-jasypt"
	"github.com/Mystery00/go-jasypt/iv"
	"github.com/Mystery00/go-jasypt/salt"
)

const defaultAlgorithm = "PBEWithHMACSHA512AndAES_256"

func Encrypt(message string, password string, algorithm string) (string, error) {
	if len(algorithm) == 0 {
		algorithm = defaultAlgorithm
	} else if !isValidAlgorithm(algorithm) {
		return "", errors.New(fmt.Sprintf("Unsupported algorithm: %s", algorithm))
	}

	// create a new instance of jasypt
	encryptor := jasypt.New(algorithm, jasypt.NewConfig(
		jasypt.SetPassword(password),
		jasypt.SetSaltGenerator(salt.RandomSaltGenerator{}),
		jasypt.SetIvGenerator(iv.RandomIvGenerator{}),
	))

	// encrypt the message
	return encryptor.Encrypt(message)
}

func Decrypt(encode string, password string, algorithm string) (string, error) {
	if len(algorithm) == 0 {
		algorithm = defaultAlgorithm
	} else if !isValidAlgorithm(algorithm) {
		return "", errors.New(fmt.Sprintf("Unsupported algorithm: %s", algorithm))
	}

	// create a new instance of jasypt
	encryptor := jasypt.New(algorithm, jasypt.NewConfig(
		jasypt.SetPassword(password),
		jasypt.SetSaltGenerator(salt.RandomSaltGenerator{}),
		jasypt.SetIvGenerator(iv.RandomIvGenerator{}),
	))
	// decrypt the message
	return encryptor.Decrypt(encode)
}

func isValidAlgorithm(algorithm string) bool {
	for _, currentAlgo := range supportedAlgorithms() {
		if currentAlgo == algorithm {
			return true
		}
	}
	return false
}

func supportedAlgorithms() []string {
	return []string{"PBEWithHMACSHA512AndAES_256", "PBEWithMD5AndDES"}
}
