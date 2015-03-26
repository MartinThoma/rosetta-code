#!/usr/bin/env python


def fib(n):
    """Calculate the n-th Fibonacci number."""
    if n <= 1:
        return n
    else:
        return fib(n-1) + fib(n-2)

if __name__ == '__main__':
    n = 36
    print("The %ith Fibonacci number is %i." % (n, fib(n)))
