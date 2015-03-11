Try to divide two integer values and see what happens

Your code should looke like this:

```python
a = 1
b = 2
c = a/b

if c > 0.1:
    print("YourLanguage does implicit conversion of integers to floats")
else:
    print("YourLanguage does integer division")
```

## Results

| Language   | Integer Division | Warning | Comments             |
| ---------- | ---------------- | ------- | -------------------- |
| Go         | Yes              | No      |                    |
