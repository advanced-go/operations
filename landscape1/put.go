package landscape1

import (
	"context"
	"errors"
	"github.com/advanced-go/operations/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"time"
)

func put[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, resource, template string, body []T, insert pgxsql.InsertFuncT[T]) (h2 http.Header, status *core.Status) {
	var e E

	if len(body) == 0 {
		status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: no entries found"))
		return nil, status
	}
	if insert == nil {
		insert = testInsert[T] //pgxsql.InsertT[T]
	}
	h2 = httpx.Forward(h2, h)
	h2.Set(core.XFrom, module.Authority)
	_, status = insert(ctx, h, resource, template, body)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return
}

func testInsert[T pgxsql.Scanner[T]](_ context.Context, _ http.Header, resource, template string, entries []T, args ...any) (tag pgxsql.CommandTag, status *core.Status) {
	switch p := any(&entries).(type) {
	case *[]Entry:
		status = insertEntry(*p)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	if status.OK() {
		tag.RowsAffected = int64(len(entries))
	}
	return
}

func insertEntry(body []Entry) *core.Status {
	if len(body) == 0 {
		return core.StatusBadRequest()
	}
	defer safeEntry.Lock()()

	for _, e := range body {
		e.CreatedTS = time.Now().UTC()
		e.Partition = lastEntry().Partition + 1
		entryData = append(entryData, e)
	}
	return core.StatusOK()
}
