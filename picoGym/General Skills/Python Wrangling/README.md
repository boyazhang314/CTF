# Python Wrangling
**Points: 10**

## Description
Python scripts are invoked kind of like programs in the Terminal... Can you run [this Python script](https://mercury.picoctf.net/static/8e33ede04d02f3765b8c6a6e24d72733/ende.py) using [this password](https://mercury.picoctf.net/static/8e33ede04d02f3765b8c6a6e24d72733/pw.txt) to get [the flag](https://mercury.picoctf.net/static/8e33ede04d02f3765b8c6a6e24d72733/flag.txt.en)?

## Hints
<details>
<summary>Hint 1</summary>

> Get the Python script accessible in your shell by entering the following command in the Terminal prompt: \
> ```$ wget https://mercury.picoctf.net/static/8e33ede04d02f3765b8c6a6e24d72733/ende.py```

</details>

<details>
<summary>Hint 2</summary>

> ```$ man python```

</details>

## Solution
Download all the necessary files and try running the provided Python script with ```python ende.py```

Output:

```Usage: ende.py (-e/-d) [file]```

Most likely ```ende``` stands for "encryption/decryption" and we can conclude the flag ```-e``` is for encrypting and ```-d``` is for decrypting

Run: 

```python ende.py -d flag.txt.en``` 

It will promt for a password with ```Please enter the password:``` \
Pass in the password provided from ```pw.txt``` to get the flag

This can be done in one line with ```python ende.py -d flag.txt.en < pw.txt```

## Flag
```picoCTF{4p0110_1n_7h3_h0us3_aa821c16}```