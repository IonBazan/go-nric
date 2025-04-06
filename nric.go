package nric

import (
	"errors"
	"fmt"
	"regexp"
)

// Custom error types
var (
	ErrInvalidFormat   = errors.New("invalid NRIC format")
	ErrInvalidChecksum = errors.New("invalid NRIC checksum")
)

type NRIC struct {
	id string
}

const (
	length                = 9
	checksumCitizen       = "JZIHGFEDCBA"
	checksumForeigner     = "XWUTRQPNMLK"
	checksumForeigner2022 = "XWUTRQPNJLK"
	prefixCitizen1900     = "S"
	prefixCitizen2000     = "T"
	prefixForeigner1900   = "F"
	prefixForeigner2000   = "G"
	prefixForeigner2022   = "M"
)

func NewNRIC(id string) (*NRIC, error) {
	n := &NRIC{id: id}
	if err := n.validate(); err != nil {
		return nil, err
	}
	return n, nil
}

func (n *NRIC) String() string {
	return n.id
}

func (n *NRIC) IsForeigner() bool {
	prefix := n.id[0]
	return prefix == prefixForeigner1900[0] ||
		prefix == prefixForeigner2000[0] ||
		prefix == prefixForeigner2022[0]
}

func (n *NRIC) Is2000() bool {
	prefix := n.id[0]
	return prefix == prefixCitizen2000[0] || prefix == prefixForeigner2000[0]
}

// IsSeriesM checks if the NRIC is from the M series (post-2022 foreigners)
func (n *NRIC) IsSeriesM() bool {
	return n.id[0] == prefixForeigner2022[0]
}

func (n *NRIC) validate() error {
	regexPattern := fmt.Sprintf(`^([%s%s][\d]{7}[%s])|([%s%s%s][\d]{7}[%s%s])$`,
		prefixCitizen1900,
		prefixCitizen2000,
		checksumCitizen,
		prefixForeigner1900,
		prefixForeigner2000,
		prefixForeigner2022,
		checksumForeigner,
		checksumForeigner2022,
	)

	regex := regexp.MustCompile(regexPattern)

	if !regex.MatchString(n.id) {
		return fmt.Errorf("%w: %s", ErrInvalidFormat, n.id)
	}

	if n.generateChecksum() != n.id[length-1] {
		return fmt.Errorf("%w: %s", ErrInvalidChecksum, n.id)
	}

	return nil
}

func (n *NRIC) generateChecksum() byte {
	checksum := n.getOffset()
	checksumWeights := []int{2, 7, 6, 5, 4, 3, 2}

	for i, weight := range checksumWeights {
		digit := int(n.id[i+1] - '0')
		checksum += digit * weight
	}

	checksumChars := n.getChecksumChars()
	return checksumChars[checksum%11]
}

func (n *NRIC) getOffset() int {
	if n.IsSeriesM() {
		return 3
	}
	if n.Is2000() {
		return 4
	}
	return 0
}

func (n *NRIC) getChecksumChars() string {
	if n.IsSeriesM() {
		return checksumForeigner2022
	}
	if n.IsForeigner() {
		return checksumForeigner
	}
	return checksumCitizen
}
