// Package baseconv Convert a number string between arbitrary bases
package baseconv

import (
	"bytes"
	"errors"
	"math/big"
)

// consts
const (
	BASE62Text = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// vars
var (
	base62          = []byte(BASE62Text)
	ErrInvalidBase  = errors.New("the base value must be between 2 and 62")
	ErrInvalidNumer = errors.New("the number string is invalid")
)

// Convert Convert a number between arbitrary bases
func Convert(num string, frombase int, tobase int) (string, error) {
	ret, err := ConvertBytes([]byte(num), frombase, tobase)
	return string(ret), err
}

// ConvertBytes ...
func ConvertBytes(num []byte, frombase int, tobase int) ([]byte, error) {
	if len(num) == 0 || frombase == tobase {
		return num, nil
	}

	if 2 > frombase || frombase > 62 || 2 > tobase || tobase > 62 {
		return nil, ErrInvalidBase
	}

	var fromdigits = base62[0:frombase]
	var todigits = base62[0:tobase]
	fromBi := big.NewInt(int64(frombase))
	toBi := big.NewInt(int64(tobase))

	x := big.NewInt(0)
	for _, digit := range num {
		x.Mul(x, fromBi)
		i := bytes.IndexByte(fromdigits, digit)
		if i < 0 {
			return nil, ErrInvalidNumer
		}
		x.Add(x, big.NewInt(int64(i)))
	}

	var res []byte
	for x.Cmp(big.NewInt(0)) > 0 {
		digit := new(big.Int).Mod(x, toBi).Uint64()
		res = append([]byte{todigits[digit]}, res...)
		x.Div(x, toBi)
	}

	return res, nil
}
