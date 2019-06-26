# decimal-slides
Slides and examples for gophercon
- install go present with https://github.com/golang/talks
- info about markup language used in https://godoc.org/golang.org/x/tools/present

# Useful resources

anz-bank decimal github repo:
<https://github.com/anz-bank/decimal>

Very nice math explanation:
<https://stackoverflow.com/questions/1089018/why-cant-decimal-numbers-be-represented-exactly-in-binary>

IBM resource and supplier of test suite used for benchmarking:
<http://speleotrove.com/decimal/>

Arbitrary precision implementations of decimal:
<https://github.com/shopspring/decimal>
<https://github.com/ericlagergren/decimal>

IEEE 754R standard (also known as IEEE 754-2008):
<https://ieeexplore.ieee.org/document/4610935>

# Notes
Notes are in a google doc:
<https://docs.google.com/document/d/1KsEQ_375gaqDkzk-gSzGg9rza0jbjBImIp8ygEe8fXk/edit>

# Benchmarking

run benchmarks with
`go test -bench=. -v` from the `examples/benchmark` directory
or run
`go test ./examples/benchmark/ -bench=. -v` from the `decimal-slides` directory
