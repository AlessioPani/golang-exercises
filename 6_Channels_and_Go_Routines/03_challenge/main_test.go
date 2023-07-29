package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("epsilon")

	wg.Wait()

	if msg != "epsilon" {
		t.Errorf("msg is not epsilon, but got %s", msg)
	}

}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("expected to find epsilon, got other")
	}

}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"
	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "universe") {
		t.Errorf("expected to find universe, got other")
	}

	if !strings.Contains(output, "cosmos") {
		t.Errorf("expected to find cosmos, got other")
	}

	if !strings.Contains(output, "world") {
		t.Errorf("expected to find world, got other")
	}
}
