package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunc1(t *testing.T) {
	should := assert.New(t)
	if should.Equal("aaa", "aaa") {
		t.Logf("!=")
	}

}
