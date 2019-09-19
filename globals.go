package ezbench

import (
	"fmt"
	"time"
)

var gSeq seqs
var gPar pars

func init() {
	gSeq = emptySequential()
	gPar = emptyPars()
}

// == Sequential ==

// Seq starts a new timer and stops the last timer if any
func Seq(name ...string) {
	if !gSeq.IsGoingOn() {
		gSeq.StartNext(gSeq.toName(name...))
		return
	}
	gSeq.CommitLast()
	gSeq.StartNext(gSeq.toName(name...))
}

// SeqPrint stops the last timer and prints all benchmark information measured with Seq method
func SeqPrint() {
	gSeq.Print()
	gSeq = emptySequential()
}

func UneasyGlobalSequentials() seqs {
	return gSeq
}

func UneasyNewSequentials() seqs {
	return emptySequential()
}

// == Once ==

// Once executes benchee and immediately prints the time taken
func Once(name string, benchee func()) {
	t := time.Now()
	benchee()
	fmt.Printf("%s took %s\n", name, time.Since(t))
}

// == Parallel ==

// Par executes benchee and stores benchmark
func Par(name string, benchee func()) {
	gPar.MeasureF(name, benchee)
}

// ParStart starts taking benchmark until ParEnd is called with the same name
func ParStart(name string) {
	gPar.StartTimer(name)
}

// ParEnd stops taking the benchmark started by last PerStart call with the same name
func ParEnd(name string) {
	gPar.EndTimer(name)
}

// ParPrint prints all benchmark information measured with Par* methods
func ParPrint() {
	gPar.Print()
	gPar = emptyPars()
}

func UneasyGlobalParallels() pars {
	return gPar
}

func UneasyNewParallels() pars {
	return emptyPars()
}
