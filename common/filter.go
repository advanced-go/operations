package common

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"strconv"
)

func FilterT[T any](values url.Values, entries []T, valid func(url.Values, T) bool) ([]T, *core.Status) {
	if values == nil {
		return entries, core.StatusBadRequest()
	}
	if len(entries) == 0 {
		return entries, core.StatusNotFound()
	}
	var result []T

	for _, e1 := range entries {
		if valid(values, e1) {
			result = append(result, e1)
		}
	}
	if len(result) == 0 {
		return result, core.StatusNotFound()
	}
	result = Order(values, result)
	return Top(values, result), core.StatusOK()
}

func Order[T any](values url.Values, entries []T) []T {
	if entries == nil || values == nil {
		return entries
	}
	s := values.Get("order")
	if s != "desc" {
		return entries
	}
	var result []T

	for i := len(entries) - 1; i >= 0; i-- {
		result = append(result, entries[i])
	}
	return result
}

func Top[T any](values url.Values, entries []T) []T {
	if entries == nil || values == nil {
		return entries
	}
	s := values.Get("top")
	if s == "" {
		return entries
	}
	cnt, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("top value is not valid: %v", s)
	}
	var result []T
	for i, e := range entries {
		if i < cnt {
			result = append(result, e)
		} else {
			break
		}
	}
	return result
}
