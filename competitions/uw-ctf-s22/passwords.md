# Passwords

## Description

> Charlie the C programmer forgot his password that will make the program print his flag! He also accidentally deleted his original source file, and may have made a mistake writing it in the first place...
>
> Bonus chall: figure out what the mistake Charlie made was :)
>
> ``[`a.out`](https://ctf.csclub.uwaterloo.ca/uploads?key=fb1a8a046c906259a1252bb53ee350f25457a3d3e145d45fa93a9708997abdc1%2Fa.out)``

## Come Out

We're giving nothing more than an executable. Let's get some information about it

<figure><img src="../../.gitbook/assets/image (11).png" alt=""><figcaption></figcaption></figure>

So it's an executable. As expected. Let's try running it

<figure><img src="../../.gitbook/assets/image (5).png" alt=""><figcaption></figcaption></figure>

It wants a password, as stated in the description, but we do not know it

Here, let's try running it through a debugger such as GDB with `gdb ./a.out`

<figure><img src="../../.gitbook/assets/image (2).png" alt=""><figcaption></figcaption></figure>

Here we can see all of the functions. Let's put a breakpoint on `main` with `b main`

## Flag

`uwctf{xkcdgoodpasswd_4676d47590b4876e}`
