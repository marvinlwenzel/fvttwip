package mymodule

import (
	"encoding/hex"
	"fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp/caddyauth"
)

func init() {
	caddy.RegisterModule(FvttCrypter{})
}

// Gizmo is an example; put your own type here.
type FvttCrypter struct{}

func (FvttCrypter) Hash(plaintext []byte) ([]byte, error) {
	fmt.Println("MLW Hash")
	fmt.Println("MLW Hash")
	fmt.Println("MLW Hash")
	fmt.Println("MLW Hash")
	fmt.Println(plaintext)
	hexString := "4c6561726e20476f21"
	return hex.DecodeString(hexString)
}
func (FvttCrypter) FakeHash() []byte {
	hexString := "4c6561726e20476f21"
	h, _ := hex.DecodeString(hexString)
	return h
}
func (FvttCrypter) Compare(hashed, plaintext []byte) (bool, error) {
	fmt.Println("MLW Compare")
	fmt.Println("MLW Compare")
	fmt.Println("MLW Compare")
	fmt.Println("MLW Compare")
	fmt.Println(hashed)
	fmt.Println(plaintext)
	// err := new(FvttCrypter)
	// if err != nil {
	// 	return false, err
	// }
	return true, nil
}

// CaddyModule returns the Caddy module information.
func (FvttCrypter) CaddyModule() caddy.ModuleInfo {
	fmt.Println("MLW caddy.ModuleInfo")
	fmt.Println("MLW caddy.ModuleInfo")
	fmt.Println("MLW caddy.ModuleInfo")
	fmt.Println("MLW caddy.ModuleInfo")
	return caddy.ModuleInfo{
		ID:  "http.authentication.hashes.fvttcryp",
		New: func() caddy.Module { return new(FvttCrypter) },
	}
}

// Interface guards
var (
	_ caddyauth.Comparer = (*FvttCrypter)(nil)
	_ caddyauth.Hasher   = (*FvttCrypter)(nil)
)
