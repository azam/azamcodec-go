package azamcodec_test

import (
	"azamcodec"
	"testing"
)

var REFERENCE_SAMPLES = map[string]uint{
	"0":                0x0,
	"1":                0x1,
	"f":                0xf,
	"h0":               0x10,
	"h1":               0x11,
	"hf":               0x1f,
	"z0":               0xf0,
	"zf":               0xff,
	"hg0":              0x100,
	"hg1":              0x101,
	"hh1":              0x111,
	"hzf":              0x1ff,
	"zzzf":             0xffff,
	"zzzzf":            0xfffff,
	"zzzzzf":           0xffffff,
	"zzzzzzf":          0xfffffff,
	"zzzzzzzf":         0xffffffff,
	"zzzzzzzzf":        0xfffffffff,
	"zzzzzzzzzf":       0xffffffffff,
	"zzzzzzzzzzf":      0xfffffffffff,
	"zzzzzzzzzzzf":     0xffffffffffff,
	"zzzzzzzzzzzzf":    0xfffffffffffff,
	"hgggggggggggg0":   0x10000000000000,
	"zzzzzzzzzzzzzf":   0xffffffffffffff,
	"zzzzzzzzzzzzzzf":  0xfffffffffffffff,
	"zzzzzzzzzzzzzzzf": 0xffffffffffffffff,
}

func TestEncodeUint(t *testing.T) {
	for expected, input := range REFERENCE_SAMPLES {
		actual := azamcodec.EncodeUint(input)
		if actual != expected {
			t.Errorf("EncodeUint(%0x) = %s; want %s", input, actual, expected)
		}
	}
}

func TestEncodeUintsEmpty(t *testing.T) {
	const expected = ""
	actual := azamcodec.EncodeUints()
	if actual != expected {
		t.Errorf("EncodeUints() = %s; want %s", actual, expected)
	}
	actual2 := azamcodec.EncodeUints(make([]uint, 0)...)
	if actual2 != expected {
		t.Errorf("EncodeUints(make([]uint, 0)...) = %s; want %s", actual2, expected)
	}
}
