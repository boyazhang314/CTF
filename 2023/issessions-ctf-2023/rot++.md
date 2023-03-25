# ROT++

## Description

> Standard Caesar ciphers are too easy! This cipher was encrypted by starting at 0 and incrementing the shift for each subsequent character in the string. Write a simple program to decrypt this cipher!
>
> Hint 1: The ‘alphabet’ for this challenge is A-Z followed by 0–9.
>
> Hint 2: retroCTF{?????\_??\_???\_????}

## ROT--

We are given a text file with the cipher `RPV47_ny_0wx_ol4b`

As described, this was encrypted with a modified Caesar cipher where the shift starts at 0 and increases by 1 for each subsequent character

* Start with a shift of 0 and increase for each character, then undo the shift
* We are also given the alphabet used, which doesn't include underscores so we can safely ignore those

```python
# Alphabet consists of all uppercase so conver cipher
cipher = 'RPV47_ny_0wx_ol4b'.upper()

# Alphabet
alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789' # Length of 36

flag = ''
shift = 0

for ch in cipher:
    # Alphabet has no underscores
    if ch == '_':
	flag += '_'
    else:
        # Reverse the shift
        flag += alphabet[alphabet.find(ch) - shift % 36]
        shift += 1
		
print(flag)

```

## Flag

`retroCTF{ROT13_IS_TOO_EASY}`
