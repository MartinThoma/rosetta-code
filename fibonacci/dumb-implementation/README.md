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
| 1   | Rust                 |  1,22s |   22 |
| 2   | Python+Rust          |  1,23s | 20+9 |
| 3   | Python               | 17,29s |   23 |