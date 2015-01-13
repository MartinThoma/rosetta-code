#!/usr/bin/env python

""" Sort an array of 10,000,000 randomly created integers."""


import random
import time

# Create the array - well, actually a list
n = 10000000
max_int = 1000000000
array = [random.randint(0, max_int) for i in range(n)]

# Sort it and measure time
t0 = time.time()
array_sorted = sorted(array)
t1 = time.time()

# Output the time
print("%0.2f seconds for sorting %i integers." % (t1-t0, n))
