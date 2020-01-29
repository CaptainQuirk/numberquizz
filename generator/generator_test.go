package generator

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGenerateStatement(t *testing.T) {
    // GIVEN
    max := 15

    // WHEN
    got := GenerateStatement(max)

    // THEN
    assert.GreaterOrEqual(t, got.Number(), 0, "Number should be greater than 0")
    assert.Less(t, got.Number(), 16, "Number should be greater than 0")
}
