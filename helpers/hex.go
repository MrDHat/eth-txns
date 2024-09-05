package helpers

import (
	"fmt"

	"github.com/mrdhat/eth-txns/errors"
)

func ConvertHexToDecimal(hexVal string) (int, error) {
	// Remove the "0x" prefix if present
	if len(hexVal) >= 2 && hexVal[:2] == "0x" {
		hexVal = hexVal[2:]
	}

	decimal := 0
	for _, char := range hexVal {
		decimal *= 16
		if char >= '0' && char <= '9' {
			decimal += int(char - '0')
		} else if char >= 'a' && char <= 'f' {
			decimal += int(char - 'a' + 10)
		} else if char >= 'A' && char <= 'F' {
			decimal += int(char - 'A' + 10)
		} else {
			// Invalid hex character
			return 0, errors.ErrInvalidHexValue
		}
	}

	return decimal, nil
}

func ConvertPositiveDecimalToHex(decimal int) string {
	return fmt.Sprintf("0x%x", decimal)
}
