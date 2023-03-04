# Audio

## Description

> _Ooooooh spoooooopy!_
>
> ``[`whoops.mp3`](https://ctf.csclub.uwaterloo.ca/uploads?key=eb30eb1b3642806cf078d8d213e10c8d5a178473465c825f206b36898a9b2b63%2Fwhoops.mp3)``

## \*Intelligible Nonsense

The MP3 file consists of extremely fast, high-pitched talking in a language that, as far as my ears can tell, does not exist on this planet

Yet at the same time, there is something distinctly _English_ about the vernacular

Let's open the file in [Audacity](https://www.audacityteam.org/) and see what we can do

<figure><img src="../../.gitbook/assets/image (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

The high pitch and speed seem to imply it was sped up, so let's try slowing it down

* `CTRL-A` to select all
* Effect > Pitch and Tempo > Change Speed
* Set the Speed Multiplier to 0.5 the original speed and Apply

It's a lot slower and deeper, but still not something we can understand

The sounds being produced sounds like the air is being sucked from the speaker's mouth. Of course, this is the opposite of what happens when humans speak, so this seems to imply reversing

* `CTRL-A` to select all
* Effect > Pitch and Tempo > Reverse

Now we can actually pick up some words, such as "Charlie" and "Fox". These words seem peculiar, and after searching around they are part of the phonetic alphabet

{% embed url="https://en.wikipedia.org/wiki/NATO_phonetic_alphabet" %}

We now listen _veryyyyyy closelyyyyy_ as we try to decipher what is being said. It takes a while.

## Flag

`uwctf{4ud10f1l3_dacb9b194dd9cec7}`
