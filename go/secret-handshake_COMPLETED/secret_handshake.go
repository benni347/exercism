package secret

import "strconv"

// Handshake returns the secret handshake
func Handshake(code uint) []string {
	binaryRep := strconv.FormatUint(uint64(code), 2)
	var result []string
	for i := len(binaryRep) - 1; i >= 0; i-- {
		if binaryRep[i] == '1' {
			if i == len(binaryRep)-1 {
				result = append(result, "wink")
			} else if i == len(binaryRep)-2 {
				result = append(result, "double blink")
			} else if i == len(binaryRep)-3 {
				result = append(result, "close your eyes")
			} else if i == len(binaryRep)-4 {
				result = append(result, "jump")
			}
		}
	}

	// Check if the fifth bit is set and reverse the slice
	if len(binaryRep) >= 5 && binaryRep[0] == '1' {
		reverseSlice(result)
	}
	return result
}

// reverseSlice reverses a slice of strings in place
func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
