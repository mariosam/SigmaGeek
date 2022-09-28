/**
 * @version GO 1.19.1
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 * $ go test
 */
package GO

import "testing"

func TestGrandePalindromoPrimo9(t *testing.T) {
	tables := []struct {
		want string
	}{
		{"318272813"}, //posicao: 129.079
	}

	for _, table := range tables {
		got := findBigPrimePalindrome9()

		if got != table.want {
			t.Errorf("Waiting for this %s but the return was this: %s", table.want, got)
		}
	}
}
