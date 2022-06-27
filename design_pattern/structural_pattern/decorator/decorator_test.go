package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewColorSquare(t *testing.T) {
	colorSquare := NewColorSquare(Square{}, "green")
	assert.Equal(t, "draw square color is: green", colorSquare.Draw())
}
