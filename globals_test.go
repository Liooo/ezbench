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

func ExampleOnce() {
	ezbench.Once("yo", func() {
		time.Sleep(time.Millisecond * 300)
	})
	// Output: yo took 300.0839ms
}

func ExampleSeq(){
	ezbench.Seq()
	time.Sleep(time.Millisecond * 100)

	ezbench.Seq()
	time.Sleep(time.Millisecond * 500)

	ezbench.Seq()
	time.Sleep(time.Millisecond * 300)
	ezbench.SeqPrint()
	// Output:
	// sequential benchmark results:
	//   1st Seq: took 105.038033ms
	//   2nd Seq: took 500.244891ms
	//   3rd Seq: took 300.928543ms
}

func ExampleSeq_named(){
	ezbench.Seq()
	time.Sleep(time.Millisecond * 200)

	ezbench.Seq("I'll take long")
	time.Sleep(time.Millisecond * 300)
	time.Sleep(time.Millisecond * 300)
	time.Sleep(time.Millisecond * 300)

	ezbench.Seq()
	time.Sleep(time.Millisecond * 100)

	ezbench.SeqPrint()
	// Output:
	// sequential benchmark results:
	//   1st Seq: took 100.85786ms
	//   I'll take long: took 900.675695ms
	//   3rd Seq: took 100.72323ms.
}

func ExamplePar(){
	ezbench.Par("short", func() { time.Sleep(time.Millisecond * 200) })
	ezbench.Par("long", func() { time.Sleep(time.Millisecond * 500) })
	ezbench.Par("long", func() { time.Sleep(time.Millisecond * 300) })

	ezbench.ParPrint()
	// Output:
	// parallel benchmark results:
	//   long: called twice, took 800.323144ms
	//   short: called once, took 203.587983ms
}

func ExampleParStart(){
	ezbench.ParStart("long")

	ezbench.Par("short", func() { time.Sleep(time.Millisecond * 200) })
	time.Sleep(time.Millisecond * 300)

	ezbench.ParEnd("long")

	ezbench.Par("long", func() { time.Sleep(time.Millisecond * 300) })

	ezbench.ParPrint()
	// Output:
	// parallel benchmark results:
	//   long: called twice, took 800.323144ms
	//   short: called once, took 203.587983ms
}
