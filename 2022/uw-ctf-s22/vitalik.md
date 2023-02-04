# Vitalik

## Description

> Vitalik's creation uses a hashing function. Find our _path_ to the mapping for nothing.

## And There's Nothing

Let's start by figuring out who Vitalik is and what is their "creation"

Google searching reveals a man named Vitalik Buterin, founder of Ethereum, which is a blockchain

<figure><img src="../../.gitbook/assets/image (4) (2).png" alt=""><figcaption><p>Man is busy</p></figcaption></figure>

As per the description, this Ethereum uses a hash function, so let's try to find what it is

<figure><img src="../../.gitbook/assets/image (15).png" alt=""><figcaption></figcaption></figure>

As we can see, it uses **Keccak-256**

{% embed url="https://emn178.github.io/online-tools/keccak_256.html" %}

This hash function is very impressive, as it's able to hide the length of the message. For the next bit, I stewed about what `Find our`` `_`path`_` ``to the mapping for nothing` meant

I tried encoding "path" as well as the whole message, then shoving it into the `uwctf{...}` format, though that was certainly not correct

I was stumped. This was the final challenge I had to solve and I was at a lost

It wasn't until I played around with Keccak-256 that I realized that even the empty string hashes to something, and the word "path" had a very different meaning...

The empty string, or "nothing", hashes to `c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470`

We can add this to the "path" of the CTF website,&#x20;

```
https://ctf.csclub.uwaterloo.ca/c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470
```

Open it with any text editor to get the flag

## Flag

`uwctf{6c48c02b6682151e}`
