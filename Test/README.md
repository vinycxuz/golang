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

### 1.3. Clean Tests
Aqui usamos as regras do triplo A (Arrange, Act, Assert). Arrange é quando configuramos tudo para a execução do teste, ou seja, a construção, inicialização e preparo do objeto que será testado. Ato,
que é quando executamos a função do nosso sistema de teste. E por fim, Assert, que é avaliar o resultado obtido. Os bons casos de testes estão resumidos na sigla FIRST:
- Fast: os testes devem ser rápidos.
- Isolated: Independente da ordem, os resultados devem ser os mesmos.
- Repeatable: os testes devem ser executados em qualquer ambiente e obter os mesmos resultados.
- Self-Validating: os tests devem nos informar se foi aprovado ou não. 
- Timely: Devemos seguir a abordagem TDD, significa que o teste deve ser escrito no tempo certo, que é antes do código de produção. 

Algo importante a se falar sobre os testes, são os números mágicos. Números mágicos são literais espalhadas sem explicação. Exemplo:

    func showMessage() {
        for i := 0; i < 33; i++ {
            fmt.Println("Hello, World!")
        }
    }

No código acima, tem sentido o número 33? Com certeza não. O ideal é substituir por constantes, na qual podemos dar números significativos, são facilmente
alterados e não há necessidade de comentários poluentes no código.

### 1.4. Funcionalidades do Go para Testes
Existem alguns recursos muitos úteis na linguagem Go. O primeiro deles é o subteste, a função t.Run conseguimos criar uma estrutura de pai-filho, que inclusive é refletido na saída do console.

    go test -v
    === RUN   TestIsPalindrome
    --- PASS: TestIsPalindrome (0.00s)
    === RUN   TestSum
    --- PASS: TestSum (0.00s)
    === RUN   TestNewPerson
    === RUN   TestNewPerson/Idade_positiva
    === RUN   TestNewPerson/Idade_negativa
    --- PASS: TestNewPerson (0.00s)
    --- PASS: TestNewPerson/Idade_positiva (0.00s)
    --- PASS: TestNewPerson/Idade_negativa (0.00s)
    PASS
    ok      awesomeProject/Test     0.302s

Outros são os testes de tabela, apresentado na função TestTableNewPerson. Os benefícios 
são:
- Testes mais limpos, pois, há uma seção projetada para coletar todos os casos de testes e outra para processar esses testes.
- Podemos adicionar testes ao nosso conjunto sem duplicar o código.
- Podemos ter subtests num teste primário.

Existem algumas opções para executarmos os testes que valem a pena dizer. Além da forma acima que detalha os testes, incluindo pai e filhos,
utilizando:

    go test -v

Também podemos executar um teste de cobertura de código, utilizando 

    go test -cover
    PASS
    coverage: 100.0% of statements
    ok      awesomeProject/Test     1.208s

Além disso, temos outras duas opções interessantes:

    go test -coverprofile=coverage.out
    go tool cover -html=coverage.out

A primeira gera um arquivo com o relatório de cobertura e a segunda abre esse relatório em um navegador, onde podemos ver detalhadamente quais 
linhas foram cobertas ou não pelos testes. Além disso, existem _tags_ flag para executar apenas testes unitários. Para executar apenas testes unitários, usamos:

    go test -tags=unit

Quando não queremos executar todos os testes e, gostaríamos de selecionar apenas os relevantes para as necessidades atuais. Exemplos situacionais:
- Testes de execução lenta, como testes de integração ou E2E.
- Testes que só podem ser executados em determinados ambientes. 
- Quando queremos excluir testes do pipeline de CI/CD
- Quando queremos excluir testes que dependem de recursos externos.

O Go oferece duas maneiras de pular os testes. Utilizando -short ou com condicionais. o TestSlowFunc ilustra o uso 
de t.Short() para pular o teste quando a flag -short é fornecida. Já o TestSkipFunc ilustra o uso de t.Skip() para pular o teste com base em uma condição.

Mas até agora vimos usando condicionais, o que é massante e pra quem já tem experiência com testes, pode ter estranhado não ter visto assert. O package testify/assert torna
os testes mais legíveis e menos verbosos. Algumas funções úteis:

    assert.Equal(t, expected, actual) // Verifica se dois valores são iguais.
    assert.NotEqual(t, expected, actual) // Verifica se dois valores não são iguais.
    assert.Nil(t, object) // Verifica se um objeto é nulo.
    assert.NotNil(t, object) // Verifica se um objeto não é nulo.
    assert.True(t, condition) // Verifica se uma condição é verdadeira.
    assert.False(t, condition) // Verifica se uma condição é falsa.
    assert.Contains(t, string, substring) // Verifica se uma string contém uma substring específica.
    assert.NotContains(t, string, substring) // Verifica se uma string não contém uma substring específica.

