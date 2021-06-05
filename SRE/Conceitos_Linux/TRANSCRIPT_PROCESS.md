- Process
O que é um processo:
É qualquer programa q esteja rodando em uma maquina.
Shell é um processo. 

- Como o Shell funciona:
O shell é um processo que intepreta o que o usuário escreve e traduz isso para o computador, devolvendo um output

Não da pra falar de Processos sem falar sobre o PID. Afinal, o que é o PID.
PID é um numero unico identificador de um programa que esta rodando ou seja, um processo. Conhecido como Process ID.
O PID é criado e gerenciado pelo Kernel.
Todo systema operacional hoje possuem um limite de numeros que podem ser criados.

- Estrutura de um Processo.
Todo programa no momento que é criado, vira um processo e tem seu conteudo "alocado" em estruturas tais como:
- Stack
- Heap
- Data Segment
- Code Segment

Cada processo possue estruturas unicas e proprias.
Quando o codigo é executado vira um processo e essas estruturas são alocadas.

Mas como se cria um processo.
- Normalmente com uma chamada para o kernel chamada de exec() ou fork()
Exec como o nome diz é executar
E fork ?
Fork em portugues é garfo, agora imagina formato de um garfo, e é literalmente isso , tem uma base unica que se "divide" outras 3 our 4 bases, dependendo do seu garfo claro.
O que acontece neste caso é, o fork ira criar um processo clone do processo atual requesitando suas proprias estruturas e alocando seus proprios recursos.
Mas nem sempre isso é verdade, pode acontecer de recursos serem compartilhados entre o processo pai e processo filho 

Assim quando se faz o fork de um processo voce pode vizualizar o PID e PPID, o que foi comentado na aula inicia do Jeferson sobre os processos pai e filhos. ( Vamos falar um pouco mais disso daqui a pouco ) 

Também pode ser usado a syscall por exemplo exec().
O exec literalmente faz o que diz, executa o programa iniciando um processo novo alocando todos os atributos necessarios comentados antes.

Bom, e ae o processo ta rodando mas o que acontece agora?
Agora temos a gerencia dos estados desse processo, digamos que este processo precisa de abrir um socket ou carregar algo em memoria. 
O que acontece:
O processo passa pelo estado de NEW ( quando ele é criado )
Apos isso ele entra em READY, onde vai iniciar a execução isso significa que ele esta pronto para ser executado com todas as suas estruturas ja endereçadas
Mas como pode existir um outro processo usando o recurso que ele precisa ou que ele depende de um processo para que ele possa executar. 
Este processo entra em WAITING.
Se ele não tem nenguma dependencia o processo recebe uma bandeira verde do kernel dizendo que pode ser executado 
Assim entra em RUNNING.
Se o processo executou tudo que precisava ele entra no estado de TERMINATED.
Caso ainda exista algo que o processo precisa executar, chamando outros recursos, ele volta para READY.
Se os recursos estão liberados ele entra em RUNNING novamente ou em WAITING caso tenha dependencia, reiniciando o ciclo.


Legal, ja entendemos um pouquinho de como um processo funciona e onde ele é carregado , vamos ver um pouco de como é na vida real.

Vamos dar uma olhada no `ls`:

Pra quem não conhece eu vou usar uma ferramenta chamada strace que vai me ajudar a ler o que o processo esta executando em quais Syscall.

------ Quem sabe o que são Syscall ?
Vamos dar uma passadinha em Syscall? Sim ou Não ? 

Sim >>> 

Syscalls:

Syscall são interfaces de kernel que o processos executam para se comunicar com o Hardware.
Alguns exemplos falamos antes como o exec() e fork() 

strace hostnane

strace uname

Hostname é bem estranho pq no final acaba chamando o uname.

Não >>>>

strace -f ls
 - aqui na primeira linha voce ve a chamada do exec(), que na implemntação do kernel do linux, ( incluindo outros tipos de exec ) este é o execve(), 
onde inicia o programa e aloca os recursos para o processo ser executado
 - voce pode ver aqui diversas outras Syscalls. "Porra Giordano, mas voce complicou com essas syscalls ai."
Bem, eu não imagino q voce consegue manter na cabeça todas as syscalls que existem. Mas tem mta ferramenta de pesquisa inclusive o proprio `man`
Esta e a forma que o processo se comunica com o Kernel para que isso possa ser "trackeado" pelo kernel fazendo a gerencia do processo.

"Caralho, muita informação cara ! Onde eu vou usar isso? "

Ai que tá a jogada, vamos fazer um hands on :
- começando com o `ls`, beleza facil pequeno, e ja tem syscall pra porra neh! sim mas muito disso não vai ser seu interesse. Agora que voce sabe do exec() e do fork(), vamos ver o write()
- Como o linux é um sistema de arquivos, o write é bem usado, olha aqui  -> write() escrevendo no console do seu shell os arquivos que o `ls` achou. Vamos agora usar um truque de filtro:
- strace -e trace=write() -f ls , simplificou a brincadeira neh.

"Ta mas como isso me ajuda?"

- Imagina um processo rodando que voce não tem acesso ao source code pra precisa entender o q ele esta fazendo ou onde esta tendo erro, onde ta parando. Voce não precisa ler o codigo.
Vamos tentar com alguns programinhas q eu criei de exemplo.

- exemplo do Seek()
- exemplo do Stat()
- agora vamos ver o fork()
- por ultimo uma api simples que podemos fazer as chamadas e ver o que esta acontecendo dentro do processo mesmo sem ter o codigo fonte.

"Mas agora como eu sei o q eu devo filtrar?" 

Vamos rodar um `strace -c -f PID`  e ver quais são as syscalls mais usadas, onde tem mais erro e onde vale a pena olhar.

agora vamos filtrar por `connect`, aqui da pra ver facil onde tem erro  de conexão com cliente ou servidor.
E na aula de Network podemos ver como juntar essa ferramenta com outras ferramentas de rede para entender a comunicação dos processos pela rede.
