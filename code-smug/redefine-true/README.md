Try to reassign the 'true' value to false.

Your code should looke like this:

```python

True = False

if True:
    print("YourLanguage does not 'support' reassignment of True :-)")
else:
    print("YourLanguage 'supports' reassignment of True :-(")
```

## Results

| Language   | Result | Comments                    |
| ---------- | ------ | --------------------------- |
| C++        | Good   |                             |
| Python 2.7 | Bad    | PEP8 checkers will warn you |
| Python 3.X | Good   |                             |