<?php
/**
 * @version PHP 7.3.29
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 */
namespace PHP;

class GrandePalindromoPrimo {    

    /**
     * Encontre o primeiro primo palíndromo de N dígitos na expansão decimal do π (3,1415…)
     * Carrega trilhoes de dados do PI https://pi.delivery
     */
    public function findPrimePalindromeInPi( $startPoint = "0", $decimal = 0 ): string {
        try {
            $resp = file_get_contents("https://api.pi.delivery/v1/pi?start=".$startPoint."&numberOfDigits=1000");
            $body = json_decode($resp);
            $number = $body->content;
    
            //percorrer o decimal do pi
            for ( $i=0; $i < strlen( $number )-$decimal; $i++ ) {
                $big = substr( $number, $i, $decimal );
    
                if ( $this->preCheck( $big ) ) {
                    if ( $this->isPalimdrome( $big ) ) {
                        echo "\nPalindrome: ".$big;
    
                        if ( $this->isPrime( $big ) ) {
                            echo "\nPrimo: ".$big;
                            echo "\nTamanho: ".$decimal;
                            echo "\nChamada de API n.: ".$startPoint;
                            echo "\nIndice da chave: ".($startPoint+$i);
    
                            return $big;
                        } else {
                            echo " - mas nao eh primo!";
                        }
                    }
                }
            } 
    
            //recursividade para chamada limitada da API.
            return $this->findPrimePalindromeInPi( $startPoint+(1000-$decimal), $decimal );
        } catch (Exception $err) {
            echo '\nExceção capturada: ', $err->getMessage();
        }
    }

    /**
     * Eliminate all invalid data to be palindrome and prime.
     */
    private function preCheck( $number ): bool {
        //se o primeiro e ultimo digitos forem diferentes, ja nao eh palindromo
        if ( $number[0] != $number[strlen($number)-1] ) return false;

        //primos nao iniciam ou terminam em par [se for palindromo o ultimo digito eh igual ao primeiro]
        if ( $number[0]%2 == 0 ) return false;

        //se for divisível por 2 ou 3 ou 5, não é primo!
        if ( ( $number % 2 ) == 0 || ( $number % 3 ) == 0 || ( $number % 5 ) == 0 ) return false;

        //se a soma dos algorismos forem divisiveis por 3 nao eh primo
        $soma = 0;
        for ( $i = 0; $i < strlen($number); $i++ ) {
            $soma += $number[$i];
        }
        if ( $soma % 3 == 0 ) return false;

        //nenhum erro encontrado
        return true;
    }

    /**
     * Check if the number is palindrome.
     */
    private function isPalimdrome( $number ): bool {
        $reverse = strrev( $number );

        return $reverse === $number;
    }

    /**
     * Check if the number is prime.
     */
    private function isPrime( $number ): bool {
        //so precisamos testar até a raiz quadrada
        $raiz = sqrt($number);

        /* tem q melhorar essa funcao
        quando tem 21 digitos leva 25 minutos calculando */
        for ( $i = 3; $i <= $raiz; $i += 2 ) {
            if ( intdiv( intval($number), $i ) == 0 ) return false;
        }

        return true;
    }
}
