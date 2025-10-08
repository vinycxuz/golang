#  Introdução ao System Design Interview
System Design é um processo de definição dos componentes e sua integração, APIs e modelos de dados para
construir sistemas de larga escala que atendam um conjunto de requisitos.
Aqui iremos revisar os seguintes conceitos: 

- Introdução com abstrações, requisitos não funcionais e estimativa de recursos
- Building blocks: DNS, Load balancers, databases, key-value, CDN, Sequencer, Monitoring, Distributed cache and messaging queues, pub-sub,
rate limiter, blob store, distributed search, distributed logging, distributed task, shared counters.
- Design Problems

## 1. Abstrações
Abstração é o processo de mostrar o funcionamento sem nos concentrar nos detalhes da implantação. É como imaginar um desenho 
de um pássaro sem se preocupar com as penas, cores ou o tipo de pássaro. Nosso objetivo é transmitir as ideias principais.

### 1.1. Remote procedure calls (RPC)
RPCs fornecem uma abstração para uma chamada de procedimento local aos desenvolvedores, ocultando a complexidade de empacotar e enviar argumentos de função para um servidor
remoto, receber os valores de retorno e gerenciar quaisquer tentativas de rede.

RPC é um protocolo de comunicação amplamente usado em sistemas distrbuídos, participa da camada de transporte e aplicação no modelo OSI.
De maneira objetivo RPC é um mecanismo que permite que um programa execute uma sub-rotina ou procedimento em outro espaço de endereço (geralmente em outro computador na rede compartilhada) e depois
retorna o resultado ao requisitor. 

São extremamente utilizadas como localização em tempo real, serviços de pesquisa e vídeos como Youtube. O google devenvoleu o gRPC, uma estrutura de código aberto que 
constrói utilizando RPC, sistemas distribuídos eficientes e de alto desempenho.

### 1.2. Modelos de consistência
Modelos de consistência fornecem abstrações para raciocinar sobre a correção de um sistema distribuído realizando leitura, gravações e mutações de dados
simultaneous.

Temos alguns tipos de consistências, eventuais, casuais, sequenciais e fortes. 
- A consistência eventual é o modelo mais fraco, onde não há garantia de que uma leitura retornará o valor mais recente escrito.
Normalmente garante _alta disponibilidade_.Exemplos são DNS ou Cassandra. 
- A consistência casual preserva a ordem das operações causalmente relacionadas. Se uma operação A causa uma operação B, então todos os nós verão A antes de B. Um exemplo são respostas de comentários na internet.
- Consistência sequencial garante que todas as operações sejam vistas na mesma ordem por todos os nós, mas não necessariamente na ordem em que foram emitidas. Um exemplo são editores colaborativos.
- Consistência forte é o modelo mais forte, onde todas as operações de leitura retornam o valor mais recente escrito. Normalmente garante _alta consistência_. A replicação síncrona é um dos principais ingredientes para alcançar essa consistência. 
Exemplos são atualizar senha de uma conta.

### 1.3. Espectro do modelo de falhas
O modelo de falhas nos fornecem uma estrutura para raciocinar sobre o impacto das falhas e maneira de lidar com elas.
- Falha-parada: um nó do sistema para de funcionar peramnentemente.
- Colidir: Um nó para de funcionar silecniosamente, mas, os outros nós não conseguem detectar.
- Falhas de omissão: Um nó falha na execução de uma ação, como enviar ou receber uma mensagem.
- Falhas temporais: Um nó executa uma ação, mas não dentro de um limite de tempo esperado.
- Falha bizantina: Um nó exibe comportamento aleatório, seja por um ataque ou um bug. São os mais difíceis de se lidar.

## 2. Requisitos não funcionais
Requisitos não funcionais são atributos do sistema que não está relacionado com as funcionalidades do sistema. 

## 2.1. Disponibilidade
Disponibilidade é a porcentagem de tempo que um serviço ou infraestrutura está acessível aos clientes e é operado em condições normais.
Para medir a disponibilidade matematicamente, usamos a fórmula:

A (Disponibilidade) = (Tempo total - Tempo de inatividade) / Tempo total

E existe a tabela dos nove noves da possível disponibilidade (pesquisar na internet).

## 2.2. Confiabilidade

                                           