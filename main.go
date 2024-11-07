package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1000
const qnt = 100000

var falhas = 0

var (
	saldo   []float32
	valores []float32
	compras [][]int
	vendas  [][]int
	relacao [][]int
	err     error
)

// Função que lê um arquivo e converte o conteúdo em um vetor de inteiros
func leVetor(filename string) ([]float32, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Leitura do arquivo linha por linha
	scanner := bufio.NewScanner(file)
	var numbers []float32

	for scanner.Scan() {
		line := scanner.Text()

		// Divide a linha em partes separadas por espaço
		values := strings.Split(line, " ")

		// Converte cada valor em inteiro e adiciona ao vetor
		for _, v := range values {
			if v != "" { // Ignora espaços em branco
				num, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				numbers = append(numbers, float32(num))
				// Se já alcançamos o tamanho N, interrompemos
				if len(numbers) == N {
					break
				}
			}
		}
		if len(numbers) == N {
			break
		}
	}

	// Verifica se houve erro na leitura
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

// Função para ler uma matriz N x N de um arquivo
func leMatrizN(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Cria um scanner para ler o arquivo
	scanner := bufio.NewScanner(file)

	// Determina o número de linhas (N) no arquivo
	var N int
	var lines []string

	// Lê todas as linhas do arquivo
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Verifica se o arquivo tem pelo menos uma linha
	if len(lines) == 0 {
		return nil, fmt.Errorf("arquivo vazio")
	}

	// Assume-se que a primeira linha do arquivo indica o número de colunas
	N = len(strings.Fields(lines[0]))

	// Cria a matriz N x N
	matrix := make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, N)
	}

	// Preenche a matriz com os valores do arquivo
	for i, line := range lines {
		if i >= N {
			break
		}
		values := strings.Fields(line)
		for j, v := range values {
			if j >= N {
				break
			}
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			matrix[i][j] = num
		}
	}

	return matrix, nil
}

// Função para ler uma matriz N x 2 de um arquivo
func leMatrizqnt(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Cria um scanner para ler o arquivo
	scanner := bufio.NewScanner(file)

	// Determina o número de linhas (N) no arquivo
	var N int
	var lines []string

	// Lê todas as linhas do arquivo
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Verifica se o arquivo tem pelo menos uma linha
	if len(lines) == 0 {
		return nil, fmt.Errorf("arquivo vazio")
	}

	// Define N como o número de linhas no arquivo
	N = len(lines)

	// Cria a matriz N x 2
	matrix := make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, 2) // Define 2 colunas para cada linha
	}

	// Preenche a matriz com os valores do arquivo
	for i, line := range lines {
		if i >= N {
			break
		}
		values := strings.Fields(line)
		if len(values) < 2 {
			return nil, fmt.Errorf("linha %d do arquivo não contém pelo menos 2 valores", i+1)
		}

		// Converte os dois valores para inteiro e os adiciona à matriz
		for j := 0; j < 2; j++ {
			num, err := strconv.Atoi(values[j])
			if err != nil {
				return nil, err
			}
			matrix[i][j] = num
		}
	}

	return matrix, nil
}

func somaVetor(vetor []float32) float32 {
	var soma float32 = 0

	for _, valor := range vetor {
		soma += valor
	}
	return soma
}

func atribuiVetores() {
	saldo, err = leVetor("saldos.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	valores, err = leVetor("valores.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	compras, err = leMatrizqnt("compras.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	vendas, err = leMatrizqnt("vendas.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	relacao, err = leMatrizN("relacao.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
}

func main() {
	gera()
	atribuiVetores()
	sequencial()
	fmt.Println(float32(falhas) / float32(qnt))
}
