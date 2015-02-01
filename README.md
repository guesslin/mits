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
