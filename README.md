[![Build Status](https://travis-ci.org/guesslin/mits.svg)](https://travis-ci.org/guesslin/mits) [![Coverage Status](https://coveralls.io/repos/guesslin/mits/badge.svg)](https://coveralls.io/r/guesslin/mits) [![GoDoc](https://godoc.org/github.com/guesslin/mits?status.svg)](https://godoc.org/github.com/guesslin/mits)

MITS
====

aka Mutual Information Term Segmentation


Formula
=======

MI(x[i], x[i+1]) = p(x[i], x[i+1]) * log(p(x[i], x[i+1]) / (p(x[i]) * p(x[i+1])))

p(x[i]): the probability of x[i] show up in all sentences
p(x[i], x[i+1]) : the probability of x[i] and x[i+1] continue show up in all sentences

Author
======

* Yu-han Lin, guesslin1986@gmail.com
