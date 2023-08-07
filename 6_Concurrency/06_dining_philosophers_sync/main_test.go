package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second

	mealOrder = []string{}
	dine()
	if len(mealOrder) != 5 {
		t.Errorf("Incorrenct lenght of slice, expected 5 but got %d", len(mealOrder))
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var varTests = []struct {
		name  string
		delay time.Duration
	}{
		{"no-delay", 0 * time.Second},
		{"quarter-sec-delay", 250 * time.Millisecond},
		{"half-second-delay", 500 * time.Millisecond},
	}

	for _, i := range varTests {
		mealOrder = []string{}

		eatTime = i.delay
		thinkTime = i.delay

		dine()
		if len(mealOrder) != 5 {
			t.Errorf("Test %s: incorrect lenght of slice, expected 5 but got %d", i.name, len(mealOrder))
		}
	}
}
