<?php
/**
 * @version PHP 7.3.29
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 * $ ./vendor/bin/phpunit GrandePalindromoPrimoTest.php
 */
namespace PHP;

use PHPUnit\Framework\TestCase;
require ("GrandePalindromoPrimo.php");

class GrandePalindromoPrimoTest extends TestCase {

    public function testGrandePalindromoPrimo() {
        $game = new GrandePalindromoPrimo();
        //Test 1: tamanho 9
        $startPoint = "128919"; //indice: 129.079
        $decimais = 9;
        $want = "318272813";
        $got = $game->findPrimePalindromeInPi( $startPoint, $decimais );
        echo "\nTest 1: retornou " . $got . " == esperado: " . $want;
        $this->assertEquals($want, $got);
        //Test 2: tamanho 21
        $startPoint = "140672630100"; //indice: 140.672.630.233
        $decimais = 21;
        $want = "151978145606541879151";
        $got = $game->findPrimePalindromeInPi( $startPoint, $decimais );
        echo "\nTest 2: retornou " . $got . " == esperado: " . $want;
        $this->assertEquals($want, $got);
        //Test 3: tamanho 27
        $startPoint = "0"; //indice: 140.672.630.233
        $decimais = 27;
        $want = "960945763984348936367549069";
        $got = $game->findPrimePalindromeInPi( $startPoint, $decimais );
        echo "\nTest 3: retornou " . $got . " == esperado: " . $want;
        $this->assertEquals($want, $got);
    }
}
