package main

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type CompressedGene struct {
	gene string
}

var bitString = big.NewInt(1)

func main() {
	_original := "TAGGGATTAACCGTTATATATATATAGCCATGGATCGATTATATAGGGATTAACCGTTATATATATATAGCCATGGATCGATTATA"
	original := ""
	for i := 0; i < 100; i++ {
		original += _original
	}
	fmt.Printf("original is %d bytes.\n", len(original))
	compressed := CompressedGene{original}
	err := compressed.compress(compressed.gene)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("compressed is %v.\n", bitString)
	fmt.Printf("compressed is %d bytes.\n", len(bitString.String()))
	decompress_str, err := compressed.decompress()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("original and decompress are the same: %v\n", original == decompress_str)

}

func (c *CompressedGene) compress(gene string) error {

	for _, nucleotide := range strings.ToUpper(gene) {
		bitString.Lsh(bitString, 2)
		if nucleotide == 'A' {
			bitString = bitString.Or(bitString, big.NewInt(0b00))
		} else if nucleotide == 'C' {
			bitString = bitString.Or(bitString, big.NewInt(0b01))
		} else if nucleotide == 'G' {
			bitString = bitString.Or(bitString, big.NewInt(0b10))
		} else if nucleotide == 'T' {
			bitString = bitString.Or(bitString, big.NewInt(0b11))
		} else {
			fmt.Println("Invalid Nucleotide:%v\n", bitString)
			return errors.New(fmt.Sprintf("Invalid Nucleotide:%s\n", bitString))
		}
	}
	return nil
}

func (c *CompressedGene) decompress() (string, error) {
	gene := ""
	l := 0
	bitStringTmp := new(big.Int)
	bitStringTmp.Set(bitString)
	for bitStringTmp.Cmp(big.NewInt(1)) == 1 {
		bitStringTmp.Div(bitStringTmp, big.NewInt(2))
		l++
	}
	for i := 0; i < l; i = i + 2 {
		_bits := big.NewInt(0)
		_bits.Rsh(bitString, uint(i))
		_bits.And(_bits, big.NewInt(0b11))
		if _bits.Cmp(big.NewInt(0b00)) == 0 {
			gene += "A"
		} else if _bits.Cmp(big.NewInt(0b01)) == 0 {
			gene += "C"
		} else if _bits.Cmp(big.NewInt(0b10)) == 0 {
			gene += "G"
		} else if _bits.Cmp(big.NewInt(0b11)) == 0 {
			gene += "T"
		} else {
			return "", errors.New(fmt.Sprintf("Invalid _bits:%s", bitString))
		}
	}
	return reverse(gene), nil
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
