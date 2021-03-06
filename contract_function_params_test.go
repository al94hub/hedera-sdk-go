package hedera

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

var functionName = string("f")

func TestSoliditySerialization(t *testing.T) {
	params := NewContractFunctionParams().
		AddInt32(16909060).
		AddInt64(0xffffff).
		AddString("this is a grin: 😁")

	result := params.build(&functionName)

	function, err := hex.DecodeString("29a2132d")
	assert.NoError(t, err)

	first, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000001020304")
	assert.NoError(t, err)
	second, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000ffffff")
	assert.NoError(t, err)
	thirdOffset, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000060")
	assert.NoError(t, err)
	thirdValueLength, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000014")
	assert.NoError(t, err)
	thirdValue, err := hex.DecodeString("746869732069732061206772696e3a20f09f9881000000000000000000000000")

	assert.Equal(t, 164, len(result), "Length of params does not match")
	assert.Equal(t, function, result[0:4], "Function signature doesn't match")
	assert.Equal(t, first, result[4:32+4], "First argument doesn't match")
	assert.Equal(t, second, result[32+4:(32*2)+4], "Second argument doesn't match")
	assert.Equal(t, thirdOffset, result[(32*2)+4:(32*3)+4], "ThirdOffset argument doesn't match")
	assert.Equal(t, thirdValueLength, result[(32*3)+4:(32*4)+4], "ThirsOffsetLength argument doesn't match")
	assert.Equal(t, thirdValue, result[(32*4)+4:(32*5)+4], "ThirdValue argument doesn't match")
}

func TestSolidityStringArraySerialization(t *testing.T) {
	result := NewContractFunctionParams().
		AddStringArray([]string{"one", "four"}).
		build(&functionName)

	function, err := hex.DecodeString("e9cc8780")
	assert.NoError(t, err)

	firstOffset, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000020")
	assert.NoError(t, err)

	firstLengthOfArray, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000002")
	assert.NoError(t, err)

	firstElementOffset, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000040")
	assert.NoError(t, err)

	secondElementOffset, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000080")
	assert.NoError(t, err)

	firstElementLength, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000003")
	assert.NoError(t, err)

	secondElementLength, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000004")
	assert.NoError(t, err)

	firstElementValue, err := hex.DecodeString("6f6e650000000000000000000000000000000000000000000000000000000000")
	assert.NoError(t, err)

	secondElementValue, err := hex.DecodeString("666f757200000000000000000000000000000000000000000000000000000000")
	assert.NoError(t, err)

	assert.Equal(t, 260, len(result), "Length of params does not match")
	assert.Equal(t, function, result[0:4], "Function signature doesn't match")
	assert.Equal(t, firstOffset, result[4:32+4], "FirstOffset doesn't match")
	assert.Equal(t, firstLengthOfArray, result[32+4:(32*2)+4], "FirstLengthOfArray doesn't match")
	assert.Equal(t, firstElementOffset, result[(32*2)+4:(32*3)+4], "FirstElementOffset doesn't match")
	assert.Equal(t, secondElementOffset, result[(32*3)+4:(32*4)+4], "SecondElementOffset doesn't match")
	assert.Equal(t, firstElementLength, result[(32*4)+4:(32*5)+4], " firstElementLength doesn't match")
	assert.Equal(t, firstElementValue, result[(32*5)+4:(32*6)+4], "firstElementLength doesn't match")
	assert.Equal(t, secondElementLength, result[(32*6)+4:(32*7)+4], " secondElementLength doesn't match")
	assert.Equal(t, secondElementValue, result[(32*7)+4:(32*8)+4], "secondElementLength doesn't match")
}
