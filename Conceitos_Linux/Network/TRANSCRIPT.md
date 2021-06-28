# Redes Linux 

Vamos falar um pouquinho de redes ?

Beleza .... bora
Pra começar se voce tem uma VM linux ajuda muito

Vamos instalar o tcpdump e o wireshark pra usar mais tarde.

Pra falar de redes temos que falar das Camadas OSI, acho que uma grande parte da galera ja viu isso em algum lugar.
Mas vamos dar uma passada rapida pra galera que não teve essa introdução. 

Bom legal conceitos da camada OSI.
Temos de 5 a 7 camadas de redes para comunição entre os computadores certo 

Camada Fisica : 
É literalmente como catalogamos a camada de "hardware" da comunicação, como os cabos de rede que vemos espalhados ou, com sorte, até organizados nos data-centers.
També é considerado os Hubs, Repetidores de sinais e as proprias ondas eletromagneticas do WIFI

Camada de Dados: ( a.k.a Data Link )
Nessa camada encontramos as implementações de Ethernet, switches e bridges 

Camada de Network:
Onde econtramos a implementação de IP 

Camada de Transport:
Onde temos a implementação de TCP e UDP

---- E como estamos falando da implementação TCP/IP as camadas de 
 --- Sessão
  ---- Apresentação 
   --- Aplicação 

São todas colocadas em uma unica camada : Aplicação

Camada de Aplicação:
Onde encontramos as implementações de serviços como:
 -- SSH
 -- DNS
 -- SMTP
 -- HTTP
 -- SNMP

Beleza!  Isso é uma introdução para conversarmos um pouco nessa ultimas 3 camadas que vimos:
 - Network : IP
 - Transport: TCP/UDP
 - Application: SSH / DNS/ HTTP

Na camada de Network ja temos a tradução de um MAC Address para um IP (ex: 192.168.0.1)
E na de Transport temos o TCP e UDP.
Ai caimos na diferença entre TCP e UDP e como funciona.

Quando se abre uma comunicação via TCP voce precisa passar pelo famoso ( Handshake ou Three-way handshake ) 
Como isso acontece:
o Servidor1 q inicia a comunicação manda uma mensagem para a outra maquina dizendo "quero falar com voce", para q isso aconteça o Servidor1 precisa setar um flag ( SYN ) nessa mensagem pra dizer q esta iniciando a conversa.
O Servidor2 recebe esta mensagem com a Flag de inicioa de conversa, e responde com uma flag (SYN-ACK) dizendo que recebeu a mensagem do Servidor1 de volta ao Servidor1
O Servidor1 responde com outra Flag (ACk)  dizendo que recebeu a resposta do Servidor2 
E assim se inicia uam Sessão TCP

Da pra ver nessa explicação que claramente um é muito menor q o outro.
 


UDP é um pouco diferente disso, quando voce comunica com um outro servidor usando UDP, simplesmente se cria o pacote e envia sem necessidade de iniciar uma sessão ou de comunicação previa, \
e também não se espera por confirmaçãop de recebimento ou nada do genero. Isso faz com que a comunição entre os servidores seja muito mais leve e simplificada. Mas também tem seus pros e contras.

- Um poco sobre o conteudo dos pacotes TCP e UDP

Conteudo de um pacote TCP:

Source Port : Ou porta de onde se inicia a comunicação do servidor de envio. 16 Bits
Destination Port: Ou porta do servidor de Destino. 16 Bits
Sequence Number: indica a sequencia de dos segmentos de dados que estão sendo enviados: 32 Bits
ACK Number: Numer que define quem é o proximo Sequence Number: 32 Bits
Data Offset: Indica o tamanho do Header TCP: 4 Bits
Reserved: um bit reservado para usos especificos onde existem varias implementações fora das normas comuns.Normalmente é setado com ozero : 3 bits
Flags: Como falamos antes pode ser um SYN, ACK, FIN, RST e normalmente é definido com um bit ligado para a flag q sera usada.
Checksum: que literamelmente é um checksum do proprio conteudo 
Urgent Pointer: Normalmente configuirado como 1 ( ligado ), comumente usado quebras de conexões me limpeza da pilha TCP e também para que o pacote seja entrege mais rapidamente à aplicação.
Options: São opcoes adicionais que se pode colcoar no pacote.

Conteudo de um pacote UDP:

Source Port : Ou porta de onde se inicia a comunicação do servidor de envio. 16 Bits
Destination Port: Ou porta do servidor de Destino. 16 Bits
Length: Tamanho do pacote UDP 
Checksum: que literamelmente é um checksum do proprio conteudo 


Por exemplo:

UDP é otimo para aplicações de stream, como a do Twitch, Netflix ou jogos onde o Handshake do TCP pode causar atrasos e retranmissões de pacote, e isso por sua vez causaria problemas nas montagens dos frames.
No TCP por outro lado, se tem a confirmação de que o servidor que voce quer se comunicar esta recebendo seus pacotes devidamente.
O TCP garante q o conteudo q esta sendo trfegado esta em um ordem e garante menos erros de transmissão, por isso acaba sendo o protocolo comum para muitas implementaçoes de aplicações que conhecemos como HTTP / SSH e etc...

Porra coisa longa e chate neh. Mas vamos ver como isso funciona na Real ?

Vou rodar aqui o exemplo da aula anterior que é uma API em Go bem simples que tem la no GitHub.
E voce pode ir fazendo junto comigo 

- Rodar API
- Abrir tcpdump
  - Comadinho legal de como funciona: tcpdump -i any host $SEU_HOST_AQUI port 8080 -w go_api.pcap ( no caso dessa API q esta rodando na 8080 ) 

- Agora vamos la abrir o wireshark e ver como isso funciona.

Tem um livro muito bom que tem varios testes legais pra voces pegarem esse jeito ai. do NoStarch. e Vai estar nos Links de estudo dessa aula.
Mas se voce acessar esse site aqui  : https://nostarch.com/packetanalysis3 da pra baixar um zip cheio de arquivos pcap do wireshark pra brincar.

Vamos ver alguns ?

- Window size = Segment number + Window Size * window scale
- FIN vs RST 
- Retransmissions




