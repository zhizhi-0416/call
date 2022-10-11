package main

import (
	"fmt"
	"testing"
)

func TestGenerateData(t *testing.T) {
	lesson := GenerateData(1)
	err := WriteToFile(&lesson)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFile(t *testing.T) {
	lesson, err := ReadFile(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(lesson.Students))
	for i := range lesson.Students {
		t.Log(*lesson.Students[i])
	}
}
