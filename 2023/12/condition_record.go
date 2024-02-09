package main

import (
	"bytes"

	"github.com/kmullin/advent-of-code/2023/common"
)

type ConditionRecord struct {
	r []springRow
}

func (r *ConditionRecord) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return common.InputEmptyErr
	}
	for _, b := range bytes.Split(text, []byte("\n")) {
		if len(b) == 0 {
			continue
		}

		sr, err := newSpringRow(b)
		if err != nil {
			return err
		}

		r.r = append(r.r, sr)
	}
	return nil
}

func (r ConditionRecord) TotalArrangements() (sum int) {
	for _, sr := range r.r {
		sum += sr.TotalArrangements()
	}
	return
}
