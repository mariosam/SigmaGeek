/**
 * @version JAVASCRIPT ECMAScript 6 
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 * $ node NomeDaClasse.test
 */
const TEST = require('tape')
const OBJ = require('./GrandePalindromoPrimo.js')

TEST('Starting test...', (t) => {
    //Test 1: tamanho 9
    let startPoint = "120" //posicao: 129.079 - 120 (arquivo) x 1000 (por aquivo) = 120.000
    let decimais = 9
    t.assert(OBJ.findBigPrimePalindrome( startPoint, decimais ) === "318272813", "Expect: 318272813")
    //Test 2: tamanho 21
    startPoint = "140670000" //posicao: 140.672.630.233
    decimais = 21
    t.assert(OBJ.findBigPrimePalindrome( startPoint, decimais ) === "151978145606541879151", "Expect: 151978145606541879151")

    t.end()
})
