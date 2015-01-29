package words

import (
    "fmt"
    "strings"
    "net/http"
    "bufio"
)

type WordList interface {
    IsValid(word string) bool
    ScoreWords(wordOne string, wordTwo string) int
    PossibleWords(letters []string) []string
    //SetDictionary([]string)
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

/*
func (w Words) SetDictionary(words []string) {

}
*/

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



func (w *Words) PossibleWords(letters []rune) []string {
    var words, result []string
    for pos, letter := range letters {
        var temp = make([]rune, len(letters))
        var start = make([]rune, 1)
        copy(temp, letters)
        start[0] = letter
        remainder := append(temp[:pos], temp[pos+1:]...)
        result = append(result, w.findwords(start, remainder, words)...)
    }
    return result
}

//TODO: Refactor to use channel for memory safety, also get working
func (w *Words) findwords(check []rune, letters []rune, words []string) []string{

    if len(letters) == 0 {
        return words
    }

    for index, letter := range(letters) {
        possible := append(check, letter)
        fmt.Println("Checking", string(possible))
        if w.IsValid(string(check)) {
            fmt.Println("Valid word found", string(possible))
            words = append(words, string(possible))

        }
        left := append(letters[:index], letters[index+1:]...)
        return w.findwords(possible, left, words)
    }

    check = append(check, letters[0])
    return w.findwords(check, letters[1:], words)

}
