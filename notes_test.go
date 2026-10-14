package main

import "testing"

func TestCombineNotes(t *testing.T) {
	oldText := "первая заметка"
	newText := "вторая заметка"

	result := combineNotes(oldText, newText)

	expected := "первая заметка\nвторая заметка"
	if result != expected {
		t.Errorf("Ожидалось %q, получили %q", expected, result)
	}
}