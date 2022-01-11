package main

type Request struct {
	Form string
	Params []Param
}

type Request2 struct {
	Form string
	Questions []Question
}

type Question struct {
	Key string
	Options []Option
}

type Option struct {
	Value  string
	Chance int
}

type Param struct {
	Key    string
	Value  string
}