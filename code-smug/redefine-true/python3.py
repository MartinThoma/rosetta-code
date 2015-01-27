#!/usr/bin/env python3

import sys

True = False


if True:
    print("Python %s does not 'support' reassignment of True :-)" %
          sys.version)
else:
    print("Python %s 'supports' reassignment of True :-(" % sys.version)
