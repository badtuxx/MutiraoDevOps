# Linux Page Cache

# Cache = Arquivos
# Buffer = Metadados

apt-get install vmtouch time

--------------------------------

free -wh

sysctl -w vm.drop_caches=3

free -wh

head -c 1G </dev/urandom > escrita

free -wh

vmtouch -v escrita

grep Dirty /proc/mem/info

sync

grep Dirty /proc/mem/info

/usr/bin/time -v wc -l escrita

vmtouch -v escrita

-------------------

free -wh

sysctl -w vm.drop_caches=3

free -wh

head -c 1G </dev/urandom > leitura

sysctl -w vm.drop_caches=3

free -wh

/usr/bin/time -v wc -l leitura

vmtouch -v leitura

free -wh

/usr/bin/time -v wc -l leitura

vmtouch -v leitura

------------------------------

