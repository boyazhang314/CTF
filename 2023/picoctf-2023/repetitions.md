# repetitions

## Description

> Can you make sense of this file?
>
> Download the file [here](https://artifacts.picoctf.net/c/477/enc\_flag).

## Hints

<details>

<summary>1</summary>

Multiple decoding is always good.

</details>

## Recursion

We have a file `enc_flag` which we can view in a text editor

```
VmpGU1EyRXlUWGxTYmxKVVYwZFNWbGxyV21GV1JteDBUbFpPYWxKdFVsaFpWVlUxWVZaS1ZWWnVh
RmRXZWtab1dWWmtSMk5yTlZWWApiVVpUVm10d1VWZFdVa2RpYlZaWFZtNVdVZ3BpU0VKeldWUkNk
MlZXVlhoWGJYQk9VbFJXU0ZkcVRuTldaM0JZVWpGS2VWWkdaSGRXCk1sWnpWV3hhVm1KRk5XOVVW
VkpEVGxaYVdFMVhSbHBWV0VKVVZGWmFWMDVHV2tkYVNHUlZDazFyY0ZkVWJGWlhZVlpLU0dWRlZs
aGkKYlRrelZERldUMkpzUWxWTlJYTkxDZz09Cg==
```

This looks like base64 encoding, so let's try to decode it

{% embed url="https://www.base64decode.org/" %}

<figure><img src="../../.gitbook/assets/image (3).png" alt=""><figcaption><p>Hmm, that's not a flag</p></figcaption></figure>

But it looks like the result is still base64 encoded

From the hint and the title of the question, let's try repeatedly decoding until we get something

* `V1RCa2MyRnRTWGRVYkZaVFltNVNjRmRXYUU5aVJUVnhWVzFhYVdGck5UWmFSVkpQWVRGbmVWVnVRbHBsYTBweVUxWmpNRTVHWjNsVgpXR1JyVFdwV2VsUlZVbE5oTURCNVZXMWFZUXBTTVZWNFZGZHdU MkpWTlVWaVJHeEVXbm93T1VOblBUMEsK`
* `WTBkc2FtSXdUbFZTYm5ScFdWaE9iRTVxVW1aaWFrNTZaRVJPYTFneVVuQlpla0pyU1ZjME5GZ3lVWGRrTWpWelRVUlNhMDB5VW1aYQpSMVV4VFdwT2JVNUViRGxEWnowOUNnPT0K`
* `Y0dsamIwTlVSbnRpWVhObE5qUmZiak56ZEROa1gyUnBZekJrSVc0NFgyUXdkMjVzTURSa00yUmZaR1UxTWpObU5EbDlDZz09Cg==`
* `cGljb0NURntiYXNlNjRfbjNzdDNkX2RpYzBkIW44X2Qwd25sMDRkM2RfZGU1MjNmNDl9Cg==`

<figure><img src="../../.gitbook/assets/image (4).png" alt=""><figcaption><p>Aha!</p></figcaption></figure>

## Flag

`picoCTF{base64_n3st3d_dic0d!n8_d0wnl04d3d_de523f49}`
