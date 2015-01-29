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
var letters = []rune("abcdefghijklmnopqrstuvwxyz")


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
    Hand    []rune
    Name    string
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
    return p.Hand
}

func (p *AnyPlayer) GetName() string {
    return p.Name
}

func (p *AnyPlayer) SetName(name string) {
    p.Name = name
}

func getLetter() rune {
    // simple random for right now
    return letters[rand.Intn(len(letters))]
}

// TODO: Different letter distributions? Scrabble frequency?
// TODO: Ensure player has vowels so they are able to make words
func (p *AnyPlayer) FillHand() {
    for i := len(p.Hand); i < HAND_SIZE; i++ {
        p.Hand = append(p.Hand, getLetter())
    }
}

func (p *AIPlayer) GetName() string {
    return "Computer"
}

// TODO: Decouple the wordlist
func (ai *EasyAIPlayer) Play() string {
    choice := ""
    possible := words.Wordlist.PossibleWords(ai.Hand)
    if len(possible) > 0 {
        choice = possible[0]
        for _, word := range possible {
            if len(word) < len(choice) {
                choice = word
            }
        }
    }
    return choice
}

// TODO: Make Medium choose a normal distribution from sorted.length rather than random
func (ai *MediumAIPlayer) Play() string {
    choice := ""
    possible := words.Wordlist.PossibleWords(ai.Hand)
    if len(possible) > 0 {
        val := rand.Intn(len(possible))
        choice = possible[val]
    }
    return choice
}

func (ai *HardAIPlayer) Play() string {
    choice := ""
    possible := words.Wordlist.PossibleWords(ai.Hand)
    if len(possible) > 0 {
        choice = possible[0]
        for _, word := range possible {
            if len(word) > len(choice) {
                choice = word
            }
        }
    }
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
        if strings.EqualFold(result, "-1") {
            //escape for if you can't spell a word, temporary until I find a better solution
            fmt.Printf("%s couldn't spell a word!\n", human.GetName())
            result = ""
            break
        } else if !words.Wordlist.IsValid(result) {
            fmt.Printf("Word %s not found\n", result)
        } else if !human.CheckLettersVsHand(result) {
            fmt.Printf("Word %s contained characters not in your hand: %s\n", result, string(human.GetHand()))
        } else {
            break
        }

    }
    return result
}

func (p *AnyPlayer) CheckLettersVsHand(word string) bool {
    hand := make([]rune, len(p.Hand))
    copy(hand, p.Hand)
    for _, letter := range word {
        index := -1
        for idx, handletter := range hand {
            if letter == handletter {
                index = idx
                break
            }
        }
        if index >= 0 {
            hand = append(hand[:index], hand[index+1:]...)
        } else {
            return false
        }
    }
    p.Hand = hand
    return true
}
