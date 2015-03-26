Each code sample in this directory has to calculate the 15th Fibonacci number.
To do so, the code has to define a dumb recursive function which works like
this:


```python
def fib(n):
    if n <= 1:
        return n
    else:
        return fib(n-1) + fib(n-2)
```