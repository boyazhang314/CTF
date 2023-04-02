# Decade Capsule

> A group of campers discovered a software Time Capsule that is set to open in 2030! Unfortunately, they're very impatient and are curious to know what is inside of it. Are you able to bust open the "Decade Capsule" a bit earlier than expected?

## Greetings From the Future!

We have an executable `capsule.exe`

Note that this is a Windows executable

<figure><img src="../../.gitbook/assets/image (1) (5).png" alt=""><figcaption></figcaption></figure>

So we can't execute as `./capsule.exe` as normal. Instead, we can run `wine capsule.exe`

<figure><img src="../../.gitbook/assets/image (6).png" alt=""><figcaption></figcaption></figure>

And the output is something like `213087208 more seconds until the capsule opens!`, which after some calculations, appears to be too long for us to wait

We can't open the capsule until 2030!

Unfortunately, time travel is not possible, however we can trick the program by setting our machine's time using `date -s '2030-01-01 00:00:00'`

The capsule can now be opened:

```
Hello, people of the future!
Thank you for waiting so long - here is the long awaited information:
retroCTF{u51ng_th3_0ld35t_tr1ck_1n_th3_b00k!!1}
```

## Flag

retroCTF{u51ng\_th3\_0ld35t\_tr1ck\_1n\_th3\_b00k!!1}
