/**
 * @version JAVA
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 */
package JAVA;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;

public class GrandePalindromoPrimo {

    public static void main(String[] args) {
        System.out.printf("Palindrome: %d\n", findPrimePalindromeInPi("128919", 9));
    }

    /**
     * Encontre o primeiro primo palíndromo de N dígitos na expansão decimal do π (3,1415…)
     * Carrega trilhoes de dados do PI https://pi.delivery
     */
    public static String findPrimePalindromeInPi(String startPoint, int decimal) {
        System.out.println("Executando teste com posicao inicial do PI em: " + startPoint + " para " + decimal + " digitos");

        try {
            URL url = new URL("https://api.pi.delivery/v1/pi?start="+startPoint+"&numberOfDigits=1000");
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("GET");

            BufferedReader in = new BufferedReader(new InputStreamReader( conn.getInputStream() ));
            String number = in.readLine();
            in.close();

            //remove os caracteres json
            number = number.substring(12, number.length()-2);

            //percorrer o decimal do pi
            for ( int i=0; i < number.length()-decimal; i++ ) {
                String big = number.substring( i, i+decimal );

                if ( preCheck( big ) ) {
                    if ( isPalimdrome( big ) ) {
                        System.out.println("Palindrome: "+big);

                        if ( isPrime( big ) ) {
                            System.out.println("Primo: "+big);
                            System.out.println("Tamanho: "+decimal);
                            System.out.println("Chamada de API n.: "+startPoint);
                            System.out.println("Indice da chave: "+(Long.parseLong(startPoint)+i));

                            return big;
                        } else {
                            System.out.println("mas nao eh primo!");
                        }
                    }
                }
            } 
        } catch ( Exception err ) {
            System.out.println( err );
            return "erro";
        }

        //verificar fim do arquivo de trilhoes para interromper looping.
        if ( startPoint.contains("99999999999") ) {
            return "EOF";
        }

        //recursividade para chamada limitada da API.
        return findPrimePalindromeInPi( String.valueOf(Long.parseLong(startPoint)+(1000-decimal)), decimal );
    }

    /**
     * Eliminate all invalid data to be palindrome and prime.
     */
    private static boolean preCheck(String number) {
        //se o primeiro e ultimo digitos forem diferentes, ja nao eh palindromo
        if ( number.charAt(0) != number.charAt(number.length()-1) ) return false;

        //primos nao iniciam ou terminam em par [se for palindromo o ultimo digito eh igual ao primeiro]
        if ( number.charAt(0)%2 == 0 ) return false;

        /** nao consegui fazer a divisao para numeros gigantes bitinteger (?) */
        if ( number.length() < 16 ) {
            //se for divisível por 2 ou 3 ou 5, não é primo!
            long check = Long.parseLong( number );

            if ( check % 2 == 0 || check % 3 == 0 || check % 5 == 0 ) return false;
        }

        //se a soma dos algorismos forem divisiveis por 3 nao eh primo
        int soma = 0;
        for ( int i = 0; i < number.length(); i++ ) {
            soma += number.charAt(i);
        }
        if ( soma % 3 == 0 ) return false;

        //nenhum erro encontrado
        return true;
    }

    /**
     * Check if the number is palindrome.
     */
    private static boolean isPalimdrome(String number) {
        StringBuilder reverse = new StringBuilder();
        reverse.append( number );
        reverse.reverse();

        return number.contentEquals( reverse );
    }

    /**
     * Check if the number is prime.
     * Adaptado de: https://www.mentebinaria.com.br/forums/topic/795-uma-otimiza%C3%A7%C3%A3o-interessante-para-testar-se-um-valor-%C3%A9-primo/
     */
    private static boolean isPrime(String number) throws Exception {
        //problemas com numeros muito grandes
        if ( number.length() > 15 ) {
            return isBigPrime( number );
        }

        //so precisamos testar até a raiz quadrada
        int raiz = (int) Math.sqrt(Integer.parseInt(number));
    
        for ( int i = 3; i <= raiz; i += 2 ) {
            if ( ( Integer.parseInt(number) % i ) == 0 ) return false;
        }
    
        return true;
    }

    /**
     * solucao temporaria para verificar se palidromos acima de 15 digitos sao primos
     * https://www.numberempire.com/primenumbers.php
     */
    private static boolean isBigPrime(String number) throws Exception {
        URL u = new URL("https://www.numberempire.com/primenumbers.php?number="+number);
        HttpURLConnection c = (HttpURLConnection) u.openConnection();
        c.setRequestMethod("GET");

        BufferedReader input = new BufferedReader(new InputStreamReader( c.getInputStream() ));
        String html = input.readLine();
        input.close();
        
        //verifica o retorno
        return html.contains(number+" is a prime");

        //BigInteger number = new BigInteger(n);
        //long raiz = number.sqrt().longValue();
        //for ( int i = 3; i <= raiz; i += 2 ) {
            //if ( BigInteger.valueOf(i).mod(number) == BigInteger.valueOf(0) ) return false;
        //}
        //return true;
    }
}
