package animals

import (
	"testing"
)

func TestElephantFeed(t *testing.T) {
	expect := "Grass"
	actual := ElephantFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestDogFeed(t *testing.T) {
	expect := "DogFood"
	actual := DogFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
