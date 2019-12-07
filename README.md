# punit
A pretty-printing library for units in Go.


```go
// ...
// blah, blah, blah...
// me computer, me do thing
// eventually: have scalar for human
x := time.Duration(6276663872206468)
fmt.Printf("%s\n")
// output: 1743h31m3.872206468s
```
**Absolutely outrageous.** That looks like garbage. How am I supposed to read
it? Here's my fix. I realized soon after implementing it generalizes to all
kinds of units, not just units of time, so the following snippet demonstrates
how to initialize the scale of units that I used to express human-readable
timespans (in the interest of helping you express human-readable everything
else).

```go
// ...blah, blah, blah...
// find another very ugly scalar
// Rob Pike ain't got sh*t on _my_ human
// this time will be different

var timeScale = ScaleOf(Scale{
	{int64(time.Hour*24*365), "year"},
	{int64(time.Hour*24*30),  "month"},
	{int64(time.Hour*24*7),   "week"},
	{int64(time.Hour*24),     "day"},
	{int64(time.Hour),        "hour"},
	{int64(time.Minute),      "minute"},
	{int64(time.Second),      "second"},
}...)

x := time.Duration(626663872206468)

timeScale.Format(x, time.Hour*24, true)
// output: 2 months, 1 week, and 5.6 days
timeScale.Format(x, time.Hour*24, false)
// output: 72.6 days
```

Boom. Perfect. Flawless. Just as flawless as my code, and almost as flawless as
my visage. The only thing you won't like is your benchmarks.
