# Reverse

## Description

> Try reversing this file? Can ya?
>
> I forgot the password to this [file](https://artifacts.picoctf.net/c/270/ret). Please find it for me?

## Undo

We have a file `ret` which is an executable

<figure><img src="../../.gitbook/assets/image (3).png" alt=""><figcaption></figcaption></figure>

Let's open this executable in [Ghidra](https://ghidra-sre.org/), and we find this interesting tidbit in the `main` function

<figure><img src="../../.gitbook/assets/Untitled (3).png" alt=""><figcaption></figcaption></figure>

## Flag

`picoCTF{3lf_r3v3r5ing_succe55ful_3f1331e7}`
