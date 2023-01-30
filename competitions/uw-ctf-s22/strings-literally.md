# Strings, literally

## Description

> Uh oh! GIMP corrupted my file while I was making a steg chall, can you still find the flag?
>
> ``[`steg.xcf`](https://ctf.csclub.uwaterloo.ca/uploads?key=12cf14b584bb0ab4fe90d0bbbc35527a682daf1add31f5a14cea4768bcc15213%2Fsteg.xcf)``

## Stringy Strings

As implied by the title, we can run the `strings` Linux command on the file to find the flag

```
strings steg.xcf
```

The output is quite long, so I directed the output into an `out` file with `strings steg.xcf > out`

<figure><img src="../../.gitbook/assets/image (2).png" alt=""><figcaption><p>Much easier to search</p></figcaption></figure>

## Flag

`uwctf{h1dd3n5t3g_54c10511eda3e8f9}`
