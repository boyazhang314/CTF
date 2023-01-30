# steg

## Description

> Yes, it's been there on the homepage all this time. [https://ct](https://ctf.uwaterloo.ca/)
>
>
>
> [f.uwaterloo.ca/](https://ctf.uwaterloo.ca/)

## Stegosaurus Rex

From the title "steg" I automatically think of steganography, where information is hidden within images, though other mediums such as music are also possible

We go to the website and play around, inspecting elements for potential hidden flags. Eventually, something catches my eye:

{% embed url="https://git.uwaterloo.ca/ctf" %}
Their GitLab is linked :eyes:
{% endembed %}

This houses all the code for the UW CTF website. We can see a very telling commit name

<figure><img src="../../.gitbook/assets/image (3).png" alt=""><figcaption></figcaption></figure>

"Chall"? "Background"? Oho!

Checking out the commit, we can see an image file change

<figure><img src="../../.gitbook/assets/image (1).png" alt=""><figcaption></figcaption></figure>

There must be something very peculiar about the new image :thinking:

I would use Photoshop if I had it. Unfortunately, the only image software I know how to use is MS Paint. Instead, let's try some online websites

{% embed url="https://www.aperisolve.com/" %}

One of the images produced shows the flag, albeit very hard to read

<figure><img src="../../.gitbook/assets/image (6).png" alt=""><figcaption></figcaption></figure>

## Flag

`uwctf{stegfrthistime_9b0a89e16a81af67}`
