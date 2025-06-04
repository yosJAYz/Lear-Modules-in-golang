package testingmodule

import "testing"

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
