package homework

var matrix = [9][9]rune{
	{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'},
	{'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'},
	{'q', 'r', 's', 't', 'u', 'v', 'w', 'x'},
	{'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F'},
	{'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N'},
	{'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V'},
	{'W', 'X', 'Y', 'Z', '0', '1', '2', '3'},
	{'4', '5', '6', '7', '8', '9', ' ', '.'}}

//MatrixCipher encodes and decodes a message based on the matrix described above,
//interchanging lettters 2 by 2 finding their intersections in the matrix
func MatrixCipher(message string) string {
	messageRunes := []rune(message)

	encryptedRunes := []rune{}

	for index := 0; index < len(messageRunes)-2; index += 2 {
		//x, y := messageRunes[index : index+1]
		first := messageRunes[index]
		second := messageRunes[index+1]
		var firstI, firstJ, secondI, secondJ int

		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if first == matrix[i][j] {
					firstI = i
					firstJ = j
				} else if second == matrix[i][j] {
					secondI = i
					secondJ = j
				}
			}
		}

		//interchange , only happens on same column , same row is treated by default with interchange
		if firstJ == secondJ {
			encryptedRunes = append(encryptedRunes, second, first)
		} else {
			encryptedRunes = append(encryptedRunes, matrix[firstI][secondJ], matrix[secondI][firstJ])
		}

	}

	if len(messageRunes)%2 == 1 {
		encryptedRunes = append(encryptedRunes, messageRunes[len(messageRunes)-1])
	}

	return string(encryptedRunes)
}
