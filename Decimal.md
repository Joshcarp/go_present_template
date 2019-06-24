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
The problem goes; Computers count in base 2, so it's simply quick and easy to represent all components of a floating point number in base 2, although this is quick and easy, math gets  in the way.
- The first standard of this was introduced in 1985 and was the first time where all programming languages could speak a common language when it came to non-integer values.
- It had its fundamental problems, like the inability to represent decimal fractions
    - Just like 1/3 can't be represented as a decimal number, `1/10` can't be represented in a binary one.
    - In some cases this doesn't matter that much; any inaccuracies might cancel out; `1/10 * 10 == 1`
    - But in some other cases it won't and the problem is compounded;
        - Take an example of a system which counts every 10'th of a second
        - To start with this may not be that much of a problem as the error might be tiny, but if we've got a system which relies on precise calculations, we can very quickly run into problems.
        - for example; we've got some code which counts in 1/10th of a second, over 10 hours or so the actual time considered by the system is vastly off from the time we're expecting
        - This was actually the case for an anti missile defense system which had a system time which was 0.34 seconds out of sync, This error caused it to miss its target by around 500m.
        - This could've been avoided many ways; one of them being using a decimal data type instead.
So what's the solution?
    - Instead of storing the number in base 2, we can store it as base 10, and our infamous problem is solved: `0.1 + 0.1 + 0.1 == 0.3`
I started working on this project, and I was possibly the most confused i've been in a long, long time.
I had the objective of implementing a the 754 2008 revision of the standard in golang, and that's where the journey began.
I didn't start from complete scratch; I was given quite a substantial code base already, and for someone who's biggest project in programming was probably copying code into unity and making a game it was a huge task.
So I started\







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
