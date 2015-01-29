package words

import (
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
