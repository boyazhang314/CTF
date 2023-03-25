# Crack Me

## Description

> Can you figure out the correct cipher being used?

## Cracked

We are given a Python script `crack.py`

The first bit gives us a cipher, which is the encrypted flag, and an alphabet

```python
flag = "C6EC@r%uLJ_F#b0c04Co4<0bIAbCEN"

alphabet = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ"+ \
            "[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
```

We also get a list of ciphers

```python
ciphers = [
           "RSA", "ELGAMAL", "DIFFIE-HELLMAN", "DES","MD5", "SHA1", "SHA256", "SHA512", "SHA3", "RSA",
           "CETACEAN", "CIPHERSABER2", "RAIL FENCE", "A1Z26", "BLOWFISH", "RC4", "RC5", "RC6",
           "RC7", "RC8", "RC9", "RC10", "RC11", "RC12", "RC13", "RC14", "RC15", "RC16", "RC17",
           "RC18", "RC19", "RC20", "RC21", "RC22", "RC23", "RC24", "RC25", "RC26", "RC27",
           "RC28", "RC29", "RC30", "RC31", "RC32", "RC33", "RC34", "RC35", "RC36", "RC37",
           "RC38", "RC39", "RC40", "RC41", "RC42", "RC43", "RC44", "RC45", "RC46", "RC47",
           "RC48", "RC49", "ROT47", "RC50", "RC51", "RC52", "RC53", "RC54", "RC55", "RC56", "RC57",
           "RC58", "Morse", "RC59", "RC60", "RC61", "RC62", "RC63", "RC64", "RC65", "RC66", "RC67",
           "RC68", "RC69", "RC70", "RC71", "RC72", "RC73", "RC74","Vigenere", "RC75", "RC76", "RC77",
           "RC78", "RC79", "RC80", "RC81", "RC82", "RC83", "RC84", "RC85", "RC86", "RC87",
           "RC88", "RC89", "RC90", "RC91", "RC92", "RC93", "RC94", "RC95", "RC96", "RC97",
           "RC98", "RC99", "RC100", "RC101","ATBASH", "RC102", "RC103", "RC104", "RC105", "RC106",
           "RC107","BACONIAN", "RC108", "RC109", "RC110", "RC111", "RC112", "RC113", "RC114", "RC115",
           "RC116", "RC117", "AFFINE","RC118", "RC119", "RC120", "RC121", "RC122", "RC123", "RC124",
           "RC125", "RC126", "RC127", "RC128", "SIMPLE SUBSTITUTION", "RC129", "RC130", "RC131", "RC132", "RC133", "AES",
           "RC134", "RC135", "RC136", "RC137", "RC138", "RC139", "RC140", "RC141", "RC142",
           "RC143", "RC144", "ROT13", "RC145", "RC146", "RC147", "BIFID","RC148", "RC149", "RC150", "RC151",
           "RC152", "RC153", "POLYBIUS SQUARE","RC154", "RC155", "RC156", "RC157", "RC158", "RC159", "RC160",
           "RC161", "RC162", "RC163", "RC164","Caesar", "RC165", "RC166", "RC167", "RC168", "RC169",
           "RC170", "RC171", "RC172", "CODES AND NOMENCLATORS","RC173", "RC174", "RC175", "RC176", "RC177", "RC178"
        ]
```

Finally, the `main` function just calls `get_flag` so let's examine it

```python
def get_flag():
    while True:
        input1 = input("Enter the correct cipher: ").upper()
        if input1 == ciphers[61]:
            key = int(input("\nEnter the key: "))
            decode_flag(flag,key)
            break
        else:
            print('Wrong shift cipher!\n')
        
   
def decode_flag(flag, key):
    decoded = ""

    for x in flag:
        index = alphabet.find(x)
        temp_index = (index + key) % len(alphabet)
        decoded = decoded + alphabet[temp_index]
    print(decoded)
```

Looks like `get_flag` takes in a cipher name then a key, which is an integer, and checks if the cipher is `ciphers[61]` from the provided list of ciphers

Since I'm not in the mood to count, let's add `print(ciphers[61])` to the main function and run the script with `python crack.py`

The cipher is ROT47 so we can enter that in, then it prompts us with `Enter the key:`

Taking a look at the code, the key is used with `decode_flag` to decrypt the cipher using the alphabet, and since the cipher is ROT47 the key must be 47

* This is due to `decode_flag` decrypting by "rotating" the flag with the key at the line `(index + key) % len(alphabet)`

<figure><img src="../../.gitbook/assets/image (16).png" alt=""><figcaption><p>Cracked</p></figcaption></figure>

## Flag

`retroCTF{y0uR3_4_cr@ck_3xp3rt}`
