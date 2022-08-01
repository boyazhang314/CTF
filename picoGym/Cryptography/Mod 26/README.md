# Mod 26
**Points: 10**

## Description
Cryptography can be easy, do you know what ROT13 is? \
```cvpbPGS{arkg_gvzr_V'yy_gel_2_ebhaqf_bs_ebg13_MAZyqFQj}```

## Hints
<details>
<summary>Hint 1</summary>

> This can be solved online if you don't want to do it by hand!

</details>

## Solution
ROT13 is a special case of the Caesar cipher, where the alphabet is rotated by 13 places. Each letter in the plaintext is replaced by the letter 13 positions down the alphabet 

a &rarr; n \
b &rarr; o \
⋮ \
z &rarr; m \
A &rarr; N \
B &rarr; O \
⋮ \
Z &rarr; M 

This cypher only affects letters in the alphabet, other characters are unchanged. As there are 26 letters, ROT13 serves as it's own inverse.

As suggested in the question, the flag can be obtained by performing ROT13 on the provided text

This can be done with any online tool, such as https://rot13.com/ 

## Flag
```picoCTF{next_time_I'll_try_2_rounds_of_rot13_ZNMldSDw}```
