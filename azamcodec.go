package azamcodec

var LOWER_ALPHABETS = []byte("0123456789abcdef")
var HIGHER_ALPHABETS = []byte("ghjkmnpqrstvwxyz")

// Given unsigned integer (`uint`) arguments, return Azam Codec encoded string.
//
// Usage:
//
//	azamcodec.EncodeUints(0x1, 0x2, 0x3) // "123"
//	azamcodec.EncodeUints(0xdeadbeef, 0x15, 0xc001) // "xytxvyyfh5wgg1"
func EncodeUints(nums ...uint) string {
	var value string = ""
	for _, num := range nums {
		value += EncodeUint(num)
	}
	return value
}

// Given an unsigned integer(`uint`), return Azam Codec encoded string.
//
// Usage:
//
//	azamcodec.EncodeUint(0x1) // "1"
//	azamcodec.EncodeUint(0xdeadbeef) // "xytxvyyf"
func EncodeUint(num uint) string {
	var value []byte = make([]byte, 0)
	var lowNybbleWritten = false
	for {
		var nybbleValue = num & 0xf
		var nybbleChar byte
		if lowNybbleWritten {
			nybbleChar = HIGHER_ALPHABETS[nybbleValue]
		} else {
			nybbleChar = LOWER_ALPHABETS[nybbleValue]
		}
		if !lowNybbleWritten {
			lowNybbleWritten = true
		}
		value = append(value, nybbleChar)
		num = num >> 4
		if num == 0 {
			break
		}
	}
	return string(reverseSlice(value))
}

func reverseSlice[T any](original []T) (reversed []T) {
	reversed = make([]T, len(original))
	copy(reversed, original)

	for i := len(reversed)/2 - 1; i >= 0; i-- {
		tmp := len(reversed) - 1 - i
		reversed[i], reversed[tmp] = reversed[tmp], reversed[i]
	}

	return
}

func Decode(input string) []uint {
	var ret []uint
	return ret
}
