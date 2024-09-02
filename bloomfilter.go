package main

import (
	"fmt"
	"hash"
	"time"

	"github.com/spaolacci/murmur3"
)

var myhasher hash.Hash32

func init() {
	myhasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
}

func myMurmurHash(key string, size int32) int32 {
	myhasher.Write([]byte(key))
	hash := myhasher.Sum32() % uint32(size)
	//Reset is important else the previous hasher values are used and we get wrong results
	myhasher.Reset()
	return int32(hash)
}

type Bloomfilter struct {
	filter []bool
	size   int32
}

func NewBloomFilter(size int32) *Bloomfilter {
	return &Bloomfilter{filter: make([]bool, size), size: size}
}

func (bf *Bloomfilter) Add(key string) {
	index := myMurmurHash(key, bf.size)
	bf.filter[index] = true
	fmt.Println("Set index:", index)
}

// TODO Find appropriate value of k and do hashing k times
func (bf *Bloomfilter) Exists(key string) bool {
	index := myMurmurHash(key, bf.size)
	return bf.filter[index]
}

func main() {
	keys := []string{"a", "b", "c", "d", "e"}
	bloom := NewBloomFilter(16)
	for _, key := range keys {
		bloom.Add(key)
	}

	for _, key := range keys {
		fmt.Println(key, bloom.Exists(key))
	}

	fmt.Println(bloom.Exists("z"))

	testMembership(bloom)
}

func testMembership(bloom *Bloomfilter) {
	stringarray := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	// Print the characters
	fmt.Println("Testing membership of characters between 'a' and 'z':")
	for _, ch := range stringarray {
		fmt.Printf(ch, bloom.Exists(string(ch)))
	}
	fmt.Println()
}
