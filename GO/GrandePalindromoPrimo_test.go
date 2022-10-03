/**
 * @version GO 1.19.1
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 * $ go test -timeout 999999s -run TestGrandePalindromoPrimo
 */
package GO

import "testing"

func TestGrandePalindromoPrimo(t *testing.T) {
	tables := []struct {
		want       string
		decimal    int
		startPoint string
	}{
		{"318272813", 9, "128919"},                    //indice: 129.079
		{"151978145606541879151", 21, "140672630100"}, //indice: 140.672.630.233
	}

	for _, table := range tables {
		got := findPrimePalindromeInPi(table.startPoint, table.decimal)

		if got != table.want {
			t.Errorf("Waiting for this %s but the return was this: %s", table.want, got)
		}
	}
}
