#Decimal
Title: something like "learning go through implementing a library"

- Started with go in December 2018
- Never worked on Go before, or even open source

- The ending of Foundations of algorithms
    - Introduction to floating point data types before
- The problems with floating point data types
    - Numbers and stuff and how 10 isn't a power of 2
    - Repeated fractions `1/3 --> 0.3333333`
        - Example; `0.1 + 0.1 + 0.1 != 0.3` (in float)
- Real life example of bad code -> aviation and counting in 1/10th of a second
    - How under many "normal" circumstances floating point errors will cancel out

- Solution
    - Base 10 counting system instead of base 2
        - Example; `0.1 + 0.1 + 0.1 == 0.3` (in decimal)
    - IEEE 754R standard released in 2008
    - How this solves the 1/10th fraction issue
    - All base 2 numbers are represented in base 10
- Go Goal; 128 Bit decimal floating point
    - Go doesn't support 128 bit integers
    - Concat of two 64 bit integers
- Fixed precision decimal vs Floating point benchmark
    -
