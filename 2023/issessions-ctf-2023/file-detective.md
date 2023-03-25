# File Detective

## Description

> I wonder what type of file the flag is....

## Deduction Time

We are given a text file `the-flag.txt`, or so it seems...

Let's try printing it out

<figure><img src="../../.gitbook/assets/image (2).png" alt=""><figcaption><p>Hmmmm</p></figcaption></figure>

Our keen eye of observation will notice the "PNG" printed at the top, so perhaps this is not a `txt` file after all :detective:

Indoubetably, we can check with `file the-flag.txt`

```
the-flag.txt: PNG image data, 1536 x 864, 8-bit/color RGB, non-interlaced
```

Rename the file with the proper extension using `mv the-flag.txt the-flag.png`

<figure><img src="../../.gitbook/assets/image.png" alt=""><figcaption><p>Another mystery solved</p></figcaption></figure>

## Flag

`retroCTF{TH4t_w45_P34nU75}`
