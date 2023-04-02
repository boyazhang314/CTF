# What is the password?

## Description

> N-Team,
>
> Recon managed to capture some web traffic from OmniFlags and think they are onto something. It looks like they are transferring something but we cannot read it.
>
> HQ

## I Spy A Picture

We're given a network packet, which we can open with [Wireshark](https://www.wireshark.org/)

<figure><img src="../../.gitbook/assets/image (1) (4).png" alt=""><figcaption></figcaption></figure>

It looks like a request was made to get `Password.PNG`

In a later packet, we can see this request was succesful:

<figure><img src="../../.gitbook/assets/image (2) (4).png" alt=""><figcaption></figcaption></figure>

Looks like the TCP packets are carrying the PNG data. We can export this data with File > Export Object >  HTTP

<figure><img src="../../.gitbook/assets/image (8) (4).png" alt=""><figcaption></figcaption></figure>

Export the PNG to get an image of the flag

<figure><img src="../../.gitbook/assets/image (5) (1) (3).png" alt=""><figcaption><p>There is the password</p></figcaption></figure>

## Flag

`magpie{wh3r3_15_7h3_p455w0rd}`
