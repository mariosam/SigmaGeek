/** 
 * @version JAVA
 * @author MARIO SAM <eu@mariosam.com.br>
 * @see I would love to work with you instead solving web code tests: hire me!
 * $ mvn clean test -Dtest=your.package.TestClassName
 */ 
package JAVA;

import static org.junit.Assert.assertEquals;
import org.junit.Test;

public class GrandePalindromoPrimoTest {

        @Test
	public void testGrandePalindromoPrimo() throws Exception {
            //Test 1
            String want = "318272813";
            int decimais = 9;
            String startPoint = "128919"; //indice: 129.079
            String got = GrandePalindromoPrimo.findPrimePalindromeInPi( startPoint, decimais );
            assertEquals(want, got);
            //Test 2
            want = "151978145606541879151";
            decimais = 21;
            startPoint = "140672630100"; //indice: 140.672.630.233
            got = GrandePalindromoPrimo.findPrimePalindromeInPi( startPoint, decimais );
            assertEquals(want, got);
	}

}
