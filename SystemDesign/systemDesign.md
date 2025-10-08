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
Confiabilidade é a probabilidade do serviço desempenhar as suas funções por um período específico.
O cálculo que utilizamos:

    MTBF (Tempo médio entre falhas) = (Tempo total decorrido - Soma do tempo de inatividade) / número total de falhas

    MTTR (Tempo médio para reparo) = Tempo total de manutenção / número total de reparos
                                           
Vale ressaltar que confiabilidade e disponibilidade são conceitos diferentes. Um sistema pode estar disponível 90% do tempo mas é confiável em apenas 80% do tempo.
Confiabilidade é medida de execução do sistema, enquanto disponibilidade é uma medida de acessibilidade do sistema.
Um ótimo exemplo é imaginar que o aplicativo pode estar funcionando corretamente, mas o servidor caiu, isso é disponibilidade. 

## 2.3. Escalabilidade
Escalabilidade é a capacidade de um sistema lidar com a quantidade crescente de carga de trabalho sem comprometer o desempenho. 
Existem dois tipos de escalabilidade, vertical e horizontal.
- Escalabilidade vertical refere-se ao fornecimento de recursos adicionais a um dispositivo existente (RAM, CPU...)
- Escalabilidade horizontal é o número de máquinas na rede, utilizamos commodity hardware para os atrativos de custo-beneficio

## 2.4. Maintainability
Maintainability(a palavra em português manutenibilidade é lamentável rs) é a probabilidade de um sistema restaurar as suas funções num tempo especificado após a ocorrência da falha.
Podemos dividir o conceito em três aspectos subjacentes:
- Operabilidade: Facilidade de garantir o bom funcionamento
- Lucidez: simplicidade do código. 
- Modificabilidade: Facilidade de implementar mudanças

MMTR (tempo médio de reparo) = Tempo total de manutenção / número total de reparos

## 2.5. Tolerância a falhas
Tolerância a falhas é a capacidade de um sistema de executar persistentemente mesmo se um ou mais de seus componentes falharem. As duas qualidades essencias que tornam toelrante a falhas é
disponibilidade e confiabilidade. Existem algumas técnicas para alcançar a tolerância a falhas:
- Replicação: manter várias cópias dos dados em diferentes máquinas. É um dos mais utilizados.
- Ponto de verificação: é uma técnica que salva o estado do sistema em um armazenamento estável para recuperação posterior em caso de falhas. É como imaginar as versões de um programa, podemos
voltar para uma versão anterior se algo der errado.

# 4. Blocos de construção
Problema de projetos se sistemas geralmente apresentam semelhanças, e nosso objetivo é separar blocos de construção para discutir detalhadamente 
o design apenas uma vez. Esses blocos de construção podem ser reutilizados em vários projetos.

## 4.1. DNS
DNS é o serviço de nomenclatura da internet. Ele traduz o nome de um domínio para um endereço de IP. Tipos comuns de registro de recursos DNS são:
- UM: Fornece o mapeamento do nome do host para um endereço IP. Exemplo tipo-nome-valor(A, relay1.main.something.io,104.18.2.119)
- NS: Fornece o nome do host que é o DNS autoritativo para um nome de domínio. Exemplo tipo-nome-valor(NS, something.io, ns1.something.io)
- CNAME: Fornece o mapeamento do nome do host para outro nome de host. Exemplo tipo-nome-valor(CNAME, www.something.io, servidor1.something.io)
- MX: Fornece o mapeamento do nome do host para um servidor de email. Exemplo tipo-nome-valor(MX, something.io, mail.something.io)
- TXT: Fornece o mapeamento do nome do host para um texto arbitrário. Exemplo tipo-nome-valor(TXT, something.io, "v=spf1 include:_spf.google.com ~all")

Cache: DNS utiliza cache em várias camadas para melhorar a latência.

Hierarquia: DNS é hierárquico, com servidores raiz no topo, seguidos por servidores de domínio de nível superior (TLD) e servidores autoritativos.
Funciona mais ou menos assim: Temos no topo o Root, após o TLD, que é o .com, .net, .oi... e por ultimo o nome autoritário do domínio, que é o google.com, something.io...

O DNS é um próprio sistema distribuído, e a sua natureza tem algumas vantagens:
- Evita se tornar um ponto único de falha (SPOF)
- Ele atinge baixa latência, pois os servidores DNS são distribuídos globalmente.
- Ele obtém mais flexibilidade durante manutenções e atualizações.

Existem 13 tipos de servidores DNS, nomeados de A a M com muitas instâncias espalhadas pelo mundo. Suas características o tornam escalável, confiável, consistente e disponível.3

## 4.2. Load Balancers
Load Balancers são usados para dividir de forma justa todas as solicitações dos clientes entre o conjunto de servidores disponíveis. Após o firewall, o load balancer é o primeiro ponto de contato para todas as solicitações dos clientes.
Eles oferecenm:
- Escalabilidade: Ao adicionar servidores, a capacidade do sistema aumenta.
- Disponibilidade: Mesmo que alguns servidores fiquem inativos ou sofram falhas, o sistema permanece.
- Desempenho: Os load balancers podem encaminhar solicitações para os servidores com menor carga.

Além disso, oferecem alguns serviços essenciais, como:
- Verificação de integridade: LB usam o protocolo heartbeat para monitorar a integridade.
- Término de TLS/SSL
- Análise preditiva: LB podem analisar o tráfego para identificar padrões e otimizar o desempenho.
- Intervenção humana reduzida: LB podem automatizar tarefas como balanceamento de carga e failover.
- Descoberta de serviço: As solicitações dos clientes são encaminhadas para servidores de hospedagem apropriados por meio de consulta ao registro de serviço.
- Segurança: Melhorar a segurnaça mitigando ataques DDos.

Quando um LB falha, o sistema pode ficar indiponível, para isso, normalmente as empresas utilizam clusters de LB.

Existem LB em escala local e em escala global. 

#### 4.2.1. Detalhes avançados de Load Balancers
Os LB distribuem as solicitações dos clientes de acordo com um algoritmo, alguns deles são:
- Round-robin scheduling: Cada servidor recebe uma solicitação por vez.
- Weighted round-robin: Cada nó recebe um peso e os LB encaminham as solicitações dos clientes de acordo com o peso do nó.
- Least Connections: Atribuir solicitações a servidores com menos conexões existentes.
- Least response time
- IP hash
- URL hash

## 7. Content Delivery Network (CDN)
CDN é um grupo de servidores proxy distribuídos geograficamente. Um servidor proxy é um servidor intermediário entre um cliente e um servidor de origem.
A CDN visa reduzir a latência, pois, se esforçam para ter largura de banda suficiente disponível no caminho e aproximar os usuários.

Uma CDN armazena principalmente dois tipos de dados: estáticos (tipo de dado que não muda com frequência e permanece nos servidores por um longo período, como imagens, vídeos) e dinâmicos
(tipo de dado que muda com frequência, como páginas da web personalizadas, feeds de notícias).

Os requisitos funcionais de uma CDN são recuperar, solicitar, entregar, pesquisar, atualizar e excluir.
Os requisitos não-funcionais são desempenho, disponibilidade, escalabilidade e confiabilidade.

Os componentes básicos de uma CDN são:
- clientes
- sistema de roteamento: determina o servidor CDN mais próximo do cliente. Ele entende aonde o conteúdo é colocado, quantas solicitações são feitas
a carga quee um conjunto específico de servidores está manipulando e o namespace URI(Uniform Resource Indetifier)