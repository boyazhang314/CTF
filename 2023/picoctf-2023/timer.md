# timer

## Description

> You will find the flag after analysing this apk
>
> Download [here](https://artifacts.picoctf.net/c/449/timer.apk).

## Hints

<details>

<summary>1</summary>

Decompile

</details>

<details>

<summary>2</summary>

mobsf or jadx

</details>

## Tick Tock

We have a file `timer.apk`, which is an Android Package Kit

As suggested by the hints, we can use `jadx` to open and decompile. Once it's done, we can search through the files for "picoCTF"

<figure><img src="../../.gitbook/assets/Untitled (2).png" alt=""><figcaption></figcaption></figure>

## Flag

`picoCTF{t1m3r_r3v3rs3d_succ355fully_17496}`
