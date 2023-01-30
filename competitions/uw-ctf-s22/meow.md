# Meow

## Description

> Please talk to `ctf.csclub.uwaterloo.ca` on the alternative port for HTTP, it will meow you your flag.

## Here Kitty Kitty

There are a ton of cat references in this description, and with the need to talk on a port, this all points to using [netcat](https://en.wikipedia.org/wiki/Netcat)

{% hint style="info" %}
**Netcat** - A tool to read and write data across network connections
{% endhint %}

To talk to the website run `ncat ctf.csclub.uwaterloo.ca`

* Or `nc`

This doesn't produce anything, most likely because we need to specify the port

Googling "alternative port for HTTP" gives the following results

<figure><img src="../../.gitbook/assets/image (6).png" alt=""><figcaption></figcaption></figure>

Ultimately, port 8080 is the one that gets the cat out of the bag

```
ncat ctf.csclub.uwaterloo.ca 8080
```

## Flag

`uwctf{netcat_bcf2e43f3c9e97e9}`
