package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ConvertHexToDecimal(hex string) string {
	decimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		// Handle conversion error
		return hex
	}
	return strconv.FormatInt(decimal, 10)
}

func ConvertBinaryToDecimal(binary string) string {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		// Handle conversion error
		return binary
	}
	return strconv.FormatInt(decimal, 10)
}

func ConvertToUppercase(word string) string {
	// Convert the word to uppercase
	return strings.ToUpper(word)
}

func ConvertToLowercase(word string) string {
	// Convert the word to lowercase
	return strings.ToLower(word)
}

func CapitalizeWord(word string) string {
	// Capitalize the first letter of the word
	if len(word) > 0 {
		return strings.ToUpper(string(word[0])) + word[1:]
	}
	return word
}

func ModifyWordsCase(text string, caseType string, wordCount int) string {
	words := strings.Fields(text)
	if wordCount > len(words) {
		wordCount = len(words)
	}

	switch caseType {
	case "up":
		for i := 0; i < wordCount; i++ {
			words[i] = ConvertToUppercase(words[i])
		}
	case "low":
		for i := 0; i < wordCount; i++ {
			words[i] = ConvertToLowercase(words[i])
		}
	case "cap":
		for i := 0; i < wordCount; i++ {
			words[i] = CapitalizeWord(words[i])
		}
	}

	return strings.Join(words, " ")
}

func main() {
	// Check if the correct number of command-line arguments is provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go input_file output_file")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the input file
	inputText, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading the input file: %v\n", err)
		os.Exit(1)
	}

	// Convert input text to string
	text := string(inputText)

	// Define regular expressions for different modifications
	hexRegex := regexp.MustCompile(`\((hex)\)`)
	binRegex := regexp.MustCompile(`\((bin)\)`)
	upRegex := regexp.MustCompile(`\((up)\)`)
	lowRegex := regexp.MustCompile(`\((low)\)`)
	capRegex := regexp.MustCompile(`\((cap)\)`)
	modifyRegex := regexp.MustCompile(`\((up|low|cap),\s*(\d+)\)`)

	// Parse and apply modifications
	modifiedText := text

	// Replace (hex) occurrences with decimal versions
	modifiedText = hexRegex.ReplaceAllStringFunc(modifiedText, func(match string) string {
		word := strings.TrimSpace(strings.TrimSuffix(match, "(hex)"))
		decimal := ConvertHexToDecimal(word)
		return decimal
	})

	// Replace (bin) occurrences with decimal versions
	modifiedText = binRegex.ReplaceAllStringFunc(modifiedText, func(match string) string {
		word := strings.TrimSpace(strings.TrimSuffix(match, "(bin)"))
		decimal := ConvertBinaryToDecimal(word)
		return decimal
	})

	// Replace (up) occurrences with uppercase words
	modifiedText = upRegex.ReplaceAllStringFunc(modifiedText, func(match string) string {
		word := strings.TrimSpace(strings.TrimSuffix(match, "(up)"))
		return ConvertToUppercase(word)
	})

	// Replace (low) occurrences with lowercase words
	modifiedText = lowRegex.ReplaceAllStringFunc(modifiedText, func(match string) string {
		word := strings.TrimSpace(strings.TrimSuffix(match, "(low)"))
		return ConvertToLowercase(word)
	})

	// Replace (cap) occurrences with capitalized words
	modifiedText = capRegex.ReplaceAllStringFunc(modifiedText, func(match string) string {
		word := strings.TrimSpace(strings.TrimSuffix(match, "(cap)"))
		return CapitalizeWord(word)
	})

	// Replace (up, n), (low, n), (cap, n) occurrences with modified words based on word count
	modifiedText = modifyRegex.ReplaceAllStringFunc(modifiedText, func(match string) string {
		groups := modifyRegex.FindStringSubmatch(match)
		caseType := groups[1]
		wordCountStr := groups[2]
		wordCount, _ := strconv.Atoi(wordCountStr)

		word := strings.TrimSpace(strings.TrimSuffix(match, groups[0]))
		return ModifyWordsCase(word, caseType, wordCount)
	})

	// Format punctuation marks
	modifiedText = strings.ReplaceAll(modifiedText, " ,", ",")
	modifiedText = strings.ReplaceAll(modifiedText, " .", ".")
	modifiedText = strings.ReplaceAll(modifiedText, " !", "!")
	modifiedText = strings.ReplaceAll(modifiedText, " ?", "?")
	modifiedText = strings.ReplaceAll(modifiedText, " :", ":")
	modifiedText = strings.ReplaceAll(modifiedText, " ;", ";")

	// Format ellipsis and question marks
	modifiedText = strings.ReplaceAll(modifiedText, "...", "...")
	modifiedText = strings.ReplaceAll(modifiedText, "!?", "!?")

	// Format words enclosed in single quotes
	modifiedText = strings.ReplaceAll(modifiedText, " '", "'")
	modifiedText = strings.ReplaceAll(modifiedText, "' ", "'")

	// Replace "a" with "an" when followed by a vowel or 'h'
	modifiedText = strings.ReplaceAll(modifiedText, " a ", " an ")

	// Write the modified text to the output file
	err = ioutil.WriteFile(outputFile, []byte(modifiedText), 0644)
	if err != nil {
		fmt.Printf("Error writing to the output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Modification complete!")
}
