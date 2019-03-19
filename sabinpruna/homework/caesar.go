package homework

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//GenerateRandomNumberFile returns a file with 100000 random numbers used for caesar cypher
func GenerateRandomNumberFile() *os.File {
	var numbers []int

	for index := 0; index < 100000; index++ {
		numbers = append(numbers, rand.Intn(27)) //smaller , just go through the alphabet once at most
	}

	file, err := os.Create("randomNumbers.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	stringNumbers := []string{}
	for _, number := range numbers {
		stringNumbers = append(stringNumbers, strconv.Itoa(number))
	}
	result := strings.Join(stringNumbers, " ")

	_, err = file.WriteString(result)
	if err != nil {
		fmt.Println(err)
	}

	return file

}

/**
 * This can be done a lot nicer by giving a func as param for caesarEncrypt/Decrypt
 * and an operation for the privates, so function calls are somewhat shorter (maybe there was something even better)
 *  if i remembered how to do it properly
 **/

//Encrypt returns the encrypted message using the Caesar cypher with given key (numbers)
func Encrypt(numbers []int, message string) string {
	messageRunes := []rune(message)
	counter := 0

	var encryptedRunes []rune

	for _, messageRune := range messageRunes {
		if counter >= 100000 {
			counter = 0
		}
		encryptedRunes = append(encryptedRunes, caesarEncrypt(messageRune, counter))
		counter++
	}
	return string(encryptedRunes)
}

//Decrypt returns the message decrypted using the Caesar cyper with given key (numbers)
func Decrypt(numbers []int, message string) string {
	messageRunes := []rune(message)
	counter := 0

	var encryptedRunes []rune

	for _, messageRune := range messageRunes {
		if counter >= 100000 {
			counter = 0
		}
		encryptedRunes = append(encryptedRunes, caesarDecrypt(messageRune, counter))
		counter++
	}
	return string(encryptedRunes)
}

//----------------------------------PRIVATES------------------------------------------

func caesarEncrypt(r rune, shift int) rune {
	// Shift character by specified number of places.
	// If beyond range, shift backward or forward.
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

func caesarDecrypt(r rune, shift int) rune {
	// Shift character by specified number of places.
	// If beyond range, shift backward or forward.
	s := int(r) - shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

//------------------------------------------------------------------------------------
