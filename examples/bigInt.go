package main

type Int struct {
	neg bool // sign
	abs nat  // absolute value of the integer
}

type nat []Word

type Word uint
