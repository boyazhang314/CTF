# XOR Starter

## Description

> XOR is a bitwise operator which returns 0 if the bits are the same, and 1 otherwise. In textbooks the XOR operator is denoted by ⊕, but in most challenges and programming languages you will see the caret `^` used instead.\
> \
> For longer binary numbers we XOR bit by bit: `0110 ^ 1010 = 1100`. We can XOR integers by first converting the integer from decimal to binary. We can XOR strings by first converting each character to the integer representing the Unicode character.\
> \
> Given the string `label`, XOR each character with the integer `13`. Convert these integers back to a string and submit the flag as `crypto{new_string}`.\
> \
> :bulb: The Python `pwntools` library has a convenient `xor()` function that can XOR together data of different types and lengths. But first, you may want to implement your own function to solve this.

| A | B | Output |
| - | - | ------ |
| 0 | 0 | 0      |
| 0 | 1 | 1      |
| 1 | 0 | 1      |
| 1 | 1 | 0      |

## EXOOR

Since we're lazy, we can use the `pwntools` library to help us

```python
from pwn import xor

# xor function uses and returns byte strings
label = b"label"
newString = b""

# Iterate over each character
for ch in label:
    # XOR with 13 then append to the string
    newString += xor(ch, 13)

print(newString)
```

If you're feeling more daring, you can code up your own `xor` function

```python
def xor(a, b):
    # Note that ASCII characters are represented by bytes
    # For consistency format to binary of 8 bits
    xorString = ""

    # Convert to binary
    # Padded to 8 bits
    binA = '{:08b}'.format(a)
    binB = '{:08b}'.format(b)
    
    # Iterate over each bit and XOR them
    for i in range(len(binA)):
        if binA[i] == binB[i]:
            xorString += "0"
        else:
            xorString += "1"
    
    # Convert binary string to text
    ch = chr(int(xorString, 2))
    b = ch.encode('utf-8') # Return byte
    return b
```

## Flag

`crypto{aloha}`
