package Test

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

/*
1. Testes
	Testes acredito que estejamos familiarizados. Em Go, testes são feitos com o sufixo _test.go. No arquivo principal, criamos as funções normalmente.
No arquivo de testes, utilizamos o package testing e criamos as funções sempre iniciando com o nome Test seguido do nome da função a ser testada.
Toda função recebe um ponteito do tipo *Testing.T.
	Em GO, é muito comum estilo de testes table-driven, onde criamos uma table com os valores de entrada e os valores esperados (exemplo no arquivo test_test.go).
A saída dos testes é executada com o comando go test, e uma saída com falha não afeta toda a execução dos testes, t.Errorf não gera um panic ou interrompe a execução.
Exemplo de saída com erro:

--- FAIL: TestIsPalindrome (0.00s)
    test_test.go:24: IsPalindrome("anA") = false; want true
FAIL
exit status 1
FAIL    awesomeProject/Test     1.065s

	Em Golang, os testes são minimalistas, não funciona como outras linguagens que oferecem frameworks de testes robustos, como JUnit, PyTest, etc.
A ausência de identificadores de teste, hooks para executar operações de setup e teardown, comparação de objetos complexos, entre outros, pode ser vista como uma limitação.
Porém, essa 'atitude' de Go é intencional, pois, ela espera que os autores dos testes façam esse trabalho por conta própria, mantendo a simplicidade e a leveza da linguagem.
Um bom teste exibe clareza e é sucinto no sintoma do problema.

No intuito de nos ajudar (se alguém tiver vendo esse material), vamos fazer um overview sobre os testes no geral e, como estou adotando o .md para isso, vou seguir por lá.
*/
