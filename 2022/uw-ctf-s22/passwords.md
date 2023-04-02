# Passwords

## Description

> Charlie the C programmer forgot his password that will make the program print his flag! He also accidentally deleted his original source file, and may have made a mistake writing it in the first place...
>
> Bonus chall: figure out what the mistake Charlie made was :)
>
> ``[`a.out`](https://ctf.csclub.uwaterloo.ca/uploads?key=fb1a8a046c906259a1252bb53ee350f25457a3d3e145d45fa93a9708997abdc1%2Fa.out)``

## Come Out

We're giving nothing more than an executable. Let's get some information about it

<figure><img src="../../.gitbook/assets/image (11) (1) (1).png" alt=""><figcaption></figcaption></figure>

So it's an executable. As expected. Let's try running it

<figure><img src="../../.gitbook/assets/image (5) (1) (2) (1).png" alt=""><figcaption><p>It never stops running...</p></figcaption></figure>

It wants a password, as stated in the description, but we do not know it

Here, let's try running it through a debugger such as GDB with `gdb ./a.out`

<figure><img src="../../.gitbook/assets/image (2) (2) (1).png" alt=""><figcaption><p>info functions</p></figcaption></figure>

We can see all of the functions. Let's put a breakpoint on `main` with `b main`. Then we can run the program with command `run` until it hits the breakpoint

<figure><img src="../../.gitbook/assets/image (9) (2).png" alt=""><figcaption></figcaption></figure>

The flag has to be stored somewhere, so let's try dumping out the assembler code with `disas/s` and get some juicy information.&#x20;

{% embed url="https://visualgdb.com/gdbreference/commands/disassemble" %}

What's really interesting is the call the `strcmp` followed by tons of `putchar` calls

<figure><img src="../../.gitbook/assets/image (6) (1) (3).png" alt=""><figcaption></figcaption></figure>

It's likely the code here is calling `strcmp` to compare the user input to the password, and if the password is correct it puts, or prints out, the flag. These hex values that are being moved with `mov` look like ASCII characters.

This is the fun part.

We painstakingly copy each of these hex values and convert them to text.

For the curious, this is the whole concatenated hex string:

```
75776374667b786b6364676f6f647061737377645f343637366434373539306234383736657d
```

That being said, I'm not sure what the mistake Charlie made was...

## Flag

`uwctf{xkcdgoodpasswd_4676d47590b4876e}`
