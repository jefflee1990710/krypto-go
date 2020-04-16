package commitment

import (
	"fmt"
	"krypto"
	"math/big"
)

// PedersenCommitment store result of a pedersen commitment
type PedersenCommitment struct {
	c big.Int
	r big.Int

	group krypto.SchnorrGroup
	g     big.Int
	h     big.Int
}

// Summary print out summary of the pedersen commitment
func (pc *PedersenCommitment) Summary() {
	fmt.Println("c = ", pc.c.Text(10))
	fmt.Println("r = ", pc.r.Text(10))
}

// Commit create commitment of message
func (pc *PedersenCommitment) Commit(group krypto.SchnorrGroup, g big.Int, h big.Int, m big.Int) {
	pc.group = group

	p := group.GetP()
	r := group.CreateRandom()

	c1 := big.Int{}
	c2 := big.Int{}

	// c1 = g^m
	c1.Exp(&g, &m, &p)
	// c2 = h^r
	c2.Exp(&h, &r, &p)

	c := big.Int{}
	c.Mul(&c1, &c2)
	c.Mod(&c, &p)

	pc.c = c
	pc.r = r
}

// Open process open stage of commitment
func (pc *PedersenCommitment) Open(group krypto.SchnorrGroup, m big.Int, c big.Int) bool {
	return false
}

// GetCommitment retrieve commitment of the commit result
func (pc *PedersenCommitment) GetCommitment() big.Int {
	return pc.c
}

// GetR retrieve r in Pedersen Commitment
func (pc *PedersenCommitment) GetR() big.Int {
	return pc.r
}
