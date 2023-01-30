# 0s and 1s

## Description

> {% code overflow="wrap" %}
> ```
> 01110101 01110111 01100011 01110100 01100110 01111011 01100010 00110001 01101110 00110100 01110010 01111001 01011111 00111000 01100001 00110010 00110110 00110110 01100001 00111001 00110101 01100011 00110001 00110011 00110101 00110001 01100100 00110000 00110010 01111101
> ```
> {% endcode %}

## Binary Time :v:

As stated in the title **0s and 1s**, the provided text solely consists of 0s and 1s, which implies a binary string

{% hint style="info" %}
**Binary string** - Sequence of bytes (8 bits) where each byte represents a text character&#x20;
{% endhint %}

And the string is conveniently already split for us!

## Byte the Text

We can easily convert this string using online websites

{% embed url="https://www.rapidtables.com/convert/number/binary-to-ascii.html" %}

However, we could also code up a simple function in Python

````python
```python
def binToText(binary):
    text = ""

    # Split the binary into bytes
    # Note this line may differ depending on the input
    binBytes = binary.split(' ')

    # Iterate over each byte
    for byte in binBytes:
        # Convert the base 2 byte into a base 10 number
        num = int(byte, 2)

        # Convert the number into its character counterpart
        char = chr(num)

        # Add to the string
        text += char
    
    return text
```
````

<details>

<summary>References for Python functions</summary>

* `int()` - [https://www.programiz.com/python-programming/methods/built-in/int](https://www.programiz.com/python-programming/methods/built-in/int)
* `chr()` - [https://www.programiz.com/python-programming/methods/built-in/chr](https://www.programiz.com/python-programming/methods/built-in/chr)

</details>

## Flag

`uwctf{b1n4ry_8a266a95c1351d02}`
