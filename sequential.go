package ezbench

import (
	"fmt"
	"time"
)

type seqKpi struct {
	name string
	d    time.Duration
}

type seqs struct {
	ks    []seqKpi
	lastT time.Time
	lastN string
}

func emptySequential() seqs {
	return seqs{
		ks:    []seqKpi{},
		lastT: time.Time{},
		lastN: "",
	}
}

func (s *seqs) CommitLast() {
	if !s.IsGoingOn() {
		return
	}
	k := seqKpi{
		name: s.lastN,
		d:    time.Since(s.lastT),
	}
	s.ks = append(s.ks, k)
}

func (s *seqs) StartNext(name string) {
	s.lastN = name
	s.lastT = time.Now()
}

func (s *seqs) IsGoingOn() bool {
	return !s.lastT.IsZero()
}

func (s *seqs) toName(name ...string) string {
	if len(name) == 0 {
		return s.nthDefaultName(len(s.ks) + 1)
	} else if len(name) == 1 {
		return name[0]
	}
	return "you like names huh"
}

func (s *seqs) nthDefaultName(n int) string {
	return fmt.Sprintf("%s Seq", toOrder(n))
}

func (s *seqs) Print() {
	s.CommitLast()
	fmt.Printf("sequential benchmark results:\n")
	for _, v := range s.ks {
		fmt.Printf("  %s: took %s\n", v.name, v.d)
	}
}
