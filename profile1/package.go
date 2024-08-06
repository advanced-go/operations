package profile1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/operations/profile1"
)

func Get(ctx context.Context) (Entry, *core.Status) {
	return entry, core.StatusOK()
}

// Need some way to periodically query access log to determine if there is a traffic spike occurring
