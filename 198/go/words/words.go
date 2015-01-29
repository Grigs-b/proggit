package words

import (
    "strings"
    "net/http"
    "bufio"
)

type WordList interface {
    IsValid(word string) bool
    ScoreWords(wordOne string, wordTwo string) int
    PossibleWords(letters []string) []string
    SetDictionary([]string)
}

type Words struct {
    dictionary map[string]bool
}

var Wordlist = Init()

func Init() *Words {
    // Pull the wordlist from a known location
    // TODO: Move magic string to configuration
    // TODO: Error handling if Init is not called before other functions
    w := new(Words)
    resp, _ := http.Get("http://www.joereynoldsaudio.com/enable1.txt")
    defer resp.Body.Close()

    words := make(map[string]bool)
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        token := scanner.Text()
        words[token]=true
    }
    w.dictionary = words
    return w
}


func (w *Words) SetDictionary(words []string) {
    for _, word := range words {
        w.dictionary[word]=true
    }
}


func (w Words) IsValid(word string) bool {
    if w.dictionary[word] {
        return true
    }
    return false
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


// TODO: Need to store words from list in a tree based structure that has at each node
// a letter and a bool signifying whether that node represents a complete word
// PossibleWords can then check for word validity using this


// Should use channel to prevent possibly exploding stack
func (w *Words) PossibleWords(letters []rune) []string {
    var result []string

    return result
}
