/**
 * @version JAVASCRIPT ECMAScript 6 
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 * $ node NomeDaClasse.test
 */
import TEST from 'tape'
import { findPrimePalindromeInPi } from './GrandePalindromoPrimo.js'

TEST('Starting test...', (t) => {
    //Test 1: tamanho 9
    let startPoint = "128919" //indice: 129.079
    let decimais = 9
    t.assert(findPrimePalindromeInPi( startPoint, decimais ) === "318272813", "Expect: 318272813")
    //Test 2: tamanho 21
    startPoint = "140672630100" //indice: 140.672.630.233
    decimais = 21
    t.assert(findPrimePalindromeInPi( startPoint, decimais ) === "151978145606541879151", "Expect: 151978145606541879151")

    t.end()
})
