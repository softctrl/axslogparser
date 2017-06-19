package axslogparser

import (
	"strings"
	"time"

	"github.com/Songmu/go-ltsv"
	"github.com/pkg/errors"
)

// LTSV access log parser
type LTSV struct {
}

// Parse for Parser interface
func (lv *LTSV) Parse(line string) (*Log, error) {
	l := &Log{}
	err := ltsv.Unmarshal([]byte(line), l)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse ltsvlog")
	}
	l.Time, _ = time.Parse(clfTimeLayout, strings.Trim(l.TimeStr, "[]"))
	if err := l.breakdownRequest(); err != nil {
		return nil, errors.Wrap(err, "failed to parse ltsvlog")
	}
	return l, nil
}
