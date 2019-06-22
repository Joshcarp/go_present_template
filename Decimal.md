# Title
How I learned go through writing a library

# Abstract
Go is known as an easy language to get started with, but how easy is it for a first year uni student to get the grasp of?
I had finished two units of programming when I got the opportunity to work on an open source go project; Implementing a decimal datatype in go.
In this talk I will be talking about how I went from not even knowing what Go was to developing an implementation of a fixed precision datatype in go, and my journey throughout.
As the infamous example of the failings of floating point numbers go; `0.1 + 0.1 + 0.1 != 0.3`
I was studying for my algorithms exam when I learned this, it didn't make much sense to me at the time;
- "Is this such a big problem, and if so, why is such a widely used "number" so fundamentally flawed?"
And so a couple of months go by and I get the opportunity to do a summer internship, and I start researching.
The problem goes; Computers count in base 2, so it's simply quick and easy to represent all aspects of a floating poiont number in base 2






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
- Go Goal; Implement a 64 bit decimal floating point library
    - Conform as closely to standard as practical
- Journey starts:
    - First off; Git and CI (not too much here)
    - My plan "Write some test cases, Copy the code, change some 64's to 128's and done!"
    - Test cases and all
        - I had never used unit tests before
            - "Think of all conceivable edge cases"
            - Don't really know what this means, but start writing tests
    - Find IBM test suite online; _containing more than 81,300 tests_
        - A lot more than I could probably think of in a summer internship





- 128 Bit decimal floating point
    - Go doesn't support 128 bit integers
    - Concat of two 64 bit integers
- Fixed precision decimal vs Floating point benchmark
    -
