package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Extracting modulus
	n := key.N
	fmt.Println("Modulus (n): %\n", n)

	// Extracting Public Exponent (e)
	e := key.PublicKey.E
	fmt.Println("Exponent (e): %x\n", e)

	// Saves pub key to PEM
	pubKeyBytes := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})
	fmt.Println("Public Key: \n &s", pubKeyPEM)
}
