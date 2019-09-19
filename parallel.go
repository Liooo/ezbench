package ezbench

import (
	"fmt"
	"sort"
	"time"
)

type parKpi struct {
	count int
	d     time.Duration
	lastT time.Time
}

func (p *parKpi) IsGoingOn() bool {
	return !p.lastT.IsZero()
}

type pars struct {
	ks map[string]*parKpi
}

func emptyPars() pars {
	return pars{
		ks: make(map[string]*parKpi),
	}
}

func (p *pars) MeasureF(name string, benchee func()) {
	if _, ok := p.ks[name]; !ok {
		p.ks[name] = &parKpi{}
	}

	t := time.Now()
	benchee()
	p.ks[name].count++
	p.ks[name].d += time.Since(t)
}

func (p *pars) StartTimer(name string) {
	if _, ok := p.ks[name]; !ok {
		p.ks[name] = &parKpi{}
	}
	p.ks[name].lastT = time.Now()
}

func (p *pars) EndTimer(name string) {
	p.ks[name].count++
	p.ks[name].d += time.Since(p.ks[name].lastT)
	p.ks[name].lastT = time.Time{}
}

func (p *pars) Print() {
	var keys []string
	for k, _ := range p.ks {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return p.ks[keys[i]].d > p.ks[keys[j]].d
	})

	fmt.Printf("parallel benchmark results:\n")
	for _, k := range keys {
		fmt.Printf("  %s: called %s, took %s\n", k, toCount(p.ks[k].count), p.ks[k].d)
		if p.ks[k].IsGoingOn() {
			fmt.Printf("  !!! forgot to call ParEnd(%q)?\n", k)
		}
	}
}
