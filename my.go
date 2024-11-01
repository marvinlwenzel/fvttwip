package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp/caddyauth"
	"golang.org/x/crypto/pbkdf2"
)

func init() {
	caddy.RegisterModule(FvttCrypter{})
}

// Gizmo is an example; put your own type here.
type FvttCrypter struct{}

func (FvttCrypter) Hash(plaintext []byte) ([]byte, error) {
	saltRandomness := make([]byte, 32)
	// [25 27 ... 40] (len=32)
	_, err := rand.Read(saltRandomness)
	if err != nil {
		return nil, err
	}

	// 191bebf98a1111dc092dfb2de19f46a1a59df14db8e234654b68452562b28a28
	saltHexString := hex.EncodeToString(saltRandomness)

	// [49 57 49 98 ... 50 56] (len=64)
	saltBytes := []byte(saltHexString)

	hash := pbkdf2.Key(plaintext, saltBytes, 1000, 64, sha512.New)
	hexHashBytes := []byte(hex.EncodeToString(hash))

	arr := make([]byte, 201)
	n := 0
	copy(arr[n:], []byte("$fvtt1a$"))
	n = 8
	copy(arr[n:], saltBytes)
	n += len(saltBytes)
	arr[n] = '$'
	n += 1
	copy(arr[n:], hexHashBytes)
	n += len(hexHashBytes)
	return arr, nil
}
func (FvttCrypter) FakeHash() []byte {
	fake := []byte("$fvtt1a$191bebf98a1111dc092dfb2de19f46a1a59df14db8e234654b68452562b28a28$c27139bff8c68a5f9f8228b7f00b41db5fc001841497d99d5422aa505f1a83c1ed37bb159857bbb8bc44905f47d3873419e26e0c3ce37d6c8b1544b645f07e1b")
	return fake
}
func (FvttCrypter) Compare(hashed, plaintext []byte) (bool, error) {
	fmt.Println("MLW Compare")
	fmt.Println("MLW Compare")
	fmt.Println("MLW Compare")
	fmt.Println("MLW Compare")
	fmt.Println("hashed")
	fmt.Println(hashed)
	fmt.Println(string(hashed))
	fmt.Println("plaintext")
	fmt.Println(plaintext)
	fmt.Println(string(plaintext))
	// err := new(FvttCrypter)
	// if err != nil {
	// 	return false, err
	// }
	return true, nil
}

// CaddyModule returns the Caddy module information.
func (FvttCrypter) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.authentication.hashes.fvttcrypt",
		New: func() caddy.Module { return new(FvttCrypter) },
	}
}

// Interface guards
var (
	_ caddyauth.Comparer = (*FvttCrypter)(nil)
	_ caddyauth.Hasher   = (*FvttCrypter)(nil)
)
