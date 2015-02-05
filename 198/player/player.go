package player

import (
    "os"
    "bufio"
    "strings"
    "math/rand"
    "fmt"
    "github.com/Grigs-b/proggit/198/words"
)

const HAND_SIZE int = 12



type Player interface {
    Play() string
    GetScore() int
    GetHand() []rune
    AddPoints(points int)
    FillHand()
    SetName(string)
    GetName() string
}

type AnyPlayer struct {
    Score   int
    Hand    words.LetterMap
    Name    string
    Words   words.WordList
}

type HumanPlayer struct {
    AnyPlayer
    Name    string
}

type AIPlayer struct {
    AnyPlayer
}

type EasyAIPlayer struct {
    AIPlayer
}

type MediumAIPlayer struct {
    AIPlayer
}

type HardAIPlayer struct {
    AIPlayer
}

func (p *AnyPlayer) GetScore() int {
    return p.Score
}

func (p *AnyPlayer) AddPoints(points int) {
    p.Score += points
}

func (p *AnyPlayer) GetHand() []rune {
    var hand []rune
    for letter, count := range p.Hand {
        for i := 0; i < count; i++ {
            hand = append(hand, letter)
        }
    }
    return hand
}

func (p *AnyPlayer) GetName() string {
    return p.Name
}

func (p *AnyPlayer) SetName(name string) {
    p.Name = name
}


// TODO: Different letter distributions? Scrabble frequency?
// TODO: Ensure player has vowels so they are able to make words
func (p *AnyPlayer) FillHand() {
    if p.Hand == nil {
        p.Hand = words.NewLetterMap("")
    }
    sum := 0
    for _, count := range p.Hand {
        sum += count
    }
    for i := sum; i < HAND_SIZE; i++ {
        p.Hand[words.GetLetter()]++
    }
}

func (p *AIPlayer) GetName() string {
    return "Computer"
}

// TODO: Add Sortby length to possible words to make choosing shortest/longest/etc easier
func (ai *EasyAIPlayer) Play() string {
    choice := ""
    done := make(chan struct{})
    defer close(done)
    i := 0
    for possible := range ai.Words.PossibleWords(done, ai.Hand) {
        if i == 0 {
            choice = possible
        }
        if len(possible) < len(choice) {
            choice = possible
        }
    }
    ai.UpdateHand(choice)
    return choice
}

// TODO: Other AI Improvements, change to try a min/max "target" length to aim for?
//       Easy AI could be something like: try 2-4 length words, then 5..12 until you find one
//       Medium could be: try 4-5 length, then 2,3,6..12
//       Hard: try 12...2
func (ai *MediumAIPlayer) Play() string {

    done := make(chan struct{})
    defer close(done)
    var compiled []string
    choice := ""
    for possible := range ai.Words.PossibleWords(done, ai.Hand) {
        compiled = append(compiled, possible)
    }

    if len(compiled) > 0 {
        val := rand.Intn(len(compiled))
        choice = compiled[val]
    }
    ai.UpdateHand(choice)
    return choice
}

func (ai *HardAIPlayer) Play() string {

    done := make(chan struct{})
    defer close(done)
    choice := ""
    for possible := range ai.Words.PossibleWords(done, ai.Hand) {
        if len(possible) > len(choice) {
            choice = possible
        }
    }
    ai.UpdateHand(choice)
    return choice
}


// TODO: duplicative of code in main
func (human HumanPlayer) ReadIO() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter word: ")
    text, _ := reader.ReadString('\n')
    return strings.TrimSpace(text)
}

func (human *HumanPlayer) Play() string {
    result := ""
    for {
        result = human.ReadIO()
        result = strings.ToLower(result)
        if strings.EqualFold(result, "-1") {
            //escape for if you can't spell a word, temporary until I find a better solution
            fmt.Printf("%s couldn't spell a word!\n", human.GetName())
            result = ""
            break
        } else if !human.Words.IsValid(result) {
            fmt.Printf("Word %s not found\n", result)
        } else if !human.UpdateHand(result) {
            fmt.Printf("Word %s contained characters not in your hand: %s\n", result, string(human.GetHand()))
        } else {
            break
        }

    }
    return result
}

func (p *AnyPlayer) UpdateHand(word string) bool {
    tocheck := words.NewLetterMap(word)

    for letter, count := range tocheck {
        if p.Hand[letter] < count {
            return false
        }
    }
    // going through twice so we don't partially remove from hand on first pass
    //  possible to condense?
    for letter, count := range tocheck {
        p.Hand[letter] -= count
    }

    return true
}
