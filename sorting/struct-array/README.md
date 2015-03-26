Each code sample in this directory has to define 3 structs (or whatever is
commonly used to store data in the language, e.g. classes or dictionaries).

The structs are planets with the attributes `name`, `aphelion`, `perihelion`,
`axis` and `radius`.

The container with those three structs should get printed:

* unsorted
* sorted by axis
* sorted by name


## Example

```python
mars = {'name': "Mars",
        'aphelion': 249.2,
        'perihelion': 206.7,
        'axis': 227939100,
        'radius': 3389.5}

earth = {'name': "Earth",
         'aphelion': 151.930,
         'perihelion': 147.095,
         'axis': 149598261,
         'radius': 6371.0}

venus = {'name': "Venus",
         'aphelion': 108.939,
         'perihelion': 107.477,
         'axis': 108208000,
         'radius': 6051.8}

planets = [mars, earth, venus]
print("unsorted: %s" % str(planets))
print("by axis:  %s" % str(sorted(planets, key=lambda n: n['axis'])))
print("by name: %s" % % str(sorted(planets, key=lambda n: n['name'])))
```

Output:

```bash
unsorted: [{'aphelion': 249.2, 'axis': 227939100, 'name': 'Mars', 'perihelion': 206.7, 'radius': 3389.5}, {'aphelion': 151.93, 'axis': 149598261, 'name': 'Earth', 'perihelion': 147.095, 'radius': 6371.0}, {'aphelion': 108.939, 'axis': 108208000, 'name': 'Venus', 'perihelion': 107.477, 'radius': 6051.8}]
by axis:  ['Venus', 'Earth', 'Mars']
by name: ['Earth', 'Mars', 'Venus']
```


## Results

Only the lines used for sorting are counted.

| #   | Language             | LOCs |
| --- | -------------------- | ---- |
| 1   | Python               |    3 |
| 2   | Go                   |   11 |