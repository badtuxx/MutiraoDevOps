
```bash
sudo tcpdump -i any port 8080 and host 192.168.0.1

sudo tcpdump -i any port 8080 and host 192.168.0.1 -A

sudo tcpdump -i any port 8080 and host 192.168.0.1 -w arquivo.pcap

sudo tcpdump -i lo port 8080 -A 
```

```bash
telnet 192.168.86.12 8080
```

```bash
nmap -Pn -p 8080 192.168.86.12
```

```bash
nc -uz 192.168.86.12 8080
nc -vuz 192.168.86.12 8080
```

Pra "escutar" em uma porta usando Netcat

```bash
nc -l -k 9999
```
