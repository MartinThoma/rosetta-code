#!/usr/bin/env python

"""Sort an array of structs.

Python knows C-structs (see https://docs.python.org/2/library/struct.html),
but usually you use dictionaries.
"""

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

by_axis = [p['name'] for p in sorted(planets, key=lambda n: n['axis'])]
print("by axis:  %s" % str(by_axis))

by_name = [p['name'] for p in sorted(planets, key=lambda n: n['name'])]
print("by name: %s" % str(by_name))
