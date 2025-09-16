package logic

import "fmt"

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	fmt.Println("Processing data:", data)
	return "Processed"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	Logic Logic
}

func (c Client) Program() {
	c.Logic.Process("Some data")
}
