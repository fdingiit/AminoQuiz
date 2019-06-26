package AminoQuiz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertSimple(t *testing.T) {
	cases := []string{
		"172.168.5.1",
		"123.123.123.123",
		"127.0.0.1",
		"0.0.0.0",
		"255.255.255.255",
	}

	wanted := []uint32{
		172<<24 + 168<<16 + 5<<8 + 1,
		123<<24 + 123<<16 + 123<<8 + 123,
		127<<24 + 0<<16 + 0<<8 + 1,
		0,
		255<<24 + 255<<16 + 255<<8 + 255,
	}

	for i, c := range cases {
		got, err := IpConvert(c)
		fmt.Printf("case: %s, wanted: %d, got: %d\n", c, wanted[i], got)
		if !assert.Nil(t, err) || !assert.Equal(t, wanted[i], got) {
			t.FailNow()
		}
	}
}

func TestConvertSpacesValid(t *testing.T) {
	cases := []string{
		"172. 168 .5. 1",
		"123    . 123 .123.     123",
		"127 . 0 . 0 .                1",
		"0   .0 .    0.     0",
		"255 .255 .255. 255",
	}

	wanted := []uint32{
		172<<24 + 168<<16 + 5<<8 + 1,
		123<<24 + 123<<16 + 123<<8 + 123,
		127<<24 + 0<<16 + 0<<8 + 1,
		0,
		255<<24 + 255<<16 + 255<<8 + 255,
	}

	for i, c := range cases {
		got, err := IpConvert(c)
		fmt.Printf("case: %s, wanted: %d, got: %d\n", c, wanted[i], got)
		if !assert.Nil(t, err) || !assert.Equal(t, wanted[i], got) {
			t.FailNow()
		}
	}
}

func TestConvertSpacesInValid(t *testing.T) {
	cases := []string{
		"   172. 168. 5. 1",
		"123.123.123.     123       ",
		"1 72.168. 5. 1",
	}

	for _, c := range cases {
		got, err := IpConvert(c)
		fmt.Printf("case: [%s], wanted a err, got: [%d, %v]\n", c, got, err)
		if !assert.NotNil(t, err) {
			t.FailNow()
		}
	}
}

func TestConvertInputInValid(t *testing.T) {
	cases := []string{
		"123",
		"123.123.123.123.",
		"123.123.123",
		".1.2.3",

		"asbc",

		"1134.1.2.3",
		"",
		"-1.-2.-3.-4",
		"256.1.2.3",
	}

	for _, c := range cases {
		got, err := IpConvert(c)
		fmt.Printf("case: [%s], wanted a err, got: [%d, %v]\n", c, got, err)
		if !assert.NotNil(t, err) {
			t.FailNow()
		}
	}
}
