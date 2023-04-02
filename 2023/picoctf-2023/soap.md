# SOAP

## Description

> The web project was rushed and no security assessment was done. Can you read the /etc/passwd file?

## Hints

<details>

<summary>1</summary>

XML external entity Injection

</details>

## Soapy Bubbles

We have a simple website with "Detail" buttons we can click

<figure><img src="../../.gitbook/assets/image (31).png" alt=""><figcaption></figcaption></figure>

Clicking one of these buttons reveals additional information, for example

```
Special Info:::: University in Kigali, Rwanda offereing MSECE, MSIT and MS EAI
```

What's interesting is, when observing through the Network tab, clicking on the buttons creates a POST request to `/data`

As given by the hint, we are supposed to perform "XML external entity injection"

Googling this, we find that there is a payload which will let us view the `/etc/passwd` file

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE foo [ <!ENTITY xxe SYSTEM "file:///etc/passwd"> ]>
<data><ID>&xxe;</ID></data>
```

Intercept the POST request using Burp Suite or something equivalent

&#x20;

<figure><img src="../../.gitbook/assets/image (43).png" alt=""><figcaption></figcaption></figure>

Replace the payload then forward the request and the website will print something interesting

<figure><img src="../../.gitbook/assets/image (42).png" alt=""><figcaption></figcaption></figure>

## Flag

`picoCTF{XML_3xtern@l_3nt1t1ty_4dbeb2ed}`
