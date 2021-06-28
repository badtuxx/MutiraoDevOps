# Strace

```bash
sudo strace -fp 2983531
sudo strace -f go run main.go 
sudo strace -e trace=read,stat -f go run main.go 
sudo strace -e trace=lseek -f python seek.py 
sudo strace -e trace=clone -f go run main.go 
```

# Perf

```bash
sudo  perf record -a -g sleep 10
sudo  perf report --sort comm,dso
```

# Ltrace

```bash
ltrace -c -f go run main.go 
```
