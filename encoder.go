package xdiff

import (
	"fmt"
	"io"
)

// Encoder knows how to interpret diff results.
type Encoder interface {
	Encode([]Delta) error
}

type plainTextEncoder struct {
	w io.Writer
}

func (pte *plainTextEncoder) Encode(deltas []Delta) error {
	if len(deltas) == 0 {
		fmt.Fprint(pte.w, "No difference.\n")
	}
	for _, d := range deltas {
		if d.Op == Update {
			fmt.Fprintf(pte.w, "%s('%s'->'%s')\n", d.Op, d.Node, d.Update)
			continue
		}
		fmt.Fprintf(pte.w, "%s('%s')\n", d.Op, d.Node)
	}
	return nil
}

// PlainTextEncoder outputs diff results in plain text format.
func PlainTextEncoder(w io.Writer) Encoder {
	return &plainTextEncoder{w}
}
