Each code sample in this directory has to calculate the 36th Fibonacci number.
To do so, the code has to define a dumb recursive function which works like
this:


```python
def fib(n):
    if n <= 1:
        return n
    else:
        return fib(n-1) + fib(n-2)
```

The answer should be:

    The 36th Fibonacci number is 14930352.

## Results

| #   | Language             | Speed  | LOCs |
| --- | -------------------- | ------ | ---- |
| 1   | Rust                 |   0,5s |    7 |
| 2   | Python               |  11,2s |    5 |