# Education Comes First

## Description

> **URGENT** W-Team,
>
> An OmniFlags educational website has been hacked by one of the captives in the space prison, who we believe has left a message for us. Can you find it?
>
> Good luck,\
> HQ

## Let's Study

<figure><img src="../../.gitbook/assets/image (20).png" alt=""><figcaption><p>Oooh bubbles</p></figcaption></figure>

We are greeted with an education page with floating bubbles. The website is pretty bare, with `index.html` , `whatWeDo.html`, and `about.html` being the only pages we can see or go to. Let's inspect the page source

We find a file called `flash.js`. This code is not very readable, however searching through there is a fragment that seems interesting

```javascript
hex2a('6d61677069657b57335f525f5337314c4c5f483352337d')
```

We have a hex string, which we can convert into readable text to get the flag

## Flag

`magpie{W3_R_S71LL_H3R3}`
