package words

import (
    "io"
    "os"
    "sync"
    "strings"
    "bufio"
)

type WordList interface {
    IsValid(word string) bool
    PossibleWords(done <-chan struct{}, letters []rune) <-chan string
    AddWord(string)
}

type Wordset struct {
    dictionary  map[string]bool
    bylength    map[int][]string
}

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}

func (s ByLength) Swap(i int,j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i int,j int) bool {
    return len(s[i]) < len(s[j])
}

func NewWordset() *Wordset {
    return &Wordset{dictionary: make(map[string]bool), bylength: make(map[int][]string)}
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
    w.bylength[len(word)] = append(w.bylength[len(word)], word)
}


func (w Wordset) IsValid(word string) bool {
    return w.dictionary[word]
}

func check(word string, letters []rune) bool {
    var tmp = word

    for _, letter := range letters {
        tmp = strings.Replace(tmp, string(letter), "", 1)
        if len(tmp) == 0 {
            return true
        }
    }
    return false
}

func (w Wordset) checkbylength(done <-chan struct{}, length int, letters []rune) <-chan string {
    result := make(chan string)
    go func() {
        defer close(result)
        for _, word := range w.bylength[length] {
            if check(word, letters) {
                select {
                case result <- word:
                    // Design Note: I could cut runtime roughly in half by adding
                    //  a return statement here, which would stop execution at the first
                    //  word of len(length) found. I choose not to in this case as
                    //  I wanted to find all words we can match and the game is responsive
                    //  enough in it's current state. But this is an easy optimization to make
                    //  if I wanted as much speed as possible
                case <-done:
                    return
                }
            }
        }

    }()

    return result
}

func merge(done <-chan struct{}, checks ...<-chan string) <-chan string {
    var wg sync.WaitGroup
    out := make(chan string)

    // Start an output goroutine for each input channel
    output := func(c <-chan string) {
        defer wg.Done()
        for n := range c {
            select {
            case out <- n:
            case <-done:
                return
            }
        }
    }
    // Set the number of goroutines we're adding
    wg.Add(len(checks))
    for _, check := range checks {
        go output(check)
    }

    // Close once all the output goroutines are done
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func (w *Wordset) PossibleWords(done <-chan struct{}, letters []rune) <-chan string {

    result := make(chan string)
    go func() {
        defer close(result)
        // create a handler function for each length word we'd like to check
        // dont have any 0 or 1 length words, so skip those, giving us 2-12
        // as indeces for our length checks, HANDLENGTH-1 total check functions
        startPos := 2
        checks := make([]<-chan string, len(letters)-1)
        for i := startPos; i <= len(letters); i++ {
            checks[i-startPos] = w.checkbylength(done, i, letters)
        }

        // Consume the merged output from all checks, handle
        for n := range merge(done, checks...) {
            select {
            case result <- n:
            case <-done:
                return
            }
        }
    }()
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
