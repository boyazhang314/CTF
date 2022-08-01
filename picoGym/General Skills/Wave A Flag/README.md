# Wave a flag
**Points: 10**

## Description
Can you invoke help flags for a tool or binary? [This program](https://mercury.picoctf.net/static/a14be2648c73e3cda5fc8490a2f476af/warm) has extraordinarily helpful information...

## Hints
<details>
<summary>Hint 1</summary>

> This program will only work in the webshell or another Linux computer.
</details>

<details>
<summary>Hint 2</summary>

> To get the file accessible in your shell, enter the following in the Terminal prompt: \
> ```$ wget https://mercury.picoctf.net/static/a14be2648c73e3cda5fc8490a2f476af/warm```
</details>

<details>
<summary>Hint 3</summary>

Run this program by entering the following in the Terminal prompt: ```$ ./warm```, but you'll first have to make it executable with ```$ chmod +x warm```
</details>

<details>
<summary>Hint 4</summary>

```-h``` and ```--help``` are the most common arguments to give to programs to get more information from them!
</details>

<details>
<summary>Hint 5</summary>

Not every program implements help features like ```-h``` and ```--help```.
</details>

## Solution
Download the program, which appears to be an executable

Permit execution permissions to run the program \
```chmod +x warm```

Then we can execute the file with ```./warm``` to get:
```Hello user! Pass me a -h to learn what I can do!```

As explained in the description, we should pass in the help flag ```-h``` \
(Note in this instance, the flag ```-help``` is not supported and returns ```I don't know what '-help' means! I do know what -h means though!```)

Run ```./warm -h``` to yield the flag

## Flag 
```picoCTF{b1scu1ts_4nd_gr4vy_755f3544}```