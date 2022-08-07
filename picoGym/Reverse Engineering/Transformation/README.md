# Transformation
**Points: 20**

## Description
I wonder what this really is... [enc](https://mercury.picoctf.net/static/2b4cea9b07db22bf4f933fddd1b8caa9/enc)

```''.join([chr((ord(flag[i]) << 8) + ord(flag[i + 1])) for i in range(0, len(flag), 2)])```

## Hints
<details>
<summary>Hint 1</summary>

> You may find some decoders online

</details>

## Solution
The provided code snippet is Python code which manipulates a string and encodes it in some fashion

```(ord(flag[i]) << 8)``` - This takes a character ```flag[i]```, calls the function ```ord()```, which returns the integer representation of the character, and bit shifts the integer left by 8 via ```<< 8``` (multiplied by 2^8 = 246)

```ord(flag[i + 1])``` - This simply returns the integer representation of character ```flag[i+1]```

These two integer values are summed together with ```+```. The final result is then converted into a character by calling the ```chr()``` function

**Note that ord() and chr() are inverse functions**

```for i in range(0, len(flag), 2)``` - This implies that the above process is repeated for every other character in ```flag```, or all even indices, as ```i``` iterates until the length of ```flag```, starting at 0 and incrementing by 2. Ultimately, this creates an array of characters

Finally, the array of characters are joined by ```''.join()``` into a singular string. Since this is the encoding process, the resulting string is the provided ```灩捯䍔䙻ㄶ形楴獟楮獴㌴摟潦弸弰㑣〷㘰摽```

*We must reverse the encoding*

First, we can split the string into an array of characters with ```list(enc)```.
We can then inverse the ```chr()``` call with ```ord()```. From the for loop above, we know each character in the array actually represents two characters in the flag

Taking a look at ```(ord(flag[i]) << 8) + ord(flag[i + 1])```,  the first integer is multiplied by 2^8 = 256, then the second integer is added to it. To reverse this process, we can call mod 256 to get the second integer. The first integer can be found by performing bitshifts to the right by 8.

These integers should be converted into characters with ```chr()```. We repeat the process for all characters in the encoding

## Flag
```picoCTF{16_bits_inst34d_of_8_04c0760d}```