#!/usr/bin/env python

import ctypes

fiblib = ctypes.CDLL("./libfibonacci.so")
fib = fiblib.fib
n = 37
print("The %ith Fibonacci number is %i." % (n, fib(n)))
