## 1. Testes

Os testes são fundamentais para o desenvolvimento de software. Com eles conseguimos detectar bugs, garantir a satisfação do cliente,
documentar, segurança, entre outros benefícios. 

Existem dois termos que frequentemente vemos ao lidar com testes: código de produção e código de teste. 
- Código de produção: é o código que implementa a funcionalidade do sistema, ou seja, o código que será executado em um ambiente de produção.
- Código de teste: é o código que não será enviado ao sistema final mas é usado para garantir que o código de produção funcione como esperado, se tornando inclusive
um bloco de construção no ciclo de desenvolvimento.

### 1.1. Tipos de Testes
Existem vários tipos de testes:
- Testes unitários: testam um único compoonente do código, como uma função ou um método.
- Testes de integração: testam o comportamento de um ou mias componentes em conjunto, como uma interação no banco de dados por exemplo.
- Testes end-to-end(E2E): abrangem todo o processo. No caso de uma API por exemplo, eles executam uma solicitação HTTP real e confirmam a resposta HTTP fornecida pelo servidor HTTP.
- Teste de UI: testes de interface.

Aqui falaremos especificamente de testes unitários, pois, são o pilar central do desenvolvimento orientado a testes.

Cobertura de código é uma unidade de medida que nos indica a quantidade de código coberto por testes. As instruções que escrevemos em nossos códigos de produção, podem ser agrupadas em três 
categorias: syntax lines, logic lines e branches. Usando o famoso exemplo do kata FizzBuzz.

    func FizzBuzz(input int) string {
	if input%5 == 0 && input%3 == 0 {
		return "FizzBuzz"
	}
	if input%5 == 0 {
		return "Buzz"
	}
	if input%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(input)
    }

Categorizando as declarações acima, temos:
- As linhas de sintaxe não contêm lógica de négocios, podem declarar funções, variáveis e estruturas.
- As linhas de lógica contêm a lógica de negócios, como expressões condicionais e loops.
- As branches são os caminhos que o código pode seguir, como o "if" acima que pode ser verdadeiro ou falso.

Existem também coberturas de código, os principais são:
- Cobertura de funções: é a razão entre o número de funções invocadas em testes e o número total de funções definidas no código de produção.
No exemplo acima, temos apenas uma função, portanto, precisamos apenas de um teste que a invoque para atingir 100% de cobertura.
- Cobertura de instruções: é a proporção do número de instruções cobertas pelos testes em relação ao número total delas no código de produção.
- Cobertura de branches: é a medição em um grupo de linhas.

### 1.2. Test-Driven Development (TDD)
O desenvolvimento orientado baseia-se em um estilo iterativo. Esses processo envolve 3 fases:
- Escreva um teste reprovado
- Faça o teste reprovado passar
- Refatore o código 

Essa abordagem deve ser seguida sempre que quisermos adicionar novas funcionalidades ao nosso programa. Também é importante falarmos sobre as suas restrições.
- Não temos permissão para escrever nenhum teste de código de produção, a menos que seja para fazer um teste unidade com falha passar.
- Não temos permissão para escrever mais de um teste de unidade do que o suficiente para falhar.
- Não temos permissão para escrever mais código de produção do que o suficiente.

Tudo começa escrevendo um teste com falha, já que não podemos escrever nenhum código de produção a menos que não tenhamos um teste para ele. Em seguida, escrevemos o código suficiente 
para o teste passar. Finalmente, podemos refatorar o código para melhorar sua estrutura sem alterar seu comportamento.

    o desenvolvimento orientado a teste não se trata de teste, se trata de design.

Imaginemos a abordagem tradicional, se quisermos criar uma feature qualquer, enquanto estamos na fase de design, podemos começar a desenvolver antes do previsto.
Ao aplicar o desenvolvimento orientado a testes, podemos ter vários benefícios como confiança em produzir códigos, pois, estamos cobertos de testes; melhorar os padrões de código por
conta da refatoração; sistema fracamente acoplado, que depende mais de abstrações que detalhes concretos de implementação, o que permite atender os requisitos do cliente e mudar rapidamente a
direção do projeto. As situações em que o TDD não deve ser aplicados são em testes lentos, requisitos poucos claros e UI/UX.



