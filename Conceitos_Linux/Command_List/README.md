#  Mutirão DevOps

 
O objetivo desse repositório é reunir a maior quantidade de comandos específicos de Linux para troubleshooting do sistema.

  
(*Usem o StackEdit:* https://stackedit.io/app#)

  
##  Checagem do SO

  

##  Sistema de arquivos

  

**Diretórios:**

Volta um nível de diretório:

    cd \..\
  
Dois níveis e assim por diante:

    cd \..\..\

  Entrar diretamente no home do usuário corrente:
 

    cd ~
  

##  Boot

Checa o tempo médio de execução do boot:

    systemd-analyze


##  Redes

Comando para checar o seu IP público:
  
    curl -s https://checkip.amazonaws.com

  
Mostrando os IPs privados:

    hostname -I

Checa se esta saindo para internet

    ping 8.8.8.8

Checa se o DNS está resolvendo

    ping google.com

Mostra todos os soquetes TCP

    ss -t -a


Mostra todos os soquetes UDP

    ss -u -a

Mostra todas as conexões SMTP estabelecidas

ss -o state established '( dport = :smtp or sport = :smtp )'

Mostra todas as conexões estabelecidas em HTTP

    ss -o state established '( dport = :http or sport = :http )'

Mostra todos os processo locais conectados ao X Server

    ss -x src /tmp/.X11-unix/*


Mostra todos os soquetes TCP em estado FIN-WAIT-1:

Lista todos os soquetes TCP sockets em estado FIN-WAIT-1 para nossa rede httpd 200.52.1/24 e olha seus temporizadores.  

    ss -o state fin-wait-1 '( sport = :http or sport = :https )' dst 200.52.1/24

Como filtrar soquetes usando estados TCP:
 
Para TCP com IPV4,

    ss -4 state NOME-FILTRO

Para TCP com IPV6,

    ss -6 state NOME-FILTRO
