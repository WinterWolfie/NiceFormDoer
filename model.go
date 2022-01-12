package main

type Config struct {
	Title string
	Form  string

	Quantity int

	Questions []Question
}

type Question struct {
	Label   string
	PostKey string
	Options []Option
}

type Option struct {
	Value  string
	Chance int
}

type Request struct {
	Form   string
	Params []Param
}
type Param struct {
	Key   string
	Value string
}
