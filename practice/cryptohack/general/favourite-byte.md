# Favourite byte

## Description

> For the next few challenges, you'll use what you've just learned to solve some more XOR puzzles.\
> \
> I've hidden some data using XOR with a single byte, but that byte is a secret. Don't forget to decode from hex first.\
> \
> `73626960647f6b206821204f21254f7d694f7624662065622127234f726927756d`

## Some of my Favourite Things

We are told the data was XORed with a single byte, which is 8 bits. This means we can brute-force all possible bytes, from `00000000` to `11111111` and XOR it with the hex until we get what looks like a flag!

```python
from pwn import xor

HEX = "73626960647f6b206821204f21254f7d694f7624662065622127234f726927756d"

byteString = bytes.fromhex(HEX)
print(byteString)

# 2 ^ 8 = 256
# A byte can represent numbers up to 256
for i in range(256):
    # Get a byte and XOR with the byte string
    xorString = xor(byteString, i)
    print(xorString)
```

<figure><img src="../../../.gitbook/assets/image (3) (1).png" alt=""><figcaption><p>We can find the flag</p></figcaption></figure>

## Flag

`crypto{0x10_15_my_f4v0ur173_by7e}`
