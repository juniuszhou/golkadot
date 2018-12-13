package bn

import (
	"math"
	"math/big"
	"sort"

	"github.com/c3systems/go-substrate/common/hexutil"
)

// FromHex creates a math/big big number from a hex string.
func FromHex(hexStr string) (*big.Int, error) {
	return hexutil.ToBN(hexStr, false)
}

// ToBN creates a BN value from a number input
func ToBN(ivalue interface{}, isLittleEndian bool) *big.Int {
	switch v := ivalue.(type) {
	case *big.Int:
		return v
	case string:
		i := new(big.Int)
		i.SetString(v, 10)
		return i
	case float32:
		return big.NewInt(int64(v))
	case float64:
		return big.NewInt(int64(v))
	case int64:
		return big.NewInt(int64(v))
	case int32:
		return big.NewInt(int64(v))
	case int16:
		return big.NewInt(int64(v))
	case int8:
		return big.NewInt(int64(v))
	case int:
		return big.NewInt(int64(v))
	case uint64:
		return big.NewInt(int64(v))
	case uint32:
		return big.NewInt(int64(v))
	case uint16:
		return big.NewInt(int64(v))
	case uint8:
		return big.NewInt(int64(v))
	case uint:
		return big.NewInt(int64(v))
	case []uint8:
		i := new(big.Int)
		i.SetBytes(v)
		hx := ToHex(i, -1)
		n, err := hexutil.ToBN(hx, isLittleEndian)
		if err != nil {
			panic(err)
		}
		return n
	}

	return new(big.Int)
}

// ToHex creates a hex value from a math/big big number. 0 inputs returns a `0x` result, BN values return the actual value as a `0x` prefixed hex value. With `bitLength` set, it fixes the number to the specified length.
func ToHex(value *big.Int, bitLength int) string {
	return hexutil.HexFixLength(value.Text(16), bitLength, true)
}

// ToUint8Slice creates a uint8 array from a big number. Empty input returns an empty uint8 array result. Convert using little-endian format if `isLittleEndian` is set.
func ToUint8Slice(value *big.Int, bitLength int, isLittleEndian bool) []uint8 {
	bufLength := int(math.Ceil(float64(bitLength) / float64(8)))

	if value == nil {
		if bitLength == -1 {
			return []uint8{}
		}
		return make([]uint8, bufLength)
	}

	if bitLength == -1 {
		output, err := hexutil.ToUint8Slice(
			ToHex(value, bitLength),
			-1,
		)
		if err != nil {
			panic(err)
		}
		return output
	}

	output := make([]uint8, bufLength)
	b := value.Bytes()

	for index := 0; index < bufLength-1; index++ {
		output[index+1] = uint8(b[index])
	}

	if isLittleEndian {
		sort.Sort(sort.Reverse(Uint8Slice(output)))
	}

	return output
}

// Uint8Slice ...
type Uint8Slice []uint8

// Len ...
func (s Uint8Slice) Len() int {
	return len(s)
}

// Less ...
func (s Uint8Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap ...
func (s Uint8Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}