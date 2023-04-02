# hideme

## Description

> Every file gets a flag.
>
> The SOC analyst saw one image been sent back and forth between two people. They decided to investigate and found out that there was more than what meets the eye [here](https://artifacts.picoctf.net/c/260/flag.png).

## Finders Keepers

We're given a `flag.png` file

<figure><img src="../../.gitbook/assets/Untitled (1).png" alt=""><figcaption></figcaption></figure>

Doesn't look that interesting, let's seewhat else this "png" is hiding

Running `binwalk flag.png` reveals that the file is actually a ZIP

We can run `unzip flag.png` to get a folder with a PNG in it

<figure><img src="../../.gitbook/assets/Untitled2 (1).png" alt=""><figcaption><p>Found it!</p></figcaption></figure>

## Flag

`picoCTF{Hiddinng_An_imag3_within_@n_ima9e_81cc7947}`
