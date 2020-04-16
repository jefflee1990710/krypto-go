package krypto

import (
	"fmt"
	"math/big"
)

// KeyPair store public key and private key
type KeyPair struct {
	pri big.Int
	pub big.Int

	p big.Int
	q big.Int
}

// GetPublicKey return public key in keypair
func (kp *KeyPair) GetPublicKey() big.Int {
	return kp.pub
}

// GetPrivateKey return private key in keypair
func (kp *KeyPair) GetPrivateKey() big.Int {
	return kp.pri
}

// GenerateNew calculate valide private key and public key in provided group setting
func (kp *KeyPair) GenerateNew(group SchnorrGroup, generator big.Int) {
	pri := group.CreateRandom()

	pub := big.Int{}
	pub.Exp(&generator, &pri, &group.p)

	kp.pri = pri
	kp.pub = pub
	kp.p = group.p
	kp.q = group.q
}

// Summary print information of p and q
func (kp *KeyPair) Summary() {
	fmt.Println("Public Key = ", kp.pub.Text(10))
	fmt.Println("Private Key = ", kp.pri.Text(10))
}
