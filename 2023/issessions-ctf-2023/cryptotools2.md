# CryptoTools2

## Description

> There are many resources easily available online that are very high quality. CyberChef ([https://gchq.github.io/CyberChef/](https://gchq.github.io/CyberChef/?ref=blog.hnolan.com)) is one of these tools. Become familiar with the website to solve this challenge!\
> \
> Hint: Pay special attention to the instructions for this challenge.

## Use The Tools Again

We are given a `file.tar.gz` file which we can extract to get our cipher

```
NzIlNjUlNzQlNmYlNDMlNTQlNDYlN2IlNzQlNjElNzMlNzQlNzklNmQlNjUlNjElNmMlN2Q=
```

As suggested, let's try cooking with CyberChef:

This looks like a base64 strings, so add the "From Base64" recipe

```
72%65%74%6f%43%54%46%7b%74%61%73%74%79%6d%65%61%6c%7d
```

This looks like hex, so add the "From Hex" recipe

<figure><img src="../../.gitbook/assets/image (18).png" alt=""><figcaption><p>Thanks for the meal</p></figcaption></figure>

## Flag

`retoCTF{tastymeal}`
