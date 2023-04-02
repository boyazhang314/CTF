# money-ware

## Description

> Flag format: picoCTF{Malwarename}
>
> The first letter of the malware name should be capitalized and the rest lowercase.
>
> Your friend just got hacked and has been asked to pay some bitcoins to `1Mz7153HMuxXTuR2R1t78mGSdzaAtNbBWX`. He doesn’t seem to understand what is going on and asks you for advice. Can you identify what malware he’s being a victim of?

## Hints

<details>

<summary>1</summary>

Some crypto-currencies abuse databases exist; check them out!

</details>

<details>

<summary>2</summary>

Maybe Google might help.

</details>

## Malware Mining

As per the hints, we can try looking for some crypto-currency abuse databases

{% embed url="https://www.bitcoinabuse.com/" %}

Try to search for the hash `1Mz7153HMuxXTuR2R1t78mGSdzaAtNbBWX`

{% embed url="https://www.bitcoinabuse.com/reports/1Mz7153HMuxXTuR2R1t78mGSdzaAtNbBWX" %}

<figure><img src="../../.gitbook/assets/image (1).png" alt=""><figcaption></figcaption></figure>

We find an interesting blog about "Petya" [https://blog.avira.com/petya-strikes-back/](https://blog.avira.com/petya-strikes-back/) which gives us a malware name

## Flag

`picoCTF{Petya}`
