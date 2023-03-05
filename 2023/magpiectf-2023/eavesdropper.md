# Eavesdropper

## Description

> Hello N-Team,
>
> According to recon, one of the captives has been sending flags discretely to our endpoints, and the key flag is going to have something to do with her favourite food. Unfortunately, you just can't remember what it was...
>
> She's got a sweet tooth, so research has determined that it has to do with some kind of dessert, like **cookies** or **cakes**. What was her favourite flavour? **vanilla**? **oreo**? **chocolate chip**? Or maybe **pistachio**?
>
> Good luck,\
> HQ

## Where's the Cookies?

We're given a network packet, which we can open with [Wireshark](https://www.wireshark.org/)

<figure><img src="../../.gitbook/assets/image (9).png" alt=""><figcaption></figcaption></figure>

There are 120000 packets! Way too many to search through, so let's try examining the packets

* Right Click > Follow > TCP Stream

<figure><img src="../../.gitbook/assets/image (24).png" alt=""><figcaption></figcaption></figure>

We can see what looks like a flag in a `magpie{...}` format, but what's inside the curly braces isn't readable text. In fact, if we iterate through the other streams, we see that they all have what appears to be a flag, with random characters in it

As mentioned by the description, the captive has been sending many flags, however there is only one "key flag", which is what we will likely need to find

Something I noticed is that there are HTTP POST requests which sends a request with a flag, then there are a bunch of 405 Not Allowed responses, which implies these requests do not contain the key flag. Thinking that the key flag would be the one that allows for a successful response, I tried CTRL-F "200" and "Success", but no packets were found

Of course manually checking each packet would be impossible, so I needed a way to extract all the possible flags then search for the supposed key flag. After some Googling, I learned that you can run the `strings` comment on a .pcapng file to get its contents, then we can run `grep` to extract all the flags

```
strings nothing_to_see_here.pcapng | grep -i magpie > out
```

We're told that the contents of the flag have something to do with cake, cookies, chocolate, or some sort of flavored desert, so let's try CTRL-F with educated guesses:

<figure><img src="../../.gitbook/assets/image (23).png" alt=""><figcaption><p>Found chocolate chip cookies <span data-gb-custom-inline data-tag="emoji" data-code="1f36a">🍪</span></p></figcaption></figure>

## Flag

`magpie{chOc0LatE_Ch1p_c0Ok1e5}`
