package main

type Direction int

const (
	L Direction = -1
	R Direction = 1
)

type Action struct {
	direction Direction
	distance  int
}
