package go_common

import "fmt"

func getHint(secret, guess string) string {
	A := 0
	B := 0

	guessBytes := []byte(guess)
	secretBytes := []byte(secret)
	guessedByteMap := make(map[int]int)
	secretByteMap := make(map[int]int)

	for guestIndex, guestValue := range guessBytes {
		if guestValue == secretBytes[guestIndex] {
			setGuessAndSecretIndexPair(guessedByteMap, secretByteMap, guestIndex, guestIndex)
			A = increaseBullNumber(A)
		}
	}

	for guestIndex, guestValue := range guessBytes {
		for secretIndex, secretValue := range secretBytes {
			if guestIndex != secretIndex && guestValue == secretValue &&
				checkSecretAndGuessIndexUsed(guessedByteMap, secretByteMap, guestIndex, secretIndex) {
				B = increaseCowNumber(B)
				setGuessAndSecretIndexPair(guessedByteMap, secretByteMap, guestIndex, secretIndex)

			}
		}
	}

	return fmt.Sprintf("%dA%dB", A, B)
}

func increaseBullNumber(a int) int {
	return a + 1
}

func increaseCowNumber(b int) int {
	return b + 1
}

func setGuessAndSecretIndexPair(guessedByteMap map[int]int, secretByteMap map[int]int, guestIndex, secretIndex int) {
	guessedByteMap[guestIndex] = 1
	secretByteMap[secretIndex] = 1
}

func checkSecretAndGuessIndexUsed(guessedByteMap, secretByteMap map[int]int, guestIndex, secretIndex int) bool {
	return secretByteMap[secretIndex] == 0 && guessedByteMap[guestIndex] == 0
}
