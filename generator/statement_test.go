package generator

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestToQuestion(t *testing.T) {
    // GIVEN
    number := 13
    representation := Representation{ "binary", 2 }

    statement := Statement{ number, representation }

    // WHEN
    expected := statement.ToQuestion()

    // THEN
    assert.Equal(t, expected, "Turn 13 into binary: ")
}

func TestSolve(t *testing.T) {
    // GIVEN
    number := 13
    representation := Representation{ "binary", 2 }

    statement := Statement{ number, representation }

    // WHEN
    expected := statement.Solve()

    // THEN
    assert.Equal(t, expected, "1101")
}

