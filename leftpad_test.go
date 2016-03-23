package leftpad

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeftPad(t *testing.T) {
	assert.Equal(t, LeftPad("foo", 3, 'x'), "foo", "Input does not match expected output")
	assert.Equal(t, LeftPad("foo", 2, 'x'), "foo", "Input does not match expected output")
	assert.Equal(t, LeftPad("foo", 6, 'x'), "xxxfoo", "Input does not match expected output")
	assert.Equal(t, LeftPad("foo", 10, '™'), "™™™™™™™foo", "Input does not match expected output")
	assert.Equal(t, LeftPad("foo", 10, '世'), "世世世世世世世foo", "Input does not match expected output")
	var r rune
	assert.Equal(t, LeftPad("foo", 10, r), "foo", "Input does not match expected output")
}
