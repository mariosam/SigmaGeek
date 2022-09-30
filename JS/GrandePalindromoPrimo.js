/**
 * @version JAVASCRIPT ECMAScript 6
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 */

import fetch from "node-fetch";

/**
 * Encontre o primeiro primo palíndromo de N dígitos na expansão decimal do π (3,1415…)
 * Carrega trilhoes de dados do PI https://pi.delivery
 */
async function findPrimePalindromeInPi( startPoint = "0", decimal = 0 ) {
    try {
        const resp = await fetch(`https://api.pi.delivery/v1/pi?start=${startPoint}&numberOfDigits=1000`);
        const body = await resp.json();
        let number = body.content;

        //percorrer o decimal do pi
        for ( let i=0; i < number.length-decimal; i++ ) {
            let big = number.substring( i, i+decimal );

            if ( preCheck( big ) ) {
                if ( isPalimdrome( big ) ) {
                    console.log(`Palindrome: ${big}`);

                    if ( isPrime( big ) ) {
                        console.log(`Primo: ${big}`);
                        console.log(`Tamanho: ${decimal}`);
                        console.log(`Chamada de API n.: ${startPoint}`);
                        console.log(`Indice da chave: ${parseInt(startPoint)+i}`);

                        return big;
                    } else {
                        console.log("mas nao eh primo!");
                    }
                }
            }
        } 

        //recursividade para chamada limitada da API.
        return findPrimePalindromeInPi( parseInt(startPoint)+(1000-decimal), decimal )
    } catch ( err ) {
        console.log( err );
    }
}

/**
 * Eliminate all invalid data to be palindrome and prime.
 */
function preCheck( number ) {
	//se o primeiro e ultimo digitos forem diferentes, ja nao eh palindromo
	if ( number[0] != number[(number.length)-1] ) return false;

    //se terminar em 5 eh divisivel por 5 (nao primo) - palindromo possui o ultimo digito igual ao primeiro
    //if ( number[0] == 5 ) return false;

    //primos nao iniciam ou terminam em par [se for palindromo o ultimo digito eh igual ao primeiro]
	if ( number[0]%2 == 0 ) return false;

    //javascript nao suporta numeros acima de 9 quatrilhoes (9.007.199.254.740.991)
    if ( number.length < 16 ) {
        //se for divisível por 2 ou 3 ou 5, não é primo!
        if ( ( parseFloat(number) % 2 ) == 0 || ( parseFloat(number) % 3 ) == 0 || ( parseFloat(number) % 5 ) == 0 ) return false;
    } else {
        let bigNumber = BigInt( number );

        //se for divisível por 2 ou 3 ou 5, não é primo!
        if ( ( bigNumber % 2n ) == 0 || ( bigNumber % 3n ) == 0 || ( bigNumber % 5n ) == 0 ) return false;
    }

	//se a soma dos algorismos forem divisiveis por 3 nao eh primo
    let soma = 0;
    for ( let i = 0; i < number.length; i++ ) {
        soma += number[i];
    }
	if ( soma % 3 == 0 ) return false;

    //nenhum erro encontrado
    return true;
}

/**
 * Check if the number is palindrome.
 */
function isPalimdrome( number ) {
    let reverse = number.split("").reverse().join("");

    return reverse == number;
}

/**
 * Check if the number is prime.
 * https://www.numberempire.com/primenumbers.php
 * Adaptado de: https://www.mentebinaria.com.br/forums/topic/795-uma-otimiza%C3%A7%C3%A3o-interessante-para-testar-se-um-valor-%C3%A9-primo/
 */
function isPrime( number ) {
    //bug do javascript para numeros gigantes.
    if ( number.length > 15 ) {
        let big = BigInt( number );

        return isBigIntPrime( big );
    }

    //so precisamos testar até a raiz quadrada
    let raiz = Math.sqrt(number);

    for ( let i = 3; i <= raiz; i += 2 ) {
        if ( ( number % i ) == 0 ) return false;
    }
    /*opcao alternativa, precisa testar e comparar performance.
    for ( let i = 7n; i <= raiz; i += 6n ) {
        if ( ( number % i ) == 0 ) return false;
        i += 4n;
        if ( ( number % i ) == 0 ) return false;
        i += 2n;
        if ( ( number % i ) == 0 ) return false;
        i += 4n;
        if ( ( number % i ) == 0 ) return false;
        i += 2n;
        if ( ( number % i ) == 0 ) return false;
        i += 4n;
        if ( ( number % i ) == 0 ) return false;
        i += 6n;
        if ( ( number % i ) == 0 ) return false;
        i += 2n;
        if ( ( number % i ) == 0 ) return false;
    }
    */

    return true;
}

//funciona para BigInt mas eh absurdamente lento.
//se conseguir obter a raiz quadrada de um bigint ja ajudaria a reduzir os calculos.
//talvez chamar uma api externa q apenas calcule o numero primo gigante.
function isBigIntPrime(p) {
    for ( let i = 2n; i * i <= p; i++ ) {
        if ( p % i === 0n ) return false;
    }

    return true;
}

/**
 * o teorema de wilson: https://rpm.org.br/cdrpm/37/4.htm
 * diz calcular todos os numeros primos
 * uma possibilidade seria calcular todos os grandes numeros e verificar se existem no decimal pi.
 */
export { findPrimePalindromeInPi }
