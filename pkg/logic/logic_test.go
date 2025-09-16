package logic

import "testing"

func TestLogic(t *testing.T) {
	lp := LogicProvider{}
	client := Client{Logic: lp}
	client.Program()
}
