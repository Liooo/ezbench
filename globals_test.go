package ezbench_test

import (
	"github.com/liooo/ezbench"
	"testing"
	"time"
)

func TestSeq(t *testing.T) {
	ezbench.Seq()
	time.Sleep(time.Millisecond * 100)
	ezbench.Seq("seq1")
	time.Sleep(time.Millisecond * 500)
	ezbench.Seq("seq2")
	time.Sleep(time.Millisecond * 300)
	ezbench.SeqPrint()

	ezbench.Seq("seq1")
	time.Sleep(time.Millisecond * 100)
	ezbench.Seq()
	time.Sleep(time.Millisecond * 100)
	ezbench.SeqPrint()
}

func TestPar(t *testing.T) {
	t.Run("parallel", func(t *testing.T) {
		ezbench.ParStart("long")
		ezbench.Par("short", func() { time.Sleep(time.Millisecond * 200) })
		time.Sleep(time.Millisecond * 300)
		ezbench.ParEnd("long")
		ezbench.Par("long", func() {
			time.Sleep(time.Millisecond * 500)
		})

		ezbench.ParPrint()
	})

	t.Run("forgot ParEnd", func(t *testing.T) {
		ezbench.ParStart("forgot")
		ezbench.ParPrint()
	})
}

func TestOnce(t *testing.T) {
	ezbench.Once("yo", func() {
		time.Sleep(time.Millisecond * 300)
	})
}

