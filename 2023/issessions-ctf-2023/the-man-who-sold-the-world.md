# The Man Who Sold the World

## Description

> Want some more points? No problem! The points are locked behind this super secret chest. Where is the key you ask? That's for you to solve. Good luck!

## Tastes Purple

We are provided an audio file `message.wav` and an image `table.png`

<figure><img src="../../.gitbook/assets/image (30).png" alt=""><figcaption><p>Purple <span data-gb-custom-inline data-tag="emoji" data-code="1f7e3">🟣</span></p></figcaption></figure>

The table doesn't provide much information, so let's listen to the message

Unfortunately, it is not music. Rather it is a series of long beeps and short beeps, implying Morse code. Instead of deciphering it by ear, I opened it with Audacity to do it visually

<figure><img src="../../.gitbook/assets/image (31) (1).png" alt=""><figcaption></figcaption></figure>

{% embed url="https://en.wikipedia.org/wiki/Morse_code" %}

We get the text `UII3CD1GTRG3GE`

Of course, this is nonsense, but let's think about the purple table again...

With the help of Googling "alphabet by alphabet table" I found that it lead me to something familiar: the Vigenère cipher

{% embed url="https://pages.mtu.edu/~shene/NSF-4/Tutorial/VIG/Vig-Base.html" %}

A cipher which takes in a key and "adds" the letters of the key to each letter of the plaintext in order to formulate the ciphertext. Equivalently, you can use the alphabet by alphabet table with the plaintext and key letters as coordinates to produce each letter in the ciphertext

Now what is the key?

Well, the table is quite _purple_

<figure><img src="../../.gitbook/assets/image (33).png" alt=""><figcaption></figcaption></figure>

## Flag

`retroCTF{FOR3NS1CEXP3RT}`
