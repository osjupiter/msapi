package main

import "io"

type stringMap map[string]string
type Complexobj interface{}

type Generator interface {
	Generate() (Complexobj, error)
}

type Filter interface {
	fillter(Complexobj) error
}

type Printer interface {
	Print(Complexobj) (io.Reader, error)
}

type Upstream interface {
	Upload(io.Reader) (io.Reader, error)
}

type Consumer interface {
	Consume(io.Reader) (Complexobj, error)
}
