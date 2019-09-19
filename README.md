# ezbench

Easiest golang benchmark tool for lazy people.

## Installation

`go get -u github.com/liooo/ezbench`


## Quick Start

### Once 

```
ezbench.Once("yo", func() { time.Sleep(time.Millisecond * 100) })
// yo took 102.656069ms
```


### Sequential

```
ezbench.Seq()
time.Sleep(time.Millisecond * 100)

ezbench.Seq()
time.Sleep(time.Millisecond * 500)

ezbench.Seq()
time.Sleep(time.Millisecond * 300)

ezbench.SeqPrint()
// sequential benchmark results:
//   1st Seq: took 104.187006ms
//   2nd Seq: took 502.810644ms
//   3rd Seq: took 301.534876ms

ezbench.Seq()
time.Sleep(time.Millisecond * 200)

ezbench.Seq("I'll take long")
time.Sleep(time.Millisecond * 300)
time.Sleep(time.Millisecond * 300)
time.Sleep(time.Millisecond * 300)

ezbench.Seq()
time.Sleep(time.Millisecond * 100)

ezbench.SeqPrint()
// sequential benchmark results:
//   1st Seq: took 100.85786ms
//   I'll take long: took 900.675695ms
//   3rd Seq: took 100.72323ms.
```


### Parallel

```
ezbench.ParStart("long")

ezbench.Par("short", func() { time.Sleep(time.Millisecond * 200) })
time.Sleep(time.Millisecond * 300)

ezbench.ParEnd("long")

ezbench.Par("long", func() { time.Sleep(time.Millisecond * 500) })
ezbench.Par("long", func() { time.Sleep(time.Millisecond * 300) })

ezbench.ParPrint()
// parallel benchmark results:
//   long: called 3 times, took 1.30814s
//   short: called once, took 203.587983ms
```
