package network1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath = "github/advanced-go/operations/network1"
)

func Profile(ctx context.Context) (Entry, *core.Status) {
	return entry, core.StatusOK()
}

// Need some way to periodically query access log to determine if there is a traffic spike occurring
