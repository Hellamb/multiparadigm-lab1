//Топ часто используемых слов в тексте без циклов и функций(и частично встроенных функций)
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	const MAX_DIF_WORDS_IN_TEXT = 10000
	const MAX_STOPWORDS = 10
	const DISPLAYED_WORDS_AMOUNT = 25
	const STOPWORDS_FILE_PATH = "stopwords.txt"
	const TEXT_FILE_PATH = "myText.txt"

	var stopwords [MAX_STOPWORDS]string
	var words [MAX_DIF_WORDS_IN_TEXT]string
	var wordsAmount [MAX_DIF_WORDS_IN_TEXT]int
	difWordsAmount := 0
	stopwordsAmount := 0

	lowChars := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	upperChars := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	// Открытие файла со стоп-словами
	file, err := os.Open(STOPWORDS_FILE_PATH)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// Перенос стоп-слов в массив
NextLine1:
	if scanner.Scan() {
		line := scanner.Text()
		charIndex := 0
		word := ""
	NextChar1:
		ch := string(line[charIndex])
		//Когда символ не буква
		if ch == " " || ch == "," || ch == "." {
			stopwords[stopwordsAmount] = word
			stopwordsAmount++
			word = ""
			charIndex++
			goto NextChar1
		}

		word += ch
		charIndex++
		//Когда нет символов в строке
		if charIndex >= len(line) {
			stopwords[stopwordsAmount] = word
			stopwordsAmount++
			word = ""
			goto NextLine1
		}
		//Когда символ обработан
		goto NextChar1
	}
	file, err = os.Open(TEXT_FILE_PATH)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner = bufio.NewScanner(file)

NextLine2:
	if scanner.Scan() {
		line := scanner.Text()
		newLine := false
		charIndex := 0
		word := ""
	NextChar2:
		if charIndex >= len(line) {
			goto NextLine2
		}
		ch := string(line[charIndex])
		abcNum := 0
		//Когда символ не буква
		if ch == " " || ch == "," || ch == "." {
			charIndex++
			goto WordProcessing
		}
		//Проверка на большую букву
	UpperCheck:
		if ch == upperChars[abcNum] {
			ch = lowChars[abcNum]
		} else if abcNum < 25 {
			abcNum++
			goto UpperCheck
		}

		word += ch
		charIndex++
		//Когда нет символов в строке
		if charIndex >= len(line) {
			newLine = true
			goto WordProcessing
		}

		goto NextChar2
		//Добавление слова в массив или увеличение значения его появления
	WordProcessing:
		wordNum := 0
	WordProcessingCycle:
		if words[wordNum] == "" {
			difWordsAmount++
			words[wordNum] = word
			wordsAmount[wordNum]++
			word = ""
			if newLine {
				goto NextLine2
			}
			goto NextChar2
		} else if words[wordNum] == word {
			wordsAmount[wordNum]++
			word = ""
			goto NextChar2
		} else if word == "" {
			if newLine {
				goto NextLine2
			}
			goto NextChar2
		}
		wordNum++
		goto WordProcessingCycle
	}

	//Сортировка пузырьком
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
	if wordsAmount[j] < wordsAmount[j+1] {
		wordsAmount[j], wordsAmount[j+1] = wordsAmount[j+1], wordsAmount[j]
		words[j], words[j+1] = words[j+1], words[j]
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
	checkedNum := 0
TopDisplay:
	checkedNum = 0
CheckCycle:
	if words[displayedIndex] == stopwords[checkedNum] {
		displayedIndex++
		goto TopDisplay
	}
	checkedNum++
	if checkedNum < stopwordsAmount {
		goto CheckCycle
	}
	displayedNum++
	fmt.Println(displayedNum, ")", words[displayedIndex], "-", wordsAmount[displayedIndex])
	displayedIndex++

	if displayedNum < DISPLAYED_WORDS_AMOUNT && displayedIndex < difWordsAmount {
		goto TopDisplay
	}

}
