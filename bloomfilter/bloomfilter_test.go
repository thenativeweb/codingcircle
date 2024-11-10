package bloomfilter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/bloomfilter"
)

func TestBloomFilter(t *testing.T) {
	bf := bloomfilter.NewBloomFilter(512, 5)

	bf.Add("www.example.com")
	bf.Add("dangerous.example.com")

	assert.True(t, bf.Contains("www.example.com"))
	assert.True(t, bf.Contains("dangerous.example.com"))
	assert.False(t, bf.Contains("www.thenativeweb.io"))
}
