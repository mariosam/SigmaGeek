/**
 * @version GO 1.19.1
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 */
package GO

import (
	"io"
	"net/http"
	"strconv"
)

/**
 * Encontre o primeiro primo palíndromo de 9 dígitos na expansão decimal do π (3,1415…)
 */
func findBigPrimePalindrome9() string {
	//carrega dados do https://angio.net com 1 milhao de digitos
	resp, _ := http.Get("https://www.angio.net/pi/digits/pi1000000.txt")
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	number := string(body)

	//percorrer o decimal do pi
	for i := 2; i < len(number)-8; i++ {
		big := number[i : i+9]

		if isPalimdrome(big) {
			intBig, _ := strconv.Atoi(big)
			b := []rune(big)

			if isPrime(intBig, int(b[0]), int(b[len(big)-1])) {
				return big
			}
		}
	}

	return "0"
}

/**
 * Check if the number is palindrome.
 */
func isPalimdrome(number string) bool {
	reverse := []rune(number)
	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}

	return string(reverse) == number
}

/**
 * Check if the number is prime.
 * https://www.calculadora-online.xyz/testar-verificar-numero-primo.php?numero=5
 */
func isPrime(number int, first int, last int) bool {
	//primos nao iniciam ou terminam em par
	if first%2 == 0 || last%2 == 0 {
		return false
	}

	//se for divisivel por 2 ou 3 nao eh primo
	if number%2 == 0 || number%3 == 0 {
		return false
	}

	var raiz int = int(math.Sqrt(float64(number)))
	for i := 5; i <= raiz; i += 4 {
		if (number % i) == 0 {
			return false
		}

		i += 2

		if (number % i) == 0 {
			return false
		}
	}

	return true
}
