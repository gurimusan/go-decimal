package decimal

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

var decimalPattern = regexp.MustCompile(`(?i)(?P<sign>[-+])?((?P<int>\d*)(\.(?P<frac>\d*))?(E(?P<exp>[-+]?\d+))?)\z`)

// Decimal class
type Decimal struct {
	sign   int32
	value  *big.Int
	expose int32
}

// NewFromString returns a new Decimal from a string representation.
func NewFromString(str string) (Decimal, error) {
	match := decimalPattern.FindStringSubmatch(str)
	result := make(map[string]string)
	for i, name := range decimalPattern.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	var signStr string
	var intStr string
	var fracStr string
	var expStr string
	var ok bool
	var err error

	if v, ok := result["sign"]; ok {
		signStr = v
	}
	if v, ok := result["int"]; ok {
		intStr = v
	}
	if v, ok := result["frac"]; ok {
		fracStr = v
	}
	if v, ok := result["exp"]; ok {
		expStr = v
	}

	sign := 0
	decimalValue := new(big.Int)
	exp := int64(0)
	if signStr == "-" {
		sign = 1
	}
	if expStr != "" {
		exp, err = strconv.ParseInt(expStr, 10, 32)
		if err != nil {
			return Decimal{}, err
		}
	}
	if fracStr != "" {
		intStr = intStr + fracStr
		exp -= int64(len(fracStr))
	}
	_, ok = decimalValue.SetString(intStr, 10)
	if !ok {
		return Decimal{}, fmt.Errorf("can't convert %s to decimal", decimalValue)
	}
	if exp < math.MinInt32 || exp > math.MaxInt32 {
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", str)
	}

	return Decimal{
		sign:   int32(sign),
		value:  decimalValue,
		expose: int32(exp),
	}, nil
}

// Add decimal.
func (d Decimal) Add(other Decimal) (Decimal, error) {
	return Decimal{}, nil
}

func (d Decimal) String() string {
	valueStr := d.value.String()
	leftdigits := int(d.expose) + len(valueStr)

	var dotplace int
	if d.expose <= 0 && leftdigits > -6 {
		dotplace = leftdigits
	} else {
		dotplace = 1
	}

	var intStr, fracStr string
	if dotplace <= 0 {
		intStr = "0"
		fracStr = "." + strings.Repeat("0", int(-dotplace)) + valueStr
	} else if dotplace >= len(valueStr) {
		intStr = valueStr + strings.Repeat("0", dotplace-len(valueStr))
		fracStr = ""
	} else {
		intStr = valueStr[:dotplace]
		fracStr = "." + valueStr[dotplace:]
	}

	var expStr string
	if leftdigits == dotplace {
		expStr = ""
	} else {
		expStr = "E" + fmt.Sprintf("%+d", leftdigits-dotplace)
	}

	if d.sign == 1 {
		return "-" + intStr + fracStr + expStr
	}
	return intStr + fracStr + expStr
}
