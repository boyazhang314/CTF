# Safe Opener 2

## Description

> What can you do with this file?
>
> I forgot the key to my safe but this [file](https://artifacts.picoctf.net/c/290/SafeOpener.class) is supposed to help me with retrieving the lost key. Can you help me unlock my safe?

## Hints

<details>

<summary>1</summary>

Download and try to decompile the file.

</details>

## Find the Key

We have a Java file `SafeOpener.class` and though we can try decompiling it, we have a quicker way of searching with `strings SafeOpener.class`

<figure><img src="../../.gitbook/assets/image (41).png" alt=""><figcaption></figcaption></figure>

## Flag

`picoCTF{SAf3_0p3n3rr_y0u_solv3d_it_b601de44}`
