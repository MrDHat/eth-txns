package helpers_test

import (
	"fmt"
	"testing"

	"github.com/mrdhat/eth-txns/errors"
	"github.com/mrdhat/eth-txns/helpers"
)

func TestConvertHexToDecimal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		err      error
	}{
		{"Valid hex with 0x prefix", "0x1A", 26, nil},
		{"Valid hex without prefix", "1A", 26, nil},
		{"Valid hex lowercase", "0xab", 171, nil},
		{"Valid hex uppercase", "0xAB", 171, nil},
		{"Valid hex mixed case", "0xAb1", 2737, nil},
		{"Zero value", "0x0", 0, nil},
		{"Invalid hex character", "0xG1", 0, errors.ErrInvalidHexValue},
		{"Empty string", "", 0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := helpers.ConvertHexToDecimal(tt.input)
			if err != tt.err {
				t.Errorf("ConvertHexToDecimal(%q) error = %v, want %v", tt.input, err, tt.err)
				return
			}
			if result != tt.expected {
				t.Errorf("ConvertHexToDecimal(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertDecimalToHex(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{0, "0x0"},
		{10, "0xa"},
		{15, "0xf"},
		{16, "0x10"},
		{26, "0x1a"},
		{255, "0xff"},
		{1000, "0x3e8"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input_%d", tc.input), func(t *testing.T) {
			result := helpers.ConvertPositiveDecimalToHex(tc.input)
			if result != tc.expected {
				t.Errorf("ConvertDecimalToHex(%d) = %s; want %s", tc.input, result, tc.expected)
			}
		})
	}
}
