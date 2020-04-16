package krypto

import (
	"math/big"
)

// SchnorrGroup ...
type SchnorrGroup struct {
	p big.Int
	q big.Int
}

// New init Schnorr Group with p and q
func (sg *SchnorrGroup) New(n int64) {
	var ONE = new(big.Int).SetInt64(1)
	var TWO = new(big.Int).SetInt64(2)
	next := true
	for next {
		s := big.Int{}
		s.SetInt64(n)
		q := CreateRandomPrime(s)
		p := big.Int{}
		p.Mul(TWO, &q)
		p.Add(&p, ONE)

		sg.p = p
		sg.q = q

		if p.ProbablyPrime(40) {
			next = false
		}
	}
}

// GetGenerator get generate base on Schnorr Group requirement
func (sg *SchnorrGroup) GetGenerator() *big.Int {
	var ONE = new(big.Int).SetInt64(1)
	var TWO = new(big.Int).SetInt64(2)

	P1 := new(big.Int)
	P1.Sub(&sg.p, ONE)
	h := CreateRandomBetween(*TWO, *P1)
	g := new(big.Int)
	g.Exp(&h, TWO, &sg.p)
	return g
}
