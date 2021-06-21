
Sniffer em uma porta especifica
```bash
sudo tcpdump -i any port 8080
```

Em uma porta E em um host especifico, também se pode usar `OR`
```bash
sudo tcpdump -i any port 8080 and host 192.168.0.1 -A
```

Porta e Host e lendo o conteudo dos pacotes
```bash
sudo tcpdump -i any port 8080 and host 192.168.0.1 -A
```

Porta e Host fazendo o output para um arquivo PCAP ( Compativel com Wireshark ) 
```bash
sudo tcpdump -i any port 8080 and host 192.168.0.1 -w arquivo.pcap
```

Especificando o Network Card que será usado o Sniffer ( neste caso localhost ) 
```bash
sudo tcpdump -i lo port 8080 -A 
```

```bash
telnet 192.168.86.12 8080
```

NMAP costuma ser uma ferrament muito mais completa para testes de rede.
Para fãs de GUI  se pode usar o [Zenmap](https://nmap.org/zenmap/)
```bash
nmap -Pn -p 8080 192.168.86.12
```

Netcat `-u` promete testar conexões UDPs ( com resálvas ) 
```bash
nc -uz 192.168.86.12 8080
nc -vuz 192.168.86.12 8080
```

Pra "escutar" em uma porta usando Netcat

```bash
nc -l -k 9999
```
