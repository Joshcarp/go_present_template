# decimal-slides
LINK TO PRESENTATION: <http://joshcarp.xyz:8080/content/decimal.slide#1>

Slides and examples for gophercon
- go present is used:
- install https://github.com/golang/talks
- godoc https://godoc.org/golang.org/x/tools/present

# Useful resources

anz-bank decimal github repo:
<https://github.com/anz-bank/decimal>

Very nice math explanation:
<https://stackoverflow.com/questions/1089018/why-cant-decimal-numbers-be-represented-exactly-in-binary>

Great blog on floating point:
<https://ciechanow.ski/exposing-floating-point/>

IBM resource and supplier of test suite used for benchmarking:
<http://speleotrove.com/decimal/>

Arbitrary precision implementations of decimal:
<https://github.com/shopspring/decimal>
<https://github.com/ericlagergren/decimal>

IEEE 754R standard (also known as IEEE 754-2008):
<https://ieeexplore.ieee.org/document/4610935>

# Other decimal libraries cited
<https://github.com/ericlagergren/decimal>

<https://github.com/shopspring/decimal>


# Notes
Notes are in a google doc:
<https://docs.google.com/document/d/1KsEQ_375gaqDkzk-gSzGg9rza0jbjBImIp8ygEe8fXk/edit>

# Benchmarking

run benchmarks with
`go test -bench=. -v` from the `examples/benchmark` directory
or run
`go test ./examples/benchmark/ -bench=. -v` from the `decimal-slides` directory
