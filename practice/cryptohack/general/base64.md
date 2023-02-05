# Base64

## Description

> Another common encoding scheme is Base64, which allows us to represent binary data as an ASCII string using an alphabet of 64 characters. One character of a Base64 string encodes 6 binary digits (bits), and so 4 characters of Base64 encode three 8-bit bytes.\
> \
> Base64 is most commonly used online, so binary data such as images can be easily included into HTML or CSS files.\
> \
> Take the below hex string, _decode_ it into bytes and then _encode_ it into Base64.\
> \
> `72bca9b68fc16ac7beeb8f849dca1d8a783e8acf9679bf9269f7bf`\
> \
> :bulb: In Python, after importing the base64 module with `import base64`, you can use the `base64.b64encode()` function. Remember to decode the hex first as the challenge description states.

## 2^6

As outlined in the description, we can decode the hex string into bytes and then encode it with Base64

```python
import base64

HEX = "72bca9b68fc16ac7beeb8f849dca1d8a783e8acf9679bf9269f7bf"

# Decode the hex into a byte string
byteString = bytes.fromhex(HEX)

# Encode the byte string with base64
b64Flag = base64.b64encode(byteString)

print(b64Flag) # This flag has a slightly different format
```

{% hint style="warning" %}
Important thing I learned: Do NOT call your Python file "base64.py" because it overrides the standard library module `base64` which causes a bunch of errors :(
{% endhint %}

## Flag

`crypto/Base+64+Encoding+is+Web+Safe/`
