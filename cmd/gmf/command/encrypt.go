package command

import (
	"fmt"

	encryptiion "gitlab.com/indev-moph/fiber-api/pkg/encryption"
)

func GenRSAKey() {
	pk, err := encryptiion.GenerateKeyPair(128)
	if err != nil {
		fmt.Println(err)
	}
	key, _ := encryptiion.PrivateKeyToBytes(pk)
	pub, _ := encryptiion.PublicKeyToBytes(&pk.PublicKey)
	fmt.Printf("PrivateKey: %s\nPublicKey: %s", key, pub)
}

func Encrypt(msg []byte, pubHex string) {

}

func Decrypt() {

}
