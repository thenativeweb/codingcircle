package bloomfilter

import (
	"crypto/sha1"
	"strconv"
)

type BloomFilter struct {
	bitArray  []bool
	size      int
	numHashes int
}

func NewBloomFilter(size, numHashes int) *BloomFilter {
	return &BloomFilter{
		bitArray:  make([]bool, size),
		size:      size,
		numHashes: numHashes,
	}
}

func (bf *BloomFilter) Add(value string) {
	for i := 0; i < bf.numHashes; i++ {
		hash := sha1.New()
		hash.Write([]byte(value + strconv.Itoa(i)))
		digest := hash.Sum(nil)

		index := int(digest[0])<<24 | int(digest[1])<<16 | int(digest[2])<<8 | int(digest[3])
		index = index % bf.size
		if index < 0 {
			index += bf.size
		}

		bf.bitArray[index] = true
	}
}

func (bf *BloomFilter) Contains(value string) bool {
	for i := 0; i < bf.numHashes; i++ {
		hash := sha1.New()
		hash.Write([]byte(value + strconv.Itoa(i)))
		digest := hash.Sum(nil)

		index := int(digest[0])<<24 | int(digest[1])<<16 | int(digest[2])<<8 | int(digest[3])
		index = index % bf.size
		if index < 0 {
			index += bf.size
		}

		if !bf.bitArray[index] {
			return false
		}
	}

	return true
}
