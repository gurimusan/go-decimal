package decimal

import (
	"testing"
)

var testDatas = map[string]string{
	"0":      "0",
	"1":      "1",
	"1.0":    "1.0",
	"1.00":   "1.00",
	"10":     "10",
	"1000":   "1000",
	"10.0":   "10.0",
	"10.1":   "10.1",
	"10.4":   "10.4",
	"10.5":   "10.5",
	"10.6":   "10.6",
	"10.9":   "10.9",
	"11.0":   "11.0",
	"1.234":  "1.234",
	"0.123":  "0.123",
	"0.012":  "0.012",
	"-0":     "-0",
	"-0.0":   "-0.0",
	"-00.00": "-0.00",
	"-1":     "-1",
	"-1.0":   "-1.0",
	"-0.1":   "-0.1",
	"-9.1":   "-9.1",
	"-9.11":  "-9.11",
	"-9.119": "-9.119",
	"-9.999": "-9.999",
}

func TestNewFromString(t *testing.T) {
	for _, s := range testDatas {
		d, err := NewFromString(s)
		if err != nil {
			t.Errorf("error while parsing %s", s)
		} else if d.String() != s {
			t.Errorf("expected %s, got %s (%s, %d)",
				s, d.String(),
				d.value.String(), d.expose)
		}
	}
}
