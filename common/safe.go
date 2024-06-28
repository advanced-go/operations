package common

import "sync"

type SafeSlice[T any] struct {
	list []T
	m    sync.Mutex
	//index map[string]T
}

func NewSafeSlice[T any](items []T) *SafeSlice[T] {
	s := new(SafeSlice[T])
	if len(items) > 0 {
		s.Append(items)
	}
	return s
}

func (s *SafeSlice[T]) Lock() {
	s.m.Lock()
}

func (s *SafeSlice[T]) Unlock() {
	s.m.Unlock()
}

func (s *SafeSlice[T]) Append(items []T) {
	//s.m.Lock()
	//defer s.m.Unlock()
	s.list = append(s.list, items...)
}

type Safe struct {
	mu sync.Mutex
}

func (s *Safe) Lock() func() {
	s.mu.Lock()
	return func() {
		s.mu.Unlock()
	}
}

func (s *Safe) Unlock() {
	s.mu.Unlock()
}

func NewSafe() *Safe {
	s := new(Safe)
	return s
}
