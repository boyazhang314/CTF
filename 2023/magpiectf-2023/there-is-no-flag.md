# There is no flag

## Description

> Come in, F-Team,
>
> One of our other branches has been working on finding a key flag from here, but weren't able to report anything back. They claim there is no flag to be found. We need you to confirm their findings, or ideally, prove them wrong.
>
> Your starting point: https://github.com/Admin-is-here-GG/React
>
> HQ

## There is a flag

The GitHub repository houses a React app. There are also 3 commits, so let's check them out:

<figure><img src="../../.gitbook/assets/image (27).png" alt=""><figcaption><p>Added what file <span data-gb-custom-inline data-tag="emoji" data-code="1f914">🤔</span></p></figcaption></figure>

The second commit is peculiar, so let's see what was added

<figure><img src="../../.gitbook/assets/image (17).png" alt=""><figcaption></figcaption></figure>

Looks like there's a binary program, so let's download it and try it out. Running it just prints "Hello World!", which isn't very interesting. Let's check out more information with `binwalk program`

<figure><img src="../../.gitbook/assets/image (26).png" alt=""><figcaption><p>Flag <span data-gb-custom-inline data-tag="emoji" data-code="1f440">👀</span></p></figcaption></figure>

There's a zip with a Flag.PNG file? We'll have to get that out! Unzip it with `unzip program` and you'll extract Flag.PNG. However we can't open it because the file contains errors

It seems the PNG is corrupted, so let's uncorrupt it

{% embed url="https://online.officerecovery.com/pixrecovery/" %}

This specific tool will watermark the picture, but you can still read the flag :)

## Flag

`magpie{m15510n_c0mpl373_w17h_r35p3c7}`
