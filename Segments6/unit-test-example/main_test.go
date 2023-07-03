package main

import (
	"testing"
	"github.com/stretchr/testify/assert" 	// "github.com/stretchr/testify/tree/master/assert"
)

func TestAddition(t *testing.T)  {
	x := 2
	y := 2

	assert.Equal(t, x, y, "x and x should be the same")
}