package handler

import "testing"

func TestHandleHomeShow(t *testing.T) {
	err := h.HandleHomeShow(c)
	t.Log(err)
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}
}
