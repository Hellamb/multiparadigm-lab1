package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	const LINES_IN_PAGE = 45
	const MAX_WORD_FREQUENCY = 100
	const MAX_DIF_WORDS_IN_TEXT = 10000
	const MAX_PAGES = 10000
	const TEXT_FILE_PATH = "myText.txt"

	var words [MAX_DIF_WORDS_IN_TEXT]string
	var wordsAmount [MAX_DIF_WORDS_IN_TEXT]int
	var wordsPages [MAX_DIF_WORDS_IN_TEXT][MAX_PAGES]int

	difWordsAmount := 0

	lowChars := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	upperChars := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	file, err := os.Open(TEXT_FILE_PATH)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	pageNum := 0
	lineNum := 0

NextPage:
	pageNum++
	lineNum = 0

NextLine:
	lineNum++
	line := scanner.Text()
	charIndex := 0
	word := ""
	newLine := false

	if scanner.Scan() {
		if lineNum > 45 {
			goto NextPage
		}

	NextChar:
		if charIndex >= len(line) {
			goto NextLine
		}
		ch := string(line[charIndex])
		abcNum := 0

		if ch == " " || ch == "." || ch == "," {
			charIndex++
			goto WordProcessing

		}

	UpperCheck:
		if ch == upperChars[abcNum] {
			ch = lowChars[abcNum]
		} else if abcNum < 25 {
			abcNum++
			goto UpperCheck
		}

		word += ch
		charIndex++
		if charIndex >= len(line) {
			newLine = true
			goto WordProcessing
		}

		goto NextChar

	WordProcessing:
		wordNum := 0
	WordProcessingCycle:
		if words[wordNum] == "" {
			difWordsAmount++
			words[wordNum] = word
			wordsAmount[wordNum]++
			wordsPages[wordNum][0]++
			wordsPages[wordNum][1] = pageNum
			word = ""
			if newLine {
				goto NextLine
			}
			goto NextChar
		} else if words[wordNum] == word {
			wordsAmount[wordNum]++
			if wordsPages[wordNum][wordsPages[wordNum][0]] != pageNum {
				wordsPages[wordNum][0]++
				wordsPages[wordNum][wordsPages[wordNum][0]] = pageNum
			}
			word = ""
			goto NextChar
		} else if word == "" {
			if newLine {
				goto NextLine
			}
			goto NextChar
		}
		wordNum++
		goto WordProcessingCycle

	}
	i, j := 0, 0
	if difWordsAmount <= 1 {
		goto Sorted
	}
Loop1:
	j = 0
	if i+1 >= difWordsAmount {
		goto Sorted
	}
Loop2:
	if words[j] > words[j+1] {
		wordsAmount[j], wordsAmount[j+1] = wordsAmount[j+1], wordsAmount[j]
		words[j], words[j+1] = words[j+1], words[j]
		wordsPages[j], wordsPages[j+1] = wordsPages[j+1], wordsPages[j]
	}
	if j+i+2 >= difWordsAmount {
		i++
		goto Loop1
	}
	j++
	goto Loop2
Sorted:

	displayedNum := 0
	displayedIndex := 0

DisplayCycle:

	if wordsAmount[displayedIndex] > 100 {
		displayedIndex++
		goto DisplayCycle
	}
	displayedNum++
	printed := 0
	fmt.Print(displayedNum, ") ", words[displayedIndex], " - ")
PrintCycle:
	printed++
	fmt.Print(wordsPages[displayedIndex][printed], " ")
	if printed < wordsPages[displayedIndex][0] {
		goto PrintCycle
	}
	fmt.Print("\n")
	displayedIndex++
	if displayedIndex < difWordsAmount {
		goto DisplayCycle
	}

}
