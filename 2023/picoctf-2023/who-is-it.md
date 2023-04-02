# who is it

## Description

> Someone just sent you an email claiming to be Google's co-founder Larry Page but you suspect a scam.
>
> Can you help us identify whose mail server the email actually originated from?
>
> Download the email file [here](https://artifacts.picoctf.net/c/499/email-export.eml). Flag: picoCTF{FirstnameLastname}

## Hints

<details>

<summary>1</summary>

whois can be helpful on IP addresses also, not only domain names.

</details>

## Stalking Time

We get an `email-export.eml` file

<figure><img src="../../.gitbook/assets/image (23).png" alt=""><figcaption><p>IP address <span data-gb-custom-inline data-tag="emoji" data-code="1f440">👀</span></p></figcaption></figure>

There is an interesting IP address `173.249.33.206`

Let's look at the identity of this address with `whois 173.249.33.206`

<figure><img src="../../.gitbook/assets/image (39).png" alt=""><figcaption><p>Found you</p></figcaption></figure>

## Flag

`picoCTF{WilhelmZwalina}`
