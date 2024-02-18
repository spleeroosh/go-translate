package handler

import "testing"

func TestHandleAuthShow(t *testing.T) {
	err := a.HandleAuthShow(c)
	t.Log(err)
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}
}
