package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	add := Add(1, 2)
	assert.Equal(t, 3, add)
}
