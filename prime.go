package krypto

import (
	"crypto/rand"
	"math/big"
)

// CreateRandomPrime create a random prime number within bit length
func CreateRandomPrime(bitLength big.Int) big.Int {
	var ZERO = new(big.Int).SetInt64(0)
	var TWO = new(big.Int).SetInt64(2)
	var THREE = new(big.Int).SetInt64(3)

	if bitLength.Cmp(TWO) == -1 {
		panic("bitLength cannot smaller then 2")
	} else if bitLength.Cmp(TWO) == 0 {
		return *THREE
	} else {
		success := false
		for !success {
			random := CreateRandom(bitLength)
			n := new(big.Int).Mod(&random, TWO)
			if n.Cmp(ZERO) != 0 { // If the number is even number, it is not prime.
				if random.ProbablyPrime(40) { // Check prime by rabin miller test
					success = true
					return random
				}
			}
		}
	}
	return *ZERO
}

// CreateRandomBetween create random number between start and end inclusively
func CreateRandomBetween(start big.Int, end big.Int) big.Int {
	var ONE = new(big.Int).SetInt64(1)

	len := new(big.Int)
	len.Sub(&end, &start)
	len.Add(len, ONE)

	n, err := rand.Int(rand.Reader, len)
	if err != nil {
		panic(err)
	}
	r := new(big.Int)
	r.Add(n, &start)
	return *r
}

// CreateRandom create a random value within bit length
func CreateRandom(bitLength big.Int) big.Int {
	max := GetMaxInt(bitLength)
	n, err := rand.Int(rand.Reader, &max)
	if err != nil {
		panic(err)
	}
	return *n
}

// GetMaxInt get maximun value of a specific bit length integer
func GetMaxInt(bitLength big.Int) big.Int {
	var ONE = new(big.Int).SetInt64(1)
	var TWO = new(big.Int).SetInt64(2)

	z1 := new(big.Int).Exp(TWO, &bitLength, nil)
	z1 = z1.Sub(z1, ONE)
	return *z1
}
