#!/usr/bin/env python


def fib(n):
    """Calculate the n-th Fibonacci number.

    Parameters
    ----------
    n : int

    Returns
    -------
    int : The n-th Fiboacci number.
    """
    if n <= 1:
        return n
    else:
        return fib(n-1) + fib(n-2)

if __name__ == '__main__':
    n = 37
    print("The %ith Fibonacci number is %i." % (n, fib(n)))
