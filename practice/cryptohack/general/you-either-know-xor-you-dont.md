# You either know, XOR you don't

## Description

> I've encrypted the flag with my secret key, you'll never be able to guess it.\
> \
> _Remember the flag format and how it might help you in this challenge!_\
> \
> `0e0b213f26041e480b26217f27342e175d0e070a3c5b103e2526217f27342e175d0e077e263451150104`

## I don't know

What I do know is that the flag format is `crypto{...}`, which we can try using&#x20;

Recall due to associativity, `flag ^ key = cipher` means `cipher ^ flag = key`

```python
from pwn import xor

HEX = "0e0b213f26041e480b26217f27342e175d0e070a3c5b103e2526217f27342e175d0e077e263451150104"

byteString = bytes.fromhex(HEX)
flagFragment = b"crypto{"

# Try XORing the flag fragment with the string to get a key fragment
print(xor(byteString, flagFragment))
```

We get an interesting byte string output:&#x20;

```
b'myXORke+y_Q\x0bHOMe$~seG8bGURN\x04DFWg)a|\x1dTM!an\x7f'
```

We can assume the key starts with `myXORkey` so let's try XORing our guess with the byte string

```python
# Test out the key
keyBit = b"myXORkey"
print(xor(keyBit, byteString))
```

Lo and behold, this yields the flag

## Flag

`crypto{1f_y0u_Kn0w_En0uGH_y0u_Kn0w_1t_4ll}`
