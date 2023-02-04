# WASM

## Description

> [https://ctf.csclub.uwaterloo.ca/ctf-uploads/noob/wasm/chall.html](https://ctf.csclub.uwaterloo.ca/ctf-uploads/noob/wasm/chall.html)

## All the Colors of the Rainbow

The website shows nothing more than a simple message

<figure><img src="../../.gitbook/assets/image (10) (1) (1).png" alt=""><figcaption></figcaption></figure>

Here `check_flag()` seems to refer to some JavaScript function

We can try running it in the Console tab with various arguments

<figure><img src="../../.gitbook/assets/image (5) (2).png" alt=""><figcaption></figcaption></figure>

Here we can learn a few things

* `check_flag` does not like any of the answers I gave and responds with a `not answer`
* `check_flag` takes in a string

After trying some more strings, I stopped to think about the hint a little more than the surface-level colours that were listed

Thinking about _technologies blue, red, yellow, blue, green, red_, there was a certain name that came to mind...

Windows.

But that didn't work. Eventually, I came to the actual solution

<figure><img src="../../.gitbook/assets/image (8).png" alt=""><figcaption><p>check_flag is quite picky...</p></figcaption></figure>

## Flag

`uwctf{W45M_4772a105377592cd}`
