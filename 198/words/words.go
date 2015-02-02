package words

import (
    "io"
    "os"
    "strings"
    "bufio"
)

type WordList interface {
    IsValid(word string) bool
    PossibleWords(letters []rune) []string
    AddWord(string)
}

type Wordset struct {
    dictionary map[string]bool
}

func NewWordset() *Wordset {
    return &Wordset{dictionary: make(map[string]bool)}
}

func (w *Wordset) LoadWordsFromFile(path string) {

    // Changed to reading local file because remote host was going very slowly
    // open input file
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        line = strings.TrimSpace(line)
        line = strings.ToLower(line)
        if err != nil && err != io.EOF {
            panic(err)
        }
        w.AddWord(line)
        if err == io.EOF {
            break
        }
    }
}


func (w *Wordset) AddWord(word string) {
    w.dictionary[word]=true
}


func (w Wordset) IsValid(word string) bool {
    if w.dictionary[word] {
        return true
    }
    return false
}


func (w *Wordset) PossibleWords(letters []rune) []string {
    var result []string
    var tmp string
    for word, _ := range w.dictionary {
        tmp = word
        for _, letter := range letters {
            tmp = strings.Replace(tmp, string(letter), "", 1)
            if len(tmp) == 0 {
                result = append(result, word)
                break
            }
        }
    }
    return result
}


func ScoreWords(wordOne string, wordTwo string) int {
    // returns positive integer if word one wins, negative integer if word two wins
    // TODO: Use score struct to save letters remaining as well as each word score
    for _, arune := range wordOne + wordTwo {
        if indexOne := strings.IndexRune(wordOne, arune); indexOne > -1 {
            if indexTwo := strings.IndexRune(wordTwo, arune); indexTwo > -1 {
                wordOne = strings.Replace(wordOne, string(arune), "", 1)
                wordTwo = strings.Replace(wordTwo, string(arune), "", 1)
            }
        }
    }
    return len(wordOne) - len(wordTwo)
}
