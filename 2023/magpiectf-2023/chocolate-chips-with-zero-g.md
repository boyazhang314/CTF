# Chocolate Chips with Zero-G

## Description

> Hello W-Team,
>
> Recon has found an old website from OmniFlags. We think there may be an insecure admin portal somewhere. See if there is anything you can find! They never did hire the best developers.
>
> Report back,\
> HQ
>
> \
> [http://srv1.2023.magpiectf.ca:1234](http://srv1.2023.magpiectf.ca:1234)

## Milky Way

<figure><img src="../../.gitbook/assets/image (10) (1).png" alt=""><figcaption></figcaption></figure>

We are brought to the moon on this website, with both a navigation bar and a sidebar, though some of the options, such as "CAREERS" and "FLAGS", link to nothing

From the description we know we're looking for an "insecure admin portal", so let's see what we can find. Inspecting the page shows that we are currently on `index.html`. Let's try searching for the admin page with `admin.html`

\-> This is done by adding the `/admin.html` path to the URL

<figure><img src="../../.gitbook/assets/image (19).png" alt=""><figcaption><p>We found the login portal</p></figcaption></figure>

Let's try logging in with any test credential, say "admin" and "admin"

<figure><img src="../../.gitbook/assets/image (3).png" alt=""><figcaption><p>No luck</p></figcaption></figure>

Inspect element again, and in the Sources tab we can find a `script.js`. Looking through it, it seems to hold the logic for login. Notably, we find this bit for when the credentials are submitted:

<figure><img src="../../.gitbook/assets/image (9) (1) (2).png" alt=""><figcaption><p>Admin</p></figcaption></figure>

Looks like it sets the cookie `admin` to `false` when we login. If we have submitted our credentials and the `admin` cookie is false, it redirects us to `denied.html`

We can go into Application > Storage > Cookies to find the `admin` cookie, set it to `true`, then refresh the page to get the flag

## Flag

`magpie{bu7-7h3-m1Lk-ju57-fl04T5-4W4y!}`
