# vitamix: time virtualization for Go

_vitamix_ is a source-to-source compiler, coupled with a runtime logic, for "virtualizing time"
in Go programs. It is intended to aid in the testing of and experimentation with
extremely time-sensitive "control" software, which is itself written in Go.

## Problem statement

Control software (like robotic control or congestion control) is software that in one
form or another listens to external events in real time and reacts with carefully-timed
responses.

Control software is notoriously hard to get right, because even the time to
compute a response (which depends on the specific implementation and the
hardware) can and does affect the overall outcomes in a non-negligible manner.

In developing control software, one wants to be able to perform
sandbox simulations, while meeting the following demands:

* _The computational time (i.e. the time it takes the software to compute a reaction to an event) is zero._ 
Being able to do this is a way of decoupling the outcomes of the response logic
(whose timing is of central importance) from the delays incurred by
implementation and hardware peculiarities.

* _Execution outcomes are deterministically exactly reproducible, both in value and in timing._ 
When writing tests for control software, the testing framework often needs to
check whether responses occur at desired times. This is hard to achieve while
using real time since even the outside temperature can affect the execution speed
of your code.

* _Executions that are long-running in real time complete instantaneously in the sandbox._
More often than not, we like to simulate a control algorithm for a long time
(e.g. minutes) to verify that its behavior is stable over time.  A long-running
control program often sleeps most of the time, while waking up at sparse
moments to respond to some external event. In a sandbox, we would really like
to be able to execute such a long-running program without waiting through the
times when the software is sleeping.

Beyond being able to skip over sleeping time and being able to remove the effect of 
your own code's run time, one can envision additional time-related needs like:

* _Simulate different and varying CPU speeds:_ This applies e.g. if one wants to
test whether the control software behaves accurately in extreme situation when
it is starved out of CPU resources by other running processes (e.g. a web
server experiencing a denial-of-service attack).

In summary, we would like to be able to execute _any_ Go program, while
eliminating 

_vitamix_ accommodates precisely these needs: _vitamix is a tool that
allows one to execute_

## Concept

## Principle of operation

_vitamix_ consists of two parts:

* Source-to-source transformer which rewrites the program source, and
* Go package which is utilized by the rewritten source and acts as
an additional runtime (additional to the Go runtime).



To install,

	$ go get github.com/petar/vitamix

## About

_vitamix_ is authored and maintained by [Petar Maymounkov](http://pdos.csail.mit.edu/~petar/). 
