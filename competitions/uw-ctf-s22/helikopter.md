# Helikopter

## Description

> Helikopter Helikopter! Or is it a mirror? What does `g?` do, In vim when given these characters? `hjpgs{ebg13_n6q5ro2}`

## GG?

I do not know what a Helikopter is, but I do vaguely know about vim, so let's start with that

I don't use vim myself so let's research about `g?`

{% embed url="https://renenyffenegger.ch/notes/development/vim/commands/g" %}

There are a few things of note:

* `g` scans lines, marks those that match a pattern, and runs a command on those lines
  * It may also take in a `range`
* There is nothing about `?`

This doesn't seem too fruitful, so let's try getting our hands dirty

## :q&#x20;

1. Open vim with `vim`
2.  Struggle and panic as you try to figure out how to actually type in the new file&#x20;

    \-> (I believe pressing the `i` key is what allows you to "insert", though I do not know if that's the correct command)
3. Insert the text `hjpgs{ebg13_n6q5ro2}` into the file
4. To call commands, press the `:` key and type in `g?` then `ENTER`

And... nothing happens?

Well except for the file seemingly deleting the text I painstakingly typed in. Also this error:

<figure><img src="../../.gitbook/assets/image (5).png" alt=""><figcaption></figcaption></figure>

Perhaps I should have tried learning vim as suggested by my university professors

### Mirror Mirror on the Wall...

Another line of thought surged regarding the reference to a mirror (I'm still avoiding the "Helikopter"), perhaps it has something to do with reversing?

I tried the text backwards. Nothing

I tried appending the reverse of the text to the end of the normal string, as though there is a "mirror" between them. Nothing

I tried a bunch of things, and to spare you the painful details, nothing seemed to work and I was at a lost

At least I learned how to exit vim

## Rotate Your Way of Thinking

Eventually, I did what I should have done at the start and noticed something peculiar about `hjpgs{ebg13_n6q5ro2}`

* It has open and closing curly braces `{}`
* There are 5 letters prior to the first bracket, similar to the format of the flags `uwctf{...}`
* Comparing the respective letters `u -> h`, `w -> j` etc. a certain pattern arises...

The letters `hjpgs` are precisely 13 characters away from `uwctf`

This is ROT 13

{% embed url="https://rot13.com/" %}

It wasn't until later that I read the comment in this StackOverflow post which provided the answer

{% embed url="https://stackoverflow.com/questions/45086981/what-are-the-vim-commands-that-start-with-g" %}

```
g?             2  Rot13 encoding operator
```

Another thing I missed was that there was literally 13 in the flag

At least I learned how to exit vim

## Flag

`uwctf{rot13_a6d5eb2}`
