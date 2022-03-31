package util

import (
	"fmt"
	"math/big"
	"strings"
)

// ARToWinston 1 AR = 1000000000000
func ARToWinston(arAmount string) (string, error) {
	var integerPart = ""
	var fractionalPart = ""

	numberOfDot := strings.Count(arAmount, ".")
	if numberOfDot > 1 {
		return "", fmt.Errorf("invalid input %v", arAmount)
	} else if numberOfDot == 1 {
		fields := strings.Split(arAmount, ".")
		integerPart = fields[0]
		fractionalPart = fields[1]
	} else { // In case of numberOfDot == 0
		integerPart = arAmount
	}

	winstonAmount := new(big.Int)
	winstonAmount, ok := winstonAmount.SetString(integerPart+"000000000000", 10)
	if !ok {
		return "", fmt.Errorf("invalid amount %v", arAmount)
	}

	if len(fractionalPart) > 12 {
		return "", fmt.Errorf("invalid input, too many fractional digits %v", arAmount)
	}

	fractionalPart = fractionalPart + strings.Repeat("0", 12-len(fractionalPart))

	fractionalAmount := new(big.Int)
	fractionalAmount, ok = fractionalAmount.SetString(fractionalPart, 10)
	if !ok {
		return "", fmt.Errorf("invalid amount %v", arAmount)
	}

	// Run here, assume arAmount = 1.2
	//   winstonAmount would be 1000000000000
	// fractionalAmount would be 200000000000
	return winstonAmount.Add(winstonAmount, fractionalAmount).String(), nil
}
