package DataStructures

/*
1. BubbleSort
	Algoitmo de ordenação simples que percorre a lista várias vezes, comparando o proximo elemento com o anterior e trocando de lugar se estiverem na ordem errada.

BubblueSort tem complexidade O(n^2) no pior caso, mas é fácil de implementar e entender.
*/

func BubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func more(v1 int, v2 int) bool {
	return v1 > v2
}

/*
2. InsertionSort
Organiza os dados em uma lista ordenada, inserindo cada novo elemento na posição correta. Funciona semelhante a como organizamos cartas por exemplo.
Complexidade O(n^2) no pior caso, mas é eficiente para listas pequenas ou quase ordenadas.
*/

func InsertionSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && comp(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

/*
3. SelectionSort
Percorre a lista várias vezes e coloca o maior valor na ultima posição e dimunuindo sucessivamente o tamanho da lista a ser ordenada.
Complexidade O(n^2) no pior caso, mas é fácil de implementar e entender.
*/

func SelectionSort(arr []int) {
	size := len(arr)
	var i, j, max int
	for i = 0; i < size; i++ {
		max = 0
		for j = 1; j < size-1-i; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[max], arr[size-1-i] = arr[size-1-i], arr[max]
	}
}

/*
4. MergeSort
MergeSort particiona a lista em várias pequenas listas, ordena todas elas e depois junta todas em uma lista ordenada.
Complexidade O(n log n) no pior caso, e é eficiente para listas grandes.
*/

func MergeSort(arr []int) {
	size := len(arr)
	tempArray := make([]int, size)
	mergeSrt(arr, tempArray, 0, size-1)
}

func mergeSrt(arr []int, tempArray []int, lowerIndex int, upperIndex int) {
	if lowerIndex >= upperIndex {
		return
	}
	middleIndex := (lowerIndex + upperIndex) / 2
	mergeSrt(arr, tempArray, lowerIndex, middleIndex)
	mergeSrt(arr, tempArray, middleIndex+1, upperIndex)
	merge(arr, tempArray, lowerIndex, middleIndex, upperIndex)
}

func merge(arr []int, tempArray []int, lowerIndex int, middleIndex int, upperIndex int) {
	lowerStart := lowerIndex
	lowerEnd := middleIndex
	upperStart := middleIndex + 1
	upperEnd := upperIndex
	count := lowerIndex
	for lowerStart <= lowerEnd && upperStart <= upperEnd {
		if arr[lowerStart] < arr[upperStart] {
			tempArray[count] = arr[lowerStart]
			lowerStart++
		} else {
			tempArray[count] = arr[upperStart]
			upperStart++
		}
		count++
	}
	for lowerStart <= lowerEnd {
		tempArray[count] = arr[lowerStart]
		count++
		lowerStart++
	}
	for upperStart <= upperEnd {
		tempArray[count] = arr[upperStart]
		count++
		upperStart++
	}
	for i := lowerIndex; i <= upperIndex; i++ {
		arr[i] = tempArray[i]
	}
}

/*
5. QuickSort
QuickSort é um algoritmo recursivo. Escolhemos um pivô e, a partir dele, colocamos valores menos à esquerda e valores maiores à direita.
Passo a passo:
- Esoclhemos o pivo
- Criar duas variáveis índices, um de índice inferior e outra de índice superior
- Fazer do índice inferior o segundo elemento do array e do índice superior o último elemento do array
- Aumentar o índice inferior até que o valor seja menor que o pivo
- Diminuir o índice superior até que o valor seja maior que o pivo
- Trocas os valores de índice inferior e índice superior
- Repetir até que o índice superior exceda o valor inferior
- Por fim, trocar o valor do pivo com o valor do índice superior
Complexidade O(n log n) no pior caso, e é eficiente para listas grandes.
*/

func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}

func QuickSort(arr []int) {
	size := len(arr)
	quickSortUtil(arr, 0, size-1)
}

func quickSortUtil(arr []int, lower int, upper int) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper
	for lower < upper {
		for arr[lower] <= pivot && lower < upper {
			lower++
		}
		for arr[upper] > pivot && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(arr, lower, upper)
		}
	}
	swap(arr, start, upper)
	quickSortUtil(arr, start, upper-1)
	quickSortUtil(arr, upper+1, stop)
}

/*
6. Quando usar cada algoritmo
- BubbleSort: quando os dados estão quase ordenados e precisa do mpinimo de execução possível.
- InsertionSort: Quando precisamos de uma ordenação estável e a lista é pequena ou quase ordenada.
- SelectionSort: ...
- MergeSort: Quando a estabilidade é importante e a lista é grande.
- QuickSort: Quando a eficiência é crucial e a lista é grande. Normalmente escolhemos quando os dados são aleatórios.
*/
