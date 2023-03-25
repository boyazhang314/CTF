# VaultChallenge

## Description

> This super secure vault got dropped and jumbled around the letters

## Break In

We have a Java file `VaultChallenge.java`

```java
import java.util.*;

class VaultChallenge {
  public static void main(String args[]) {
    Scanner input = new Scanner(System.in);
    String pass = input.next();
    if (checkPass(pass)) {
        System.out.println("I'm in! B)");
    } else {
        System.out.println("Access Denied.");
    }
  }

  public boolean checkPass(String password) {
    return password.length() == 32 &&
           password.charAt(0)  == 'd' &&
           password.charAt(29) == '4' &&
           password.charAt(4)  == 'r' &&
           password.charAt(2)  == '5' &&
           password.charAt(23) == 'r' &&
           password.charAt(3)  == 'c' &&
           password.charAt(17) == '4' &&
           password.charAt(1)  == '3' &&
           password.charAt(7)  == 'b' &&
           password.charAt(10) == '_' &&
           password.charAt(5)  == '4' &&
           password.charAt(9)  == '3' &&
           password.charAt(11) == 't' &&
           password.charAt(15) == 'c' &&
           password.charAt(8)  == 'l' &&
           password.charAt(12) == 'H' &&
           password.charAt(20) == 'c' &&
           password.charAt(14) == '_' &&
           password.charAt(6)  == 'm' &&
           password.charAt(24) == '5' &&
           password.charAt(18) == 'r' &&
           password.charAt(13) == '3' &&
           password.charAt(19) == '4' &&
           password.charAt(21) == 'T' &&
           password.charAt(16) == 'H' &&
           password.charAt(27) == '2' &&
           password.charAt(30) == '5' &&
           password.charAt(25) == '_' &&
           password.charAt(22) == '3' &&
           password.charAt(28) == '3' &&
           password.charAt(26) == '1' &&
           password.charAt(31) == '6';
    }
}
```

Looks like the password has been mixed up and the characters are out of order, but using the number in the `charAt` we can piece it back together

## Flag

`retroCTF{d35cr4mbl3_tH3_cH4r4cT3r5_123456}`
