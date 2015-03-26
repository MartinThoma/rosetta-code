#!/usr/bin/env python

import os


def main():
    script_path = os.path.realpath(__file__)
    print(script_path)
    print(os.path.dirname(script_path))

if __name__ == '__main__':
    main()
