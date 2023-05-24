package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input file> <output file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	text := string(inputData)
	formattedText := formatText(text)

	err = ioutil.WriteFile(outputFile, []byte(formattedText), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("Text conversion successful")
}

func formatText(text string) string {
    // Define regex patterns
    hexPattern := regexp.MustCompile(`\b([0-9A-Fa-f]+)\s+\(hex\)`)
    binPattern := regexp.MustCompile(`\b([01]+)\s+\(bin\)`)
    upPattern := regexp.MustCompile(`\b(\w+)\s+\(up\)`)
    lowPattern := regexp.MustCompile(`\b(\w+)\s+\(low\)`)
    capPattern := regexp.MustCompile(`\b(\w+)\s+\(cap\)`)
    upNumPattern := regexp.MustCompile(`\b((?:\w+\s+){0,}\w+)\s+\(up,\s*(\d+)\)`)
    lowNumPattern := regexp.MustCompile(`\b((?:\w+\s+){0,}\w+)\s+\(low,\s*(\d+)\)`)
    capNumPattern := regexp.MustCompile(`\b((?:\w+\s+){0,}\w+)\s+\(cap,\s*(<IPAddress>)`)
    puncPattern := regexp.MustCompile(`\s+([.,!?:;])`)
    puncGroupPattern := regexp.MustCompile(`([.,!?:;])\s+([.,!?:;])`)
    quotePattern := regexp.MustCompile(`'\s+([^']{1,})\s+'`)
    aAnPattern := regexp.MustCompile(`\ba\s+([aeiouhAEIOUH])`)

    // Replace hex and bin
    text = hexPattern.ReplaceAllStringFunc(text, func(s string) string {
        hexNum := hexPattern.FindStringSubmatch(s)[1]
        num, _ := strconv.ParseInt(hexNum, 16, 64)
        return fmt.Sprintf("%d", num)
    })
    text = binPattern.ReplaceAllStringFunc(text, func(s string) string {
        binNum := binPattern.FindStringSubmatch(s)[1]
        num, _ := strconv.ParseInt(binNum, 2, 64)
        return fmt.Sprintf("%d", num)
    })

    // Replace up, low and cap
    text = upPattern.ReplaceAllStringFunc(text, func(s string) string {
        word := upPattern.FindStringSubmatch(s)[1]
        return strings.ToUpper(word)
    })
    text = lowPattern.ReplaceAllStringFunc(text, func(s string) string {
        word := lowPattern.FindStringSubmatch(s)[1]
        return strings.ToLower(word)
    })
    text = capPattern.ReplaceAllStringFunc(text, func(s string) string {
        word := capPattern.FindStringSubmatch(s)[1]
        return strings.Title(word)
    })

    // Replace upNum, lowNum and capNum
    text = upNumPattern.ReplaceAllStringFunc(text, func(s string) string {
        match := upNumPattern.FindStringSubmatch(s)
        numWordsStr := match[1]
        numWordsSlice := strings.Fields(numWordsStr)
        numWordsToChangeStr := match[2]
        numWordsToChangeInt, _ := strconv.Atoi(numWordsToChangeStr)

        if numWordsToChangeInt > len(numWordsSlice) {
            numWordsToChangeInt = len(numWordsSlice)
        }

        for i:=len(numWordsSlice)-numWordsToChangeInt; i<len(numWordsSlice); i++ {
            numWordsSlice[i] = strings.ToUpper(numWordsSlice[i])
        }

        return strings.Join(numWordsSlice," ")
        
    })
    text = lowNumPattern.ReplaceAllStringFunc(text, func(s string) string {
        match := lowNumPattern.FindStringSubmatch(s)
        numWordsStr := match[1]
        numWordsSlice := strings.Fields(numWordsStr)
        numWordsToChangeStr := match[2]
        numWordsToChangeInt, _ := strconv.Atoi(numWordsToChangeStr)

        if numWordsToChangeInt > len(numWordsSlice) {
            numWordsToChangeInt = len(numWordsSlice)
        }

        for i:=len(numWordsSlice)-numWordsToChangeInt; i<len(numWordsSlice); i++ {
            numWordsSlice[i] = strings.ToLower(numWordsSlice[i])
        }

        return strings.Join(numWordsSlice," ")
        
    })
    text = capNumPattern.ReplaceAllStringFunc(text, func(s string) string {
        match := capNumPattern.FindStringSubmatch(s)
        numWordsStr := match[1]
        numWordsSlice := strings.Fields(numWordsStr)
        numWordsToChangeStr := match[2]
        numWordsToChangeInt, _ := strconv.Atoi(numWordsToChangeStr)

        if numWordsToChangeInt > len(numWordsSlice) {
            numWordsToChangeInt = len(numWordsSlice)
        }

        for i:=len(numWordsSlice)-numWordsToChangeInt; i<len(numWordsSlice); i++ {
            numWordsSlice[i] = strings.Title(numWordsSlice[i])
        }

        return strings.Join(numWordsSlice," ")
        
    })

    // Replace punctuation
    text = puncPattern.ReplaceAllString(text,"$1")
    text = puncGroupPattern.ReplaceAllString(text,"$1$2")
    text = quotePattern.ReplaceAllString(text,"'$1'")

    // Replace a/an
	text = aAnPattern.ReplaceAllStringFunc(text, func(s string) string {
		return "an " + aAnPattern.FindStringSubmatch(s)[1]
	})

	return text
}
