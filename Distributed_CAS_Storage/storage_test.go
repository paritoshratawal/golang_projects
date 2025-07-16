package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	path := CASPathTransformFunc("Greates Porn movies")
	fmt.Println(path)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	store := NewStore(opts)

	data := bytes.NewReader([]byte("some jpeg bytes hi hello"))

	if err := store.WriteStream("My Special Picture", data); err != nil {
		t.Error(err)
	}
}
