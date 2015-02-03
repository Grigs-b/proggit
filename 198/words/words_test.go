package words

import (
    "fmt"
    "reflect"
    "testing"
)

func TestEqualWordsTie(t *testing.T) {
    want := 0
    word := "because"
    if result := ScoreWords(word, word); !reflect.DeepEqual(result, want) {
        t.Errorf("scoreWords(%s, %s) = %+v, want %+v", word, word, result, want)
    }
}

func TestTieLeftOverLettersEqual(t *testing.T) {
    want := 0
    playerOne := "ballet" // "bl" left over
    playerTwo := "talent" // "tn" left over
    if result := ScoreWords(playerOne, playerTwo); !reflect.DeepEqual(result, want) {
        t.Errorf("scoreWords(%s, %s) = %+v, want %+v", playerOne, playerTwo, result, want)
    }
}

func TestPlayerOneWinsReturnsPositive(t *testing.T) {
    want := 2
    playerOne := "ballet" // "be" left over
    playerTwo := "tall" // nothing left
    if result := ScoreWords(playerOne, playerTwo); !reflect.DeepEqual(result, want) {
        t.Errorf("scoreWords(%s, %s) = %+v, want %+v", playerOne, playerTwo, result, want)
    }
}

func TestPlayerTwoWinsReturnsNegative(t *testing.T) {
    want := -1
    playerOne := "pwn" // no matches, 3 letters
    playerTwo := "rekt" // no matches, 4 letters
    if result := ScoreWords(playerOne, playerTwo); !reflect.DeepEqual(result, want) {
        t.Errorf("scoreWords(%s, %s) = %+v, want %+v", playerOne, playerTwo, result, want)
    }
}

func TestWordlistWordExistsReturnsTrue(t *testing.T) {
    want := true

    w := NewWordset()
    w.AddWord("ball")
    w.AddWord("court")
    w.AddWord("base")
    if result := w.IsValid("court"); !reflect.DeepEqual(result, want) {
        t.Errorf("IsValid(%s) = %+v, want %+v", "court", result, want)
    }
}

func TestWordlistWordDoesntExistReturnsFalse(t *testing.T) {
    want := false

    w := NewWordset()
    w.AddWord("ball")
    w.AddWord("court")
    w.AddWord("base")
    if result := w.IsValid("baseball"); !reflect.DeepEqual(result, want) {
        t.Errorf("IsValid(%s) = %+v, want %+v", "baseball", result, want)
    }
}

func TestPossibleWordsFindsCorrectSetGoldenPath(t *testing.T) {
    want := []string{"cab", "back"}
    var result []string
    w := NewWordset()
    w.AddWord("cab")
    w.AddWord("truth")
    w.AddWord("back")
    w.AddWord("front")
    hand := []rune{'a', 'b', 'c','k'}
    done := make(chan struct{})
    for entry := range w.PossibleWords(done, hand) {
        result = append(result, entry)
    }
    if !reflect.DeepEqual(result, want) {
        t.Errorf("PossibleWords(%s) = %+v, want %+v", hand, result, want)
    }
}

func TestPossibleWordsMultipleLettersDontCount(t *testing.T) {
    var want, result []string

    w := NewWordset()
    w.AddWord("cab")
    w.AddWord("truth")
    w.AddWord("back")
    w.AddWord("front")
    hand := []rune{'a', 'k', 'c','k'}
    done := make(chan struct{})
    for entry := range w.PossibleWords(done, hand) {
        result = append(result, entry)
    }
    if !reflect.DeepEqual(result, want) {
        t.Errorf("PossibleWords(%s) = %+v, want %+v", hand, result, want)
    }
}

func TestPossibleWordsNonePossibleReturnsEmptyList(t *testing.T) {
    var want, result []string

    w := NewWordset()
    hand := []rune{'a', 'b', 'c','k'}
    done := make(chan struct{})
    for entry := range w.PossibleWords(done, hand) {
        result = append(result, entry)
    }
    if !reflect.DeepEqual(result, want) {
        t.Errorf("PossibleWords(%s) = %+v, want %+v", hand, result, want)
    }
}

func BenchmarkPossibleWords(b *testing.B) {
    w := NewWordset()
    w.LoadWordsFromFile("../data/wordset.txt")
    hand := []rune{'a', 'b', 'c','d', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}
    done := make(chan struct{})
    var result []string
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for entry := range w.PossibleWords(done, hand) {
            result = append(result, entry)
        }
        fmt.Println(result)
    }
}
