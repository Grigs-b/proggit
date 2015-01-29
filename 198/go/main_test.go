package main

import (
    "reflect"
    "testing"
)

func TestEqualWordsTie(t *testing.T) {
    want := 0
    word := "because"
    if result := scoreWords(word, word); !reflect.DeepEqual(result, want) {
        t.Errorf("scoreWords(%s, %s) = %+v, want %+v", word, word, result, want)
    }
}

func TestTieLeftOverLettersEqual(t *testing.T) {
    want := 0
    playerOne := "ballet" // "bl" left over
    playerTwo := "talent" // "tn" left over
    if result := scoreWords(playerOne, playerTwo); !reflect.DeepEqual(result, want) {
        t.Errorf("scoreWords(%s, %s) = %+v, want %+v", playerOne, playerTwo, result, want)
    }
}

func TestPlayerOneWinsReturnsPositive(t *testing.T) {
    want := 2
    playerOne := "ballet" // "be" left over
    playerTwo := "tall" // nothing left
    if result := scoreWords(playerOne, playerTwo); !reflect.DeepEqual(result, want) {
        t.Errorf("scoreWords(%s, %s) = %+v, want %+v", playerOne, playerTwo, result, want)
    }
}

func TestPlayerTwoWinsReturnsNegative(t *testing.T) {
    want := -1
    playerOne := "pwn" // no matches, 3 letters
    playerTwo := "rekt" // no matches, 4 letters
    if result := scoreWords(playerOne, playerTwo); !reflect.DeepEqual(result, want) {
        t.Errorf("scoreWords(%s, %s) = %+v, want %+v", playerOne, playerTwo, result, want)
    }
}
