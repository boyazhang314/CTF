# RunedMyDay

## Description

> 'A group of time traveling vikings has been wrecking havoc around the world! Our attempts to communicate with them have only ended in hostility. All what we know is that it sounds like they speak a dialect of English, but its hard to understand them. While they also have a written language, we can't seem to understand that either. Fortunately, we do have a letter that we believe was written by them on December 5th, 2022. If you can decipher their written language, perhaps we'd be able to write to them and calm them down?
>
> Format: retroCTF{The Magic Words Here}'

## Fear not Death

We are given a picture of what appears to be a letter, consisting of odd symbols. Upon further inspection, it seems like each symbol corresponds to a letter in the English alphabet

Now the fun part:

<figure><img src="../../.gitbook/assets/image (15).png" alt=""><figcaption></figcaption></figure>

Using the fact that the letter was written on December 5th, we can decrypt some letters which can further be used to figure out words and guess more letters

Eventually, we get an interesting paragraph:

```
However Charons boat will not respond to our pleas
We chanted the magic words known as the flag
omajesticlegendaryboathearourpleasandbringushome
but the boat will not respond
```

## Flag

`retroCTF{Oh Majestic Legendary Boat Hear Our Pleas And Bring Us Home}`
