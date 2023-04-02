# findme

## Description

> Help us test the form by submiting the username as `test` and password as `test!`

## Hints

<details>

<summary>1</summary>

any redirections?

</details>

## Searching Hard

The website has two fields for username and password

<figure><img src="../../.gitbook/assets/image (40).png" alt=""><figcaption></figcaption></figure>

As directed in the question, we try username `test` and password `test!`

Note that trying anything else leads you to a prompt `try username:test and password:test!`

<figure><img src="../../.gitbook/assets/image (36).png" alt=""><figcaption></figcaption></figure>

We get sent to the `/home` page, but wait! Was it just me, or did the URL flash a bit when we came here? That means there was some sort of redirection

Using the Network tab we find that there were two redirections

* `/next-page/id=cGljb0NURntwcm94aWVzX2Fs`
* `/next-page/id=bF90aGVfd2F5XzI1YmJhZTlhfQ==`

These appear to be base64 encoded strings, however both look a bit short. So let's concatenate the two and decrypt `cGljb0NURntwcm94aWVzX2FsbF90aGVfd2F5XzI1YmJhZTlhfQ==` to get the flag

## Flag

`picoCTF{proxies_all_the_way_25bbae9a}`
