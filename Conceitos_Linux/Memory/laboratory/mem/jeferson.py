#!/usr/bin/env python3

import sys

pid = input("Digite o PID (ID do Processo) do Giordano: ")
mem_file = open(f'/proc/{pid}/mem', "rb+")

address = int(input("Digite o endereço de memória da cerveja favorita do Giordano: "), 16)
mem_file.seek(address)

beer = input("Digite a cerveja favorita do Jeferson: ") + "\n"
mem_file.write(str.encode(beer))
mem_file.flush()

print("Tudo pronto, agora você já pode ir ver a cerveja favorita do Giordano!")
