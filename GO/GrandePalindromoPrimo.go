/**
 * @version GO 1.19.1
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 */
package GO

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
)

/**
 * Encontre o primeiro primo palíndromo de N dígitos na expansão decimal do π (3,1415…)
 * Carrega trilhoes de dados do PI https://pi.delivery
 */
func findPrimePalindromeInPi(startPoint string, decimal int) string {
	fmt.Println("Executando teste com posicao inicial do PI em: " + startPoint + " para " + strconv.Itoa(decimal) + " digitos")
	pos, _ := strconv.Atoi(startPoint)

	resp, _ := http.Get("https://api.pi.delivery/v1/pi?start=" + startPoint + "&numberOfDigits=1000")
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	number := string(body)

	//remove os caracteres json
	number = number[12 : len(number)-3]

	//percorrer o decimal do pi
	for i := 0; i < len(number)-decimal; i++ {
		big := number[i : i+decimal]

		if preCheck(big) {
			if isPalimdrome(big) {
				fmt.Println("Palindrome: ", big)

				if isPrime(big) {
					fmt.Println("Primo: ", big)
					fmt.Println("Tamanho: ", decimal)
					fmt.Println("Chamada de API n.: ", startPoint)
					fmt.Println("Indice da chave: ", pos+i)

					return big
				} else {
					fmt.Println("mas nao eh primo!")
				}
			}
		}
	}

	//verificar fim do arquivo de trilhoes para interromper looping.
	if strings.Contains(startPoint, "99999999999") {
		return "EOF"
	}

	//recursividade para chamada limitada da API.
	return findPrimePalindromeInPi(strconv.Itoa(pos+(1000-decimal)), decimal)
}

/**
 * Eliminate all invalid data to be palindrome and prime.
 */
func preCheck(number string) bool {
	//se o primeiro e ultimo digitos forem diferentes, ja nao eh palindromo
	if number[0] != number[len(number)-1] {
		return false
	}

	//primos nao iniciam ou terminam em par [se for palindromo o ultimo digito eh igual ao primeiro]
	b := []rune(number)
	if int(b[0])%2 == 0 {
		return false
	}

	//se for divisível por 2 ou 3 ou 5, não é primo!
	if len(number) < 16 {
		check, _ := strconv.Atoi(number)
		//check, _ := strconv.ParseInt(number, 10, 64)

		if check%2 == 0 || check%3 == 0 || check%5 == 0 {
			return false
		}
	}

	//se a soma dos algorismos forem divisiveis por 3 nao eh primo
	soma := 0
	for i := 0; i < len(number); i++ {
		digito, _ := strconv.Atoi(string(number[i]))

		soma += digito
	}
	return soma%3 != 0

	//Todo número palíndromo com uma quantidade par de algarismo é divisível por 11

	//nenhum erro encontrado
	//return true
}

/**
 * Check if the number is palindrome.
 */
func isPalimdrome(number string) bool {
	//inverte a string para comparar com original
	reverse := []rune(number)
	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}

	return string(reverse) == number
}

/**
 * Check if the number is prime.
 * Adaptado de: https://www.mentebinaria.com.br/forums/topic/795-uma-otimiza%C3%A7%C3%A3o-interessante-para-testar-se-um-valor-%C3%A9-primo/
 */
func isPrime(n string) bool {
	if len(n) > 15 {
		return isOnlinePrime(n)
	}

	number, _ := strconv.ParseInt(n, 10, 64)

	//so precisamos testar até a raiz quadrada
	var raiz int64 = int64(math.Sqrt(float64(number)))
	var i int64

	for i = 3; i <= raiz; i += 2 {
		if int64(number)%i == 0 {
			return false
		}
	}

	return true
}

/**
 * solucao temporaria para verificar se palidromos acima de 15 digitos sao primos
 * https://www.numberempire.com/primenumbers.php
 */
func isOnlinePrime(number string) bool {
	resp, _ := http.Get("https://www.numberempire.com/primenumbers.php?number=" + number)
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	html := string(body)

	//verifica o retorno
	check := strings.Index(html, number+" is a prime")

	return check != -1
}
