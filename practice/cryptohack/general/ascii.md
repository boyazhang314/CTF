# ASCII

## Description

> ASCII is a 7-bit encoding standard which allows the representation of text using the integers 0-127.\
> \
> Using the below integer array, convert the numbers to their corresponding ASCII characters to obtain a flag.\
> \
> `[99, 114, 121, 112, 116, 111, 123, 65, 83, 67, 73, 73, 95, 112, 114, 49, 110, 116, 52, 98, 108, 51, 125]`\
> \
> :bulb: In Python, the `chr()` function can be used to convert an ASCII ordinal number to a character (the `ord()` function does the opposite).

## Hit the Right Chord

Using the `chr()` Python function, we can code up a script to take in these values and output the corresponding ASCII to get the flag

```python
NUMBERS = [99, 114, 121, 112, 116, 111, 123, 65, 83, 67, 73, 73, 95, 112, 114, 49, 110, 116, 52, 98, 108, 51, 125]

flag = ""

# Iterate over each number
for num in NUMBERS:
    # Convert to ASCII and append to the string
    flag += chr(num)

print(flag)
```

## Flag

`crypto{ASCII_pr1nt4bl3}`
