package main

import "fmt"

type node[T any] struct {
	Data T
	prev *node[T]
	next *node[T]
}

type list[T any] struct {
	first *node[T]
	last *node[T]
}

func (l list[T]) add(data T) *node[T] {
	n := node[T] {
		Data: T,
		prev: l.last,
	}

	if l.first == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	l.last.next = &n
	l.last = &n
	return &n
}
