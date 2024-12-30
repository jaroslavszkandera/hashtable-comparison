package testdata

import (
	"math/rand"
	"strings"
)

func generateRandomString(minLength, maxLength int) string {
	if minLength > maxLength {
		panic("Min length must be less than or equal to max length")
	}
	length := rand.Intn(maxLength-minLength+1) + minLength
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var randomString strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		randomString.WriteByte(charset[randomIndex])
	}
	return randomString.String()
}

func GenerateTestData(numValues int) []string {
	testData := make([]string, numValues)
	for i := 0; i < numValues; i++ {
		testData[i] = generateRandomString(10, 100)
	}
	return testData
}
