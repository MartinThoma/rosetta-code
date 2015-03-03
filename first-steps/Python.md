## Very first steps in Python

Python is an interpreted, dynamic language. It is shipped with most if not all
Linux distributions.

Python has multiple implementations. The best-known one is Cython.

Python scripts can be executed with Cython like this:

    $ python scriptname.py

and with another implementation called "PyPy" like this:

    $ pypy scriptname.py

## Strengths

Python is strong if you want to quickly solve simple problems (**scripting**),
if you want to execute other programs and "orchestrate" it, but Bash scripts
become too complicated.

Python is also very for **scientific calculations** due to fast and reliable
packages like [numpy](http://www.numpy.org/) and [scipy](http://www.scipy.org/).
These libraries have some parts with are written in Fortran and which are
wrapped to get a Python interface. This makes them astonishingly fast.


## Weaknesses

When it comes to pure, raw computing power it is much slower than C. However,
this assumes that you already know very good how to implement the solution and
get it done with C.