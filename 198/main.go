package main

import (
    "os"
    "fmt"
    "time"
    "strings"
    "math/rand"
    "bufio"
    "github.com/Grigs-b/proggit/198/words"
    "github.com/Grigs-b/proggit/198/player"
)

const ROUNDS int = 5



func GetInput(showtext string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(showtext)
    text, _ := reader.ReadString('\n')
    return strings.TrimSpace(text)
}


func WaitForValidOpponent(wordset words.WordList) player.Player {
    for {
        t := GetInput("Enter your selection: ")
        fmt.Println("Selected", string(t))
        switch {
        case strings.EqualFold(t, "1"):
            return &player.EasyAIPlayer{AIPlayer: player.AIPlayer{AnyPlayer: player.AnyPlayer{Name:"Easy",Words:wordset}}}
        case strings.EqualFold(t, "2"):
            return &player.MediumAIPlayer{AIPlayer: player.AIPlayer{AnyPlayer: player.AnyPlayer{Name:"Medium",Words:wordset}}}
        case strings.EqualFold(t, "3"):
            return &player.HardAIPlayer{AIPlayer: player.AIPlayer{AnyPlayer: player.AnyPlayer{Name:"Hard",Words:wordset}}}
        case strings.EqualFold(t, "4"):
            p2 := GetInput("Player2 enter your name: ")
            return &player.HumanPlayer{AnyPlayer: player.AnyPlayer{Name:p2,Words:wordset}}
        default:
            fmt.Println("Invalid Choice, Select 1-4\n")
        }
    }
}

func PlayRound(round int, p1 player.Player, p2 player.Player) {
    fmt.Printf("Turn %d -- Points %s: %d %s: %d \n", round, p1.GetName(), p1.GetScore(), p2.GetName(), p2.GetScore())
    p1.FillHand()
    p2.FillHand()
    fmt.Printf("%s letters: ", p1.GetName())
    fmt.Printf(" %s", string(p1.GetHand()))
    fmt.Println("")

    player1word := p1.Play()
    fmt.Printf("%s Selects \"%s\" \n", p1.GetName(), player1word)


    fmt.Printf("%s letters: ", p2.GetName())
    fmt.Printf(" %s", string(p2.GetHand()))

    fmt.Println("")
    player2word := p2.Play()

    fmt.Printf("%s Selects \"%s\" \n", p2.GetName(), player2word)
    score := words.ScoreWords(player1word, player2word)
    winner := "Tie"

    switch {
    case score > 0 :
        winner = p1.GetName()
        fmt.Printf("%s scores %d points\n", p1.GetName(), score)
        p1.AddPoints(score)
        break
    case score < 0 :
        winner = p2.GetName()
        fmt.Printf("%s scores %d points\n", p2.GetName(), -score)
        p2.AddPoints(-score)
        break
    }

    fmt.Printf("\n%s vs %s -- %s wins\n", player1word, player2word, winner)

}

func main() {
    var wordset = words.NewWordset()
    wordset.LoadWordsFromFile("data/wordset.txt")
    //Production:
    rand.Seed(time.Now().UTC().UnixNano())
    //Testing
    //rand.Seed(12)

    fmt.Println("Welcome to Words with Enemies!")
    p1name := GetInput("Player 1 enter your name: ")
    fmt.Println("Select enemy difficulty:")
    fmt.Println("1: Easy")
    fmt.Println("2: Medium")
    fmt.Println("3: Hard")
    fmt.Println("4: Human")


    player1 := &player.HumanPlayer{AnyPlayer: player.AnyPlayer{Name:p1name, Words:wordset}}
    player2 := WaitForValidOpponent(wordset)

    for round := 0; round < ROUNDS; round++ {
        PlayRound(round, player1, player2)
    }
    fmt.Printf("Final Score - %s: %d - %s: %d \n", player1.GetName(), player1.GetScore(), player2.GetName(), player2.GetScore())
    if player1.GetScore() > player2.GetScore() {
        fmt.Printf("%s wins!", player1.GetName())
    } else {
        fmt.Printf("%s wins!", player2.GetName())
    }
    fmt.Println("")
}
