# YATD

## Description

> Yet another cipher to decode:
>
> ```
> FH4E7LC@Ecf0243hbfh37_ge35fcN
> ```

## Cipher Find

As implied by the challenge description, we are given a cipher to decrypt

Considering the length, there is a good chance this string directly maps to a flag in some way, so let's try comparing the first 5 characters `FH4E7` with `uwctf`

* F: 70 -> u: 117
* H: 72 -> w: 119
* 4: 52 -> c: 99
* E: 69 -> t: 116
* 7: 55 -> f: 102

We can see that each time there is an offset of 47 in the ASCII representations of the characters

This implies ROT47

{% embed url="https://www.browserling.com/tools/rot47" %}

## Flag

`uwctf{rot47_acb9379bf086bd74}`
