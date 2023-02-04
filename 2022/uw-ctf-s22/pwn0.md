# Pwn0

## Description

> ```
> nc uwctf.ml 6001
> ```
>
> ``[`chall0.c`](https://ctf.csclub.uwaterloo.ca/uploads?key=9fcf37dba50432d06cc889773d031adddfbc2575b99745aa970e6e55cc5ba4a1%2Fchall0.c)``

## Get Buffed

As ordered, we run the netcat command to be greeted with a nice message

```
Welcome to Buffer Overflow 0
Can you hack me?
```

{% hint style="info" %}
**Buffer Overflow** - Occurs when the data surpasses the storage of some memory buffer
{% endhint %}

We try overflowing the buffer by injecting a very long string, such as `AAAAAAAAAAAAAAAAAAAAAAAA`&#x20;

* The idea is that there is some buffer which stores the input string, which has a predetermined set size
* Send in as large of an input as possible to try to exceed the buffer's limit

This immediately prints the flag for us

Note that at the time, the source code was not provided, so this was a shot in the dark

## Flag

`uwctf{buff3r0v3rf10w0_3f8361abc77504c5}`
