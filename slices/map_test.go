package slices_test

import (
	"errors"
	"testing"

	"github.com/boundedinfinity/commons/slices"
	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	type Type1 struct {
		thing string
	}

	type Type2 struct {
		thing string
	}

	expected := []Type2{{thing: "a"}, {thing: "b"}}
	input := []Type1{{thing: "a"}, {thing: "b"}}
	actual := slices.Map(input, func(t1 Type1) Type2 {
		return Type2{thing: t1.thing}
	})

	assert.ElementsMatch(t, expected, actual)
}

func Test_MapErr_NoErr(t *testing.T) {
	type Type1 struct {
		thing string
	}

	type Type2 struct {
		thing string
	}

	expected := []Type2{{thing: "a"}, {thing: "b"}}
	input := []Type1{{thing: "a"}, {thing: "b"}}
	actual, err := slices.MapErr(input, func(t1 Type1) (Type2, error) {
		return Type2{thing: t1.thing}, nil
	})

	assert.ElementsMatch(t, expected, actual)
	assert.Nil(t, err)
}

func Test_MapErr_WithErr(t *testing.T) {
	type Type1 struct {
		thing string
	}

	type Type2 struct {
		thing string
	}

	expected := []Type2{}
	input := []Type1{{thing: "a"}, {thing: "b"}}
	actual, err := slices.MapErr(input, func(t1 Type1) (Type2, error) {
		return Type2{}, errors.New("map error")
	})

	assert.ElementsMatch(t, expected, actual)
	assert.NotNil(t, err)
}
