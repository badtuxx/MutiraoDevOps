#!/usr/bin/env python3

import time
import mmap

with mmap.mmap(-1, 64) as mm:
    beer = input("Digite a cerveja favorita do Giordano: ")
    mm.write(str.encode(beer))

    print("Vou printar a cerveja favorita do Giordano para todo o sempre...")
    while True:
        mm.seek(0)
        print(mm.readline().decode("utf-8").strip())
        time.sleep(1)
