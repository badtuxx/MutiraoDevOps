
#  Mutirão DevOps

 
O objetivo desse repositório é reunir a maior quantidade de comandos específicos de Linux para troubleshooting do sistema.

##  Checagem do SO

Mostra as informações de CPU:

    lscpu

Checagem dos núcleos de processador físicos e lógicos:

    cat /proc/cpuinfo | grep processor | wc -l




##  Sistema de arquivos

**Diretórios:**

Volta um nível de diretório:

  

    cd \..\

  

Dois níveis e assim por diante:

  

    cd \..\..\

  

Entrar diretamente no home do usuário corrente:

  

    cd ~

  

##  Boot

  

##  Redes
Checagem de listening de portas

    ss -antpu

##  Internet


Comando para checar o seu IP público:
  
    curl -s https://checkip.amazonaws.com

  
Mostrando os IPs privados:

    hostname -I
