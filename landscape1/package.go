package landscape1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	PkgPath           = "github/advanced-go/operations/landscape1"
	partitionResource = "partition"
)

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) (entries []Entry, status *core.Status) {
	return get[core.Log, Entry](ctx, h, values, partitionResource, "", nil)
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) (status *core.Status) {
	if r == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status1 := json2.New[[]Entry](r.Body, r.Header)
		if !status1.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return status1
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), partitionResource, "", body, nil)
}
