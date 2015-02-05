package words

import (
    "sort"
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

func TestPlayerOneEmptyPlayerTwoScoresFullLengthIsNegative(t *testing.T) {
    want := -3
    playerOne := ""
    playerTwo := "win"
    if result := ScoreWords(playerOne, playerTwo); !reflect.DeepEqual(result, want) {
        t.Errorf("scoreWords(%s, %s) = %+v, want %+v", playerOne, playerTwo, result, want)
    }
}

func TestPlayerTwoEmptyPlayerOneScoresFullLength(t *testing.T) {
    want := 3
    playerOne := "won"
    playerTwo := ""
    if result := ScoreWords(playerOne, playerTwo); !reflect.DeepEqual(result, want) {
        t.Errorf("scoreWords(%s, %s) = %+v, want %+v", playerOne, playerTwo, result, want)
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
    hand := NewLetterMap("abck")
    done := make(chan struct{})
    for entry := range w.PossibleWords(done, hand) {
        result = append(result, entry)
    }
    sort.Sort(ByLength(result))
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
    hand := NewLetterMap("akkc")
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
    hand := NewLetterMap("abkc")
    done := make(chan struct{})
    for entry := range w.PossibleWords(done, hand) {
        result = append(result, entry)
    }
    if !reflect.DeepEqual(result, want) {
        t.Errorf("PossibleWords(%s) = %+v, want %+v", hand, result, want)
    }
}

// Easy optimization possible, check "checkbylength" function for explaination
func BenchmarkPossibleWords(b *testing.B) {
    w := NewWordset()
    w.LoadWordsFromFile("../data/wordset.txt")
    hand := NewLetterMap("abcdefghijkl")
    done := make(chan struct{})
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var compiled []string
        for result := range w.PossibleWords(done, hand) {
            compiled = append(compiled, result)
        }
    }
}
